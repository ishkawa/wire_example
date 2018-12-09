// +build wireinject

package mock_domain

import (
	"github.com/golang/mock/gomock"
	"github.com/google/wire"
	"github.com/ishkawa/wire_example/domain"
)

type MockRepositorySet struct {
	MockFooRepository *MockFooRepository
}

func NewProvider(ctrl *gomock.Controller) *MockRepositorySet {
	return &MockRepositorySet{
		MockFooRepository: NewMockFooRepository(ctrl),
	}
}

func ProvideFooRepository(repositorySet *MockRepositorySet) *MockFooRepository {
	return repositorySet.MockFooRepository
}

var providerSet = wire.NewSet(
	domain.WireSet,
	ProvideFooRepository,
	wire.Bind(new(domain.FooRepository), new(MockFooRepository)))

func InitializeFooService(provider *MockRepositorySet) (fooService domain.FooService) {
	wire.Build(providerSet)
	return
}
