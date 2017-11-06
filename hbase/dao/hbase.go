package dao

import (
	"context"

	"go-exercise/hbase/model"

	"bytes"
	"fmt"
	"github.com/tsuna/gohbase/hrpc"
	"strconv"
	"time"
)

const (
	_t         = "local:user"
	_f         = "base"
	_cUsername = "username"
	_cEmail    = "email"
	_timeout   = time.Millisecond * 100
)

var (
	_fB         = []byte(_f)
	_cUsernameB = []byte(_cUsername)
	_cEmailB    = []byte(_cEmail)
)

func rowKey(id int64) string {
	return strconv.FormatInt(id, 10)
}

func (d *Dao) PutUser(c context.Context, u *model.User) (err error) {
	values := map[string]map[string][]byte{
		_f: {_cUsername: []byte(u.Username),
			_cEmail: []byte(u.Email),
		}}
	key := rowKey(u.ID)
	ctx, cancel := context.WithTimeout(c, _timeout)
	defer cancel()
	var req *hrpc.Mutate
	if req, err = hrpc.NewPutStr(ctx, _t, key, values); err != nil {
		fmt.Errorf("hrpc.NewPutStr(%s, %s) error(%v)", _t, key, err)
		return
	}
	if _, err = d.client.Put(req); err != nil {
		fmt.Errorf("dao.client.Put(%s) errir(%v)", key, err)
	}
	return
}

func (d *Dao) User(c context.Context, id int64) (res *model.User, err error) {
	key := rowKey(id)
	ctx, cancel := context.WithTimeout(c, _timeout)
	defer cancel()
	var req *hrpc.Get
	if req, err = hrpc.NewGetStr(ctx, _t, key); err != nil {
		fmt.Errorf("hrpc.NewGetStr(%s, %s) error(%v)", _t, key, err)
		return
	}
	var r *hrpc.Result
	if r, err = d.client.Get(req); err != nil {
		fmt.Errorf("dao.client.Get(%s) errir(%v)", key, err)
	}
	if len(r.Cells) == 0 {
		return
	}
	u := new(model.User)
	for _, cell := range r.Cells {
		if bytes.Equal(cell.Family, _fB) {
			switch {
			case bytes.Equal(cell.Qualifier, _cUsernameB):
				u.Username = string(cell.Value)
			case bytes.Equal(cell.Qualifier, _cEmailB):
				u.Email = string(cell.Value)
			}

		}
	}
	return
}
