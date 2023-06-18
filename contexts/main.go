package main

import (
	"context"
	"fmt"
)

func doSomething(ctx context.Context) {
	// ret val as is
	fmt.Printf("doSomething: name's value is %s\n", ctx.Value("name"))

	// def 2nd ctx, with diff key value
	anotherCtx := context.WithValue(ctx, "name", "topaz")
	// prints base on given context
	doAnother(anotherCtx)

	// print (1st) ctx again, to show its not mixed with 2nd
	fmt.Printf("doSomething: name's value is %s\n", ctx.Value("name"))
}

func doAnother(ctx context.Context) {
	fmt.Printf("doAnother: name's value is %s\n", ctx.Value("name"))
}

func main() {
	// def 1st ctx
	ctx := context.Background()

	// add key & vlaue to the ctx
	ctx = context.WithValue(ctx, "name", "anna")

	doSomething(ctx)
}
