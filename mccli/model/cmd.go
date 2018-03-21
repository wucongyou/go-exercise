package model

import "strings"

type Cmd struct {
	Name        string
	MinParamNum int
	MaxParamNum int
	Params      []string
	Usage       string
}

const (
	_cmdDel      = "del"
	_cmdSet      = "set"
	_cmdGet      = "get"
	_cmdKeys     = "keys"
	_cmdFlushAll = "flushall"
)

var (
	CmdDel = &Cmd{
		Name:        _cmdDel,
		MinParamNum: 1,
		MaxParamNum: 1,
		Usage:       "del key",
	}
	CmdSet = &Cmd{
		Name:        _cmdSet,
		MinParamNum: 2,
		MaxParamNum: 4,
		Usage:       "set key value [expire seconds or absolute ts, default 0] [flag, default 0]",
	}
	CmdGet = &Cmd{
		Name:        _cmdGet,
		MinParamNum: 1,
		MaxParamNum: 2,
		Usage:       "get key [binary|string, default string]",
	}
	CmdKeys = &Cmd{
		Name:        _cmdKeys,
		MinParamNum: 1,
		MaxParamNum: 1,
		Usage:       "keys pattern",
	}
	CmdFlushAll = &Cmd{
		Name:        _cmdFlushAll,
		MinParamNum: 1,
		MaxParamNum: 1,
		Usage:       "flushall",
	}
)

// ParseCMD parse cmd string to cmd.
func ParseCMD(key string) *Cmd {
	switch strings.ToLower(key) {
	case _cmdDel:
		return CmdDel
	case _cmdSet:
		return CmdSet
	case _cmdGet:
		return CmdGet
	case _cmdKeys:
		return CmdKeys
	case _cmdFlushAll:
		return CmdFlushAll
	default:
		return nil
	}
}
