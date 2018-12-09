package domain

import "context"

type FooService interface {
	Create(ctx context.Context, name string) (created *Foo, err error)
	Duplicate(ctx context.Context, id int64) (duplicated *Foo, err error)
}

func NewFooService(fooRepository FooRepository) FooService {
	return &fooService{fooRepository: fooRepository}
}

type fooService struct {
	fooRepository FooRepository
}

func (service *fooService) Create(ctx context.Context, name string) (created *Foo, err error) {
	if name == "" {
		err = ErrInvalidArgument
		return
	}
	created = &Foo{Name: name}
	err = service.fooRepository.Put(ctx, created)
	return
}

func (service *fooService) Duplicate(ctx context.Context, id int64) (duplicated *Foo, err error) {
	source, err := service.fooRepository.Get(ctx, id)
	if err != nil {
		return
	}

	duplicated = &Foo{Name: source.Name}
	err = service.fooRepository.Put(ctx, duplicated)
	return
}
