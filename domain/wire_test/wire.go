// +build wireinject

package wire_test

import (
	"github.com/golang/mock/gomock"
	"github.com/google/wire"
	"github.com/ishkawa/wire_example/domain"
	"github.com/ishkawa/wire_example/domain/mock_domain"
)

func InitializeFooService(ctrl *gomock.Controller) (fooService domain.FooService) {
	wire.Build(domain.WireSet, mock_domain.WireSet)
	return
}
