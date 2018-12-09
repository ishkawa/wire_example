package main

import (
	"context"
	"log"

	"github.com/ishkawa/wire_example/infra"
)

func main() {
	fooService := InitializeFooService()
	ctx := context.Background()

	err := fooService.Duplicate(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}

	infra.PrintStoreDump()
}
