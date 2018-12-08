package domain_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ishkawa/wire_example/domain/wire_test"
	"github.com/stretchr/testify/assert"
)

func TestDuplicate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := wire_test.InitializeFooService(ctrl)

	ctx := context.Background()
	duplicated, err := service.Duplicate(ctx, 123)
	assert.NoError(t, err)
	assert.Equal(t, "foo", duplicated.Name)
}
