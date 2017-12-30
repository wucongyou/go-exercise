package jvm

import (
	"testing"
)

func TestParseFile(t *testing.T) {
	res, err := ParseFile("/tmp/HelloWorld.class")
	if err != nil {
		t.Errorf("failed to parse file, error(%v)", err)
		t.FailNow()
	}
	str, err := res.Format()
	if err != nil {
		t.FailNow()
	}
	t.Logf("format: \n%s", str)
}
