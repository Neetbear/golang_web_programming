package service

import (
	"github.com/google/uuid"
	"golang_web_programming/internal/dto"
	"golang_web_programming/internal/repository"
)

type MemberService interface {
	Create(request dto.CreateRequest) (dto.CreateResponse, error)
	Update(request dto.UpdateRequest) (dto.UpdateResponse, error)
	Delete(id string) error
	Get(id string) (dto.GetResponse, error)
}

type application struct {
	repository repository.Repository
}

func NewApplication(repository repository.Repository) MemberService {
	return &application{repository: repository}
}

func (app *application) Create(request dto.CreateRequest) (dto.CreateResponse, error) {
	if app.repository.IsExistBlackList(request.UserName, request.MembershipType) {
		return dto.CreateResponse{}, ErrCantSignUp
	}
	membership := NewMembership(request)
	m, err := app.repository.Create(membership)
	if err != nil {
		return dto.CreateResponse{}, err
	}
	return dto.NewCreateResponse(m.ID(), m.Type()), nil
}

func (app *application) Update(request dto.UpdateRequest) (dto.UpdateResponse, error) {
	m, err := app.repository.Update(&repository.UpdateRequest{
		ID:             request.ID,
		UserName:       request.UserName,
		MembershipType: request.MembershipType,
	})
	if err != nil {
		return dto.UpdateResponse{}, err
	}
	return dto.NewUpdateResponse(m.ID(), m.Username(), m.Type()), nil
}

func (app *application) Delete(id string) error {
	if err := app.repository.Delete(id); err != nil {
		return err
	}
	return nil
}

func (app *application) Get(id string) (dto.GetResponse, error) {
	m := app.repository.Get(id)
	if m == nil {
		return dto.GetResponse{}, nil
	}
	return dto.NewGetResponse(m.ID(), m.Username(), m.Type()), nil
}

func NewMembership(request dto.CreateRequest) *repository.Membership {
	return repository.CreateMembership(uuid.New().String(), request.UserName, request.MembershipType)
}
