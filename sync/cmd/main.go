package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	eg, ec := errgroup.WithContext(ctx)
	eg.Go(func() error {
		time.Sleep(time.Second * 10)
		return errors.New("timeout")
	})
	err := eg.Wait()
	fmt.Printf("error: %v, ec err: %v", err, ec.Err())
}
