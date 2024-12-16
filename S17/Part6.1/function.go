package function

import (
	"context"
	"time"
)

func main() {
	InternalCancellationExample()
}

func InternalCancellationExample() {
	ctx := context.Background() // ساخت یک کانتکس
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		for {
			if time.Now().Second()%15 == 0 {
				cancel()
			}
		}
	}()

	go Process1(ctx, 0)
	Process1(ctx, 0)
}

func Process1(ctx context.Context, num int) {
	for {
		select {
		case <-ctx.Done():
			println("task cancelled")
			return
		default:
			num++
			println("processing: ", num)
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func Process2(ctx context.Context, num int) {
	for {
		select {
		case <-ctx.Done():
			println("task cancelled")
			return
		default:
			num--
			println("processing: ", num)
			time.Sleep(time.Millisecond * 500)
		}
	}
}
