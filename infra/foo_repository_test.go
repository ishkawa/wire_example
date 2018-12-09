package infra_test

import (
	"context"
	"testing"

	"github.com/ishkawa/wire_example/domain"
	"github.com/ishkawa/wire_example/infra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type FooRepositoryTestSuite struct {
	suite.Suite

	ctx  context.Context
	repo domain.FooRepository
}

func TestFooRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(FooRepositoryTestSuite))
}

func (suite *FooRepositoryTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.repo = infra.NewFooRepository()
}

func (suite *FooRepositoryTestSuite) TearDownTest() {
	infra.ResetStore()
}

func (suite *FooRepositoryTestSuite) TestPutGet() {
	foo := &domain.Foo{ID: 123, Name: "Jane"}
	err := suite.repo.Put(suite.ctx, foo)
	assert.NoError(suite.T(), err)

	got, err := suite.repo.Get(suite.ctx, foo.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), foo, got)
}
