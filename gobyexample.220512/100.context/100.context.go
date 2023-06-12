package main

import "context"

func main() {
	ctx := context.Background()
	ctx2 := context.WithValue(ctx, "name", "wukong")
	print(ctx)
	print(ctx2)
	// context.TODO()
}
