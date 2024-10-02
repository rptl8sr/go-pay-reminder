package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	res, err := Handler(ctx)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res.Body)
}
