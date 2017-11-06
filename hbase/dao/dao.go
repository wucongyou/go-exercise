package dao

import "github.com/tsuna/gohbase"

type Dao struct {
	client gohbase.Client
}

func New(zk string) *Dao {
	return &Dao{
		client: gohbase.NewClient(zk),
	}
}
