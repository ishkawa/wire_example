// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/ishkawa/wire_example/domain"
	"github.com/ishkawa/wire_example/infra"
)

// Injectors from wire.go:

func InitializeFooService() domain.FooService {
	fooRepository := infra.NewFooRepository()
	fooService := domain.NewFooService(fooRepository)
	return fooService
}

// wire.go:

var wireSet = wire.NewSet(domain.WireSet, infra.WireSet)
