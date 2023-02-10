package internal

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Create(request CreateRequest) (CreateResponse, error) {
	switch {
	case s.repository.ExistName(request.UserName):
		return CreateResponse{}, errors.New("이미 존재하는 사용자 이름")
	case request.UserName == "":
		return CreateResponse{}, errors.New("빈 사용자 이름")
	case request.MembershipType == "":
		return CreateResponse{}, errors.New("빈 멤버십")
	case !slices.Contains([]string{"naver", "toss", "payco"}, request.MembershipType):
		return CreateResponse{}, errors.New("멤버십 타입 오류")
	}

	membership := Membership{uuid.New().String(), request.UserName, request.MembershipType}
	s.repository.Create(membership)
	return CreateResponse{
		ID:             membership.ID,
		MembershipType: membership.MembershipType,
	}, nil
}

func (s *Service) GetByID(id string) (GetResponse, error) {
	switch {
	case id == "" || !s.repository.Exist(id):
		return GetResponse{}, errors.New("찾을 수 없는 사용자 이름")
	}

	membership, err := s.repository.GetById(id)
	if err != nil {
		return GetResponse{}, nil
	}
	return GetResponse{
		ID:             membership.ID,
		UserName:       membership.UserName,
		MembershipType: membership.MembershipType,
	}, nil
}

func (s *Service) Update(request UpdateRequest) (UpdateResponse, error) {
	if !s.repository.Exist(request.ID) {
		return UpdateResponse{}, errors.New("찾을 수 없는 사용자 이름")
	}

	membership := s.repository.data[request.ID]
	switch {
	case membership.UserName != request.UserName && s.repository.ExistName(request.UserName):
		return UpdateResponse{}, errors.New("이미 존재하는 사용자 이름")
	case request.UserName == "":
		return UpdateResponse{}, errors.New("빈 사용자 이름")
	case request.MembershipType == "":
		return UpdateResponse{}, errors.New("빈 멤버십")
	case !slices.Contains([]string{"naver", "toss", "payco"}, request.MembershipType):
		return UpdateResponse{}, errors.New("멤버십 타입 오류")
	}

	membership, err := s.repository.Update(Membership(request))
	if err != nil {
		return UpdateResponse{}, nil
	}
	return UpdateResponse{
		ID:             membership.ID,
		UserName:       membership.UserName,
		MembershipType: membership.MembershipType,
	}, nil
}

func (s *Service) DeleteById(id string) error {
	switch {
	case id == "" || !s.repository.Exist(id):
		return errors.New("찾을 수 없는 사용자 이름")
	}

	err := s.repository.DeleteById(id)
	if err != nil {
		return err
	}
	return nil
}
