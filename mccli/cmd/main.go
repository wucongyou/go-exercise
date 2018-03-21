package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"

	"go-exercise/mccli/dao"
	"go-exercise/mccli/model"

	"github.com/bradfitz/gomemcache/memcache"
)

var (
	ErrSyntaxErr         = errors.New("ERR syntax error")
	ErrNilCommand        = errors.New("ERR nil command")
	ErrUnknownCommand    = errors.New("ERR unknown command")
	ErrWrongArgumentsNum = errors.New("ERR wrong number of arguments")
)

func main() {
	var (
		host string
		port int64
		cmd  string
	)
	flag.StringVar(&host, "h", "127.0.0.1", "host")
	flag.Int64Var(&port, "p", 11211, "port")
	flag.StringVar(&cmd, "cmd", "", "cmd")
	flag.Parse()
	d := dao.New(&dao.Conf{
		Host: host,
		Port: port,
	})
	if _, err := d.Get("ping"); err != nil && err != memcache.ErrCacheMiss {
		fmt.Printf("fail to connect %s:%d, error(%v)", host, port, err)
		return
	}
	if cmd != "" {
		args := strings.Split(cmd, " ")
		if len(args) == 0 {
			return
		}
		if args[0] == "exit" {
			return
		}
		recv(d, args)
		return
	}

	fmt.Printf("connected to %s:%d\n", host, port)
	for {
		fmt.Printf("%s:%d> ", host, port)
		cmdLine := ""
		scanLine(&cmdLine)
		args := strings.Split(cmdLine, " ")
		if len(args) == 0 {
			continue
		}
		if args[0] == "exit" {
			return
		}
		recv(d, args)
	}
}

func scanLine(s *string) {
	var c byte
	var err error
	var b []byte
	for err == nil {
		_, err = fmt.Scanf("%c", &c)
		if c != '\n' {
			b = append(b, c)
		} else {
			break
		}
	}
	*s = string(b)
}

func recv(d *dao.Dao, args []string) {
	if len(args) == 0 || args[0] == "" {
		fmt.Println(ErrNilCommand)
		return
	}
	cmd := model.ParseCMD(args[0])
	if cmd == nil {
		fmt.Printf("%v '%s'\n", ErrUnknownCommand, args[0])
		return
	}
	paramNum := len(args) - 1
	if paramNum < cmd.MinParamNum || paramNum > cmd.MaxParamNum {
		fmt.Printf("%v for '%s' command\nUsage: %s\n", ErrWrongArgumentsNum, cmd.Name, cmd.Usage)
		return
	}
	switch cmd {
	case model.CmdSet: // set key value [expiration] [flags]
		key := args[1]
		value := args[2]
		exp := int64(0)
		flags := uint64(0)
		var err error
		if paramNum >= 3 {
			exp, err = strconv.ParseInt(args[3], 10, 32)
			if err != nil {
				err = ErrSyntaxErr
				fmt.Printf("%v, expiration should be int\n", err)
				return
			}
		}
		if paramNum >= 4 {
			flags, err = strconv.ParseUint(args[4], 10, 32)
			if err != nil {
				err = ErrSyntaxErr
				fmt.Printf("%v, flags should be int\n", err)
				return
			}
		}
		item := &memcache.Item{
			Key:        key,
			Value:      []byte(value),
			Expiration: int32(exp),
			Flags:      uint32(flags),
		}
		if err = d.Set(item); err != nil {
			fmt.Println(err)
			return
		}
	case model.CmdDel: // del key
		key := args[1]
		if err := d.Delete(key); err != nil {
			fmt.Println(model.FromErr(err).Format())
			return
		}
	case model.CmdGet: // get key [binary|string]
		key := args[1]
		item, err := d.Get(key)
		if err != nil {
			fmt.Println(model.FromErr(err).Format())
			return
		}
		vtp := ""
		if paramNum >= 2 {
			vtp = args[2]
		}
		fmt.Println(model.FromItem(item, vtp).Format())
		return
	case model.CmdKeys:
		fmt.Println("keys command not impled yet")
	case model.CmdFlushAll:
		err := d.FlushAll()
		if err != nil {
			fmt.Println(model.FromErr(err).Format())
			return
		}
	}
	fmt.Println("OK")
}
