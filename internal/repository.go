package internal

import (
	"errors"
	"strconv"

	"golang.org/x/exp/slices"
)

type Repository struct {
	data map[string]Membership
}

func NewRepository(data map[string]Membership) *Repository {
	return &Repository{data: data}
}

func (r *Repository) exist(id string) bool {
	_, ok := r.data[id]
	return ok
}

func (r *Repository) existName(name string) bool {
	for _, value := range r.data {
		if value.UserName == name {
			return true
		}
	}
	return false
}

func (r *Repository) Create(request CreateRequest) (*Membership, error) {
	switch {
	case r.existName(request.UserName):
		return nil, errors.New("이미 존재하는 사용자 이름")
	case request.UserName == "":
		return nil, errors.New("빈 사용자 이름")
	case request.MembershipType == "":
		return nil, errors.New("빈 멤버십")
	case !slices.Contains([]string{"naver", "toss", "payco"}, request.MembershipType):
		return nil, errors.New("멤버십 타입 오류")
	}
	membership := Membership{ID: strconv.Itoa(len(r.data) + 1), UserName: request.UserName, MembershipType: request.MembershipType}
	r.data[membership.ID] = membership
	return &membership, nil
}

func (r *Repository) Update(request UpdateRequest) (*Membership, error) {
	if !r.exist(request.ID) {
		return nil, errors.New("찾을 수 없는 사용자 이름")
	}

	membership := r.data[request.ID]
	switch {
	case membership.UserName != request.UserName && r.existName(request.UserName):
		return nil, errors.New("이미 존재하는 사용자 이름")
	case request.UserName == "":
		return nil, errors.New("빈 사용자 이름")
	case request.MembershipType == "":
		return nil, errors.New("빈 멤버십")
	case !slices.Contains([]string{"naver", "toss", "payco"}, request.MembershipType):
		return nil, errors.New("멤버십 타입 오류")
	}

	membership.UserName = request.UserName
	membership.MembershipType = request.MembershipType
	r.data[request.UserName] = membership
	return &membership, nil
}

func (r *Repository) Delete(id string) error {
	switch {
	case id == "" || !r.exist(id):
		return errors.New("찾을 수 없는 사용자 이름")
	}
	delete(r.data, id)
	return nil
}
