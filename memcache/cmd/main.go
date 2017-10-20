package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"

	"go-exercise/memcache/dao"
)

var (
	ErrSyntaxErr = errors.New("ERR syntax error")
)

func main() {
	var (
		host string
		port int64
	)
	flag.StringVar(&host, "h", "127.0.0.1", "host")
	flag.Int64Var(&port, "p", 11211, "port")
	flag.Parse()
	d := dao.New(&dao.Conf{
		Host: host,
		Port: port,
	})
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
	if len(args) == 0 {
		fmt.Println("Err nil command")
		return
	}
	cmd := parseCMD(args[0])
	if cmd == nil {
		fmt.Printf("ERR unknown command '%s'\n", args[0])
		return
	}
	paramNum := len(args) - 1
	if paramNum < cmd.MinParamNum || paramNum > cmd.MaxParamNum {
		fmt.Printf("ERR wrong number of arguments for '%s' command\nUsage: %s\n", cmd.Name, cmd.Usage)
		return
	}
	var err error
	switch cmd {
	// set key value [expiration] [flags]
	case _cmdSet:
		key := args[1]
		value := args[2]
		exp := int64(0)
		flags := uint64(0)
		if paramNum >= 3 {
			exp, err = strconv.ParseInt(args[3], 10, 32)
			if err != nil {
				err = ErrSyntaxErr
				fmt.Printf("%v, expiration should be int\nUsage: %s\n", err, cmd.Usage)
				return
			}
		}
		if paramNum >= 4 {
			flags, err = strconv.ParseUint(args[4], 10, 32)
			if err != nil {
				err = ErrSyntaxErr
				fmt.Printf("%v, flags should be int\nUsage: %s\n", err, cmd.Usage)
				return
			}
		}
		if err = d.Set(key, []byte(value), int32(exp), uint32(flags)); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("OK")
		return
		// del key
	case _cmdDel:
		key := args[1]
		if err = d.Delete(key); err != nil {
			resp := new(dao.Resp)
			resp.FromErr(err)
			fmt.Println(resp.Format())
			return
		}
		fmt.Println("OK")
		return
		// get key [binary|string]
	case _cmdGet:
		key := args[1]
		vtp := ""
		if paramNum >= 2 {
			vtp = args[2]
		}
		res := d.GetWithType(key, vtp)
		fmt.Println(res.Format())
		return
	case _cmdKeys:
		fmt.Println("keys command not impled yet")
		return
	}
}

type Cmd struct {
	Name        string
	MinParamNum int
	MaxParamNum int
	Params      []string
	Usage       string
}

const (
	_actionDel  = "del"
	_actionSet  = "set"
	_actionGet  = "get"
	_actionKeys = "keys"
)

var (
	_cmdDel = &Cmd{
		Name:        _actionDel,
		MinParamNum: 1,
		MaxParamNum: 1,
		Usage:       "del key",
	}
	_cmdSet = &Cmd{
		Name:        _actionSet,
		MinParamNum: 2,
		MaxParamNum: 4,
		Usage:       "set key value [expire seconds or absolute ts, default 0] [flag, default 0]",
	}
	_cmdGet = &Cmd{
		Name:        _actionGet,
		MinParamNum: 1,
		MaxParamNum: 2,
		Usage:       "get key [binary|string, default string]",
	}
	_cmdKeys = &Cmd{
		Name:        _actionKeys,
		MinParamNum: 1,
		MaxParamNum: 1,
		Usage:       "keys pattern",
	}
)

func parseCMD(key string) *Cmd {
	switch key {
	case _actionDel:
		return _cmdDel
	case _actionSet:
		return _cmdSet
	case _actionGet:
		return _cmdGet
	case _actionKeys:
		return _cmdKeys
	default:
		return nil
	}
}
