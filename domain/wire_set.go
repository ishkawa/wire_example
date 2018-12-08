package domain

import (
	"github.com/google/wire"
)

var WireSet = wire.NewSet(NewFooService)
