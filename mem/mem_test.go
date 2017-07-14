package mem

import (
	"testing"
)

func TestMake(t *testing.T) {
	b := &Book{
		ID:     1,
		Name:   "1Q84",
		Author: "Murakami Haruki",
		Price:  36,
	}
	m := make(map[int64]*Book, 0)
	t.Logf("m == nil: %t", m == nil)
	m[b.ID] = b
	t.Logf("m: %v", m)
	n := new(map[int64]*Book)
	t.Logf("*n == nil: %t", *n == nil)
	*n = make(map[int64]*Book, 0)
	t.Logf("*n == nil: %t", *n == nil)
	(*n)[b.ID] = b
	t.Logf("n: %v", n)
}
