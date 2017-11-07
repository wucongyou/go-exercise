package dao

import (
	"context"
	"encoding/hex"
	"testing"

	"go-exercise/hbase/model"
)

func TestDao_PutUser(t *testing.T) {
	once.Do(startDao)
	u := &model.User{
		Mtime:    "2016-09-06 00:10:12",
		Uid:      75174,
		ID:       0x00111863,
		Username: "wcy",
		Email:    "foo@bar.com",
	}
	if err := d.PutUser(context.TODO(), u); err != nil {
		t.Errorf("dao.PutUser(%+v) error(%v)", u, err)
		t.FailNow()
	}
	if res, err := d.User(context.TODO(), u.ID); err != nil {
		t.Errorf("dao.User(%d) error(%v)", u.ID, err)
		t.FailNow()
	} else {
		t.Logf("res: %v", res)
	}
}

func TestDao_RowKey(t *testing.T) {
	tStr := "2016-09-06 00:10:12"
	uidStr := "75174"
	idStr := "1"
	k := rk(tStr, uidStr, idStr)
	t.Logf("res: %s", hex.EncodeToString(k))
}
