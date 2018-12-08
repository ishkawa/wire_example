package domain

import "context"

//go:generate mockgen -destination mock_domain/foo_repository.go github.com/ishkawa/wire_example/domain FooRepository
type FooRepository interface {
	Get(ctx context.Context, id int64) (foo *Foo, err error)
	Put(ctx context.Context, foo *Foo) (err error)
}
