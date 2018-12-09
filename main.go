package main

import (
	"context"
	"log"

	"github.com/ishkawa/wire_example/infra"
)

func main() {
	fooService := InitializeFooService()
	ctx := context.Background()

	created, err := fooService.Create(ctx, "John")
	if err != nil {
		log.Fatal(err)
	}

	_, err = fooService.Duplicate(ctx, created.ID)
	if err != nil {
		log.Fatal(err)
	}

	infra.PrintStoreDump()
}
