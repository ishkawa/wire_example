package domain_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ishkawa/wire_example/domain"
	"github.com/ishkawa/wire_example/domain/mock_domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type FooServiceTestSuite struct {
	suite.Suite

	ctx      context.Context
	ctrl     *gomock.Controller
	provider *mock_domain.Provider
	service  domain.FooService
}

func TestFooServiceTestSuite(t *testing.T) {
	suite.Run(t, new(FooServiceTestSuite))
}

func (suite *FooServiceTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.ctrl = gomock.NewController(suite.T())
	suite.provider = mock_domain.NewProvider(suite.ctrl)
	suite.service = mock_domain.InitializeFooService(suite.provider)
}

func (suite *FooServiceTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *FooServiceTestSuite) TestCreate() {
	name := "Dave"
	allocatedID := int64(123)

	suite.provider.MockFooRepository.EXPECT().
		Put(suite.ctx, gomock.Any()).
		Do(func(ctx context.Context, created *domain.Foo) {
			assert.Equal(suite.T(), int64(0), created.ID)
			assert.Equal(suite.T(), name, created.Name)
			created.ID = allocatedID
		})

	created, err := suite.service.Create(suite.ctx, name)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), allocatedID, created.ID)
	assert.Equal(suite.T(), name, created.Name)
}

func (suite *FooServiceTestSuite) TestCreateInvalid() {
	_, err := suite.service.Create(suite.ctx, "")
	assert.Error(suite.T(), err, domain.ErrInvalidArgument)
}

func (suite *FooServiceTestSuite) TestDuplicate() {
	source := &domain.Foo{ID: 123, Name: "Source"}
	allocatedID := int64(456)

	suite.provider.MockFooRepository.EXPECT().
		Get(suite.ctx, source.ID).
		Return(source, nil)

	suite.provider.MockFooRepository.EXPECT().
		Put(suite.ctx, gomock.Any()).
		Do(func(ctx context.Context, duplicated *domain.Foo) {
			assert.Equal(suite.T(), int64(0), duplicated.ID)
			assert.Equal(suite.T(), source.Name, duplicated.Name)
			duplicated.ID = allocatedID
		})

	duplicated, err := suite.service.Duplicate(suite.ctx, source.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), allocatedID, duplicated.ID)
	assert.Equal(suite.T(), source.Name, duplicated.Name)
}
