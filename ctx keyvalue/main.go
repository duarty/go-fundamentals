package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "tkn")
	booking(ctx)
}

func booking(ctx context.Context) {
	token := ctx.Value("tkn")
	fmt.Println(token)
}
