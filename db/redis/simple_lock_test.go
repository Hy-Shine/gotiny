package redis

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/spf13/cast"
)

func TestLock(t *testing.T) {
	rCfg := &cfg{
		Address: "127.0.0.1",
		Port:    6379,
	}

	client = GetInstance(*rCfg)
	ctx := context.Background()
	type args struct {
		n int
		r bool
	}

	ch := make(chan args)
	go func() {
		for i := 0; i < 100; i++ {
			go func(i int) {
				got, _ := Lock(ctx, "test", 15*time.Second, cast.ToString(i))
				if got {
					fmt.Printf("N0.%d got lock\n", i)
				} else {
					fmt.Printf("N0.%d failed\n", i)
				}
				ch <- args{r: got, n: i}
			}(i)
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v.n, v.r)
	}
}
