package mock_domain

import (
	"github.com/google/wire"
	"github.com/ishkawa/wire_example/domain"
)

var WireSet = wire.NewSet(
	NewMockFooRepository, wire.Bind(new(domain.FooRepository), new(MockFooRepository)))
