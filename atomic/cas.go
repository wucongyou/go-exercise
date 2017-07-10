package atomic

import (
	"fmt"
	"sync/atomic"
)

func CAS() {
	var v uint64 = 0
	vPtr := &v
	ok := atomic.CompareAndSwapUint64(vPtr, 0, 1)
	fmt.Printf("ok: %v\n", ok)
	ok = atomic.CompareAndSwapUint64(vPtr, 0, 1)
	fmt.Printf("ok: %v\n", ok)
	fmt.Printf("v: %d, vPtr: %v\n", v, &vPtr)

	nowV := atomic.LoadUint64(vPtr)
	fmt.Printf("now v: %d, now vPtr: %v\n", nowV, &nowV)
}
