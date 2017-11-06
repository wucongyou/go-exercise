package dao

import "sync"

var (
	once sync.Once
	d    *Dao
)

func startDao() {
	d = New("127.0.0.1:2181")
}
