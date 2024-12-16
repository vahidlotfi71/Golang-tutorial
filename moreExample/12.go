package main

import (
	"context"
	"fmt"
)

func processRequest(ctx context.Context) {
	userId := ctx.Value("userId")
	if userId == nil {
		println("user not found")
	}
	fmt.Printf("userId:%+v ", userId)

}

func main() {
	ctx := context.WithValue(context.Background(), "userId", 212121)

	processRequest(ctx)
}
