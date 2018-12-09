// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/ishkawa/wire_example/domain"
	"github.com/ishkawa/wire_example/infra"
)

var wireSet = wire.NewSet(domain.WireSet, infra.WireSet)

func InitializeFooService() (fooService domain.FooService) {
	wire.Build(wireSet)
	return
}
