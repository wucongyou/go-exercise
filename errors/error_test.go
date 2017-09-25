package errors

import (
	"errors"
	"testing"
)

func TestClient(t *testing.T) {
	err := errors.New("net http request error")
	stackErr1 := Wrap("dao.Identify(), failed to call identify api", err)
	stackErr2 := Wrap("service.Identify()", stackErr1)
	stackErr3 := Wrap("http.Identify()", stackErr2)
	t.Logf("stack 1: \n%s", StackTrace(stackErr1))
	t.Logf("stack 2: \n%s", StackTrace(stackErr2))
	t.Logf("stack 3: \n%s", StackTrace(stackErr3))
}
