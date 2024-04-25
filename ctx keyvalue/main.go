package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", " ")
	booking(ctx)
}

func booking(ctx context.Context) {
	token := ctx.Value("tkn")
	fmt.Println(token)
}
