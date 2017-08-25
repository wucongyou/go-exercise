package concurrent

import (
	"fmt"

	"golang.org/x/sync/errgroup"
)

// Start start example.
func Start() {
	var g errgroup.Group
	for i := 0; i < 3; i++ {
		seq := i
		g.Go(func() error {
			fmt.Printf("ok: %d\n", seq)
			if seq == 1 {
				return fmt.Errorf("error: %d", seq)
			}
			return nil
		})
	}
	err := g.Wait()
	fmt.Println(err)
}
