// +build wireinject

package test_injector

import (
	"github.com/golang/mock/gomock"
	"github.com/google/wire"
	"github.com/ishkawa/wire_example/domain"
	"github.com/ishkawa/wire_example/domain/mock_domain"
)

type MockRepositorySet struct {
	MockFooRepository *mock_domain.MockFooRepository
}

func NewRepositorySet(ctrl *gomock.Controller) *MockRepositorySet {
	return &MockRepositorySet{
		MockFooRepository: mock_domain.NewMockFooRepository(ctrl),
	}
}

func ProvideFooRepository(repositorySet *MockRepositorySet) *mock_domain.MockFooRepository {
	return repositorySet.MockFooRepository
}

var providerSet = wire.NewSet(
	domain.WireSet,
	ProvideFooRepository,
	wire.Bind(new(domain.FooRepository), new(mock_domain.MockFooRepository)))

func InitializeFooService(provider *MockRepositorySet) (fooService domain.FooService) {
	wire.Build(providerSet)
	return
}
