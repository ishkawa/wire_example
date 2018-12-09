package infra

import "github.com/google/wire"

var WireSet = wire.NewSet(NewFooRepository)
