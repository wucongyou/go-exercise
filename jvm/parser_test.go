package jvm

import (
	"encoding/json"
	"testing"
)

func TestParseFile(t *testing.T) {
	res, err := ParseFile("/tmp/HelloWorld.class")
	if err != nil {
		t.Errorf("failed to parse file, error(%v)", err)
		t.FailNow()
	}
	str, _ := json.Marshal(res)
	t.Logf("res: %s", str)
}
