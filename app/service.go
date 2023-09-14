package app

import (
	"context"

	"github.com/c7/todo.c7.se/todo"
	"github.com/google/uuid"
)

type Service struct {
	repository todo.Repository
}

func NewService(repository todo.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s Service) Add(_ context.Context, description string) (*todo.Item, error) {
	item := s.repository.Add(description)

	return item, nil
}

func (s Service) Remove(_ context.Context, id uuid.UUID) error {
	s.repository.Remove(id)

	return nil
}

func (s Service) Update(_ context.Context, id uuid.UUID, completed bool, description string) (*todo.Item, error) {
	item := s.repository.Update(id, completed, description)

	return item, nil
}

func (s Service) Search(_ context.Context, search string) (todo.List, error) {
	items := s.repository.Search(search)

	return items, nil
}

func (s Service) Get(_ context.Context, id uuid.UUID) (*todo.Item, error) {
	item := s.repository.Get(id)

	return item, nil
}

func (s Service) Sort(_ context.Context, ids []uuid.UUID) error {
	s.repository.Reorder(ids)

	return nil
}

func (s Service) List(context.Context) (todo.List, error) {
	return s.repository.All(), nil
}
