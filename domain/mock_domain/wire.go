// +build wireinject

package mock_domain

import (
	"github.com/golang/mock/gomock"
	"github.com/google/wire"
	"github.com/ishkawa/wire_example/domain"
)

type Provider struct {
	MockFooRepository *MockFooRepository
}

func NewProvider(ctrl *gomock.Controller) *Provider {
	return &Provider{
		MockFooRepository: NewMockFooRepository(ctrl),
	}
}

func ProvideFooRepository(provider *Provider) *MockFooRepository {
	return provider.MockFooRepository
}

func InitializeFooService(provider *Provider) (fooService domain.FooService) {
	wire.Build(
		domain.WireSet,
		ProvideFooRepository,
		wire.Bind(new(domain.FooRepository), new(MockFooRepository)))

	return
}
