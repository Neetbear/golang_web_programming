package internal

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidRequest = errors.New("err Invalid Input")
)

type Application struct {
	repository Repository
}

func NewApplication(repository Repository) *Application {
	return &Application{repository: repository}
}

func (app *Application) Create(request CreateRequest) (CreateResponse, error) {
	if !request.IsValid() {
		return CreateResponse{}, ErrInvalidRequest
	}

	m, err := app.repository.Create(request.ToCreateMember())
	if err != nil {
		return CreateResponse{}, wrapRepositoryError(err)
	}
	return m.ToCreateResponse(), nil
}

func (app *Application) Update(request UpdateRequest) (UpdateResponse, error) {
	if !request.IsValid() {
		return UpdateResponse{}, ErrInvalidRequest
	}

	m, err := app.repository.Update(request.ToUpdateMember())
	if err != nil {
		return UpdateResponse{}, wrapRepositoryError(err)
	}
	return m.ToUpdateResponse(), nil
}

func (app *Application) Delete(id string) error {
	if id == "" {
		return ErrInvalidRequest
	}
	if err := app.repository.Delete(id); err != nil {
		return wrapRepositoryError(err)
	}

	return nil
}

func (app *Application) Get(id string) (GetResponse, error) {
	if id == "" {
		return GetResponse{}, ErrInvalidRequest
	}

	m := app.repository.Get(id)
	if m == nil {
		return GetResponse{}, nil
	}
	return m.ToGetResponse(), nil
}

func wrapRepositoryError(err error) error {
	return fmt.Errorf("repository err: %v", err)
}
