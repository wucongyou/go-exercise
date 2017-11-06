package dao

import (
	"context"
	"testing"

	"go-exercise/hbase/model"
)

func TestDao_PutUser(t *testing.T) {
	once.Do(startDao)
	u := &model.User{
		ID:       1,
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
