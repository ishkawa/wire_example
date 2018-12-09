package infra

import (
	"context"
	"errors"
	"log"

	"github.com/ishkawa/wire_example/domain"
)

var store = map[int64]*domain.Foo{}

type fooRepository struct{}

func NewFooRepository() domain.FooRepository {
	return &fooRepository{}
}

func (repo *fooRepository) Get(ctx context.Context, id int64) (foo *domain.Foo, err error) {
	if found, ok := store[id]; ok {
		foo = found
	} else {
		err = errors.New("no such entity")
	}
	return
}

func (repo *fooRepository) Put(ctx context.Context, foo *domain.Foo) (err error) {
	if foo.ID == 0 {
		foo.ID = int64(len(store) + 1)
	}
	store[foo.ID] = foo
	return
}

func PrintStoreDump() {
	for _, foo := range store {
		log.Println(foo)
	}
}
