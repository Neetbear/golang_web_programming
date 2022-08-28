package internal

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrDuplicateName = errors.New("err duplicate username")
	ErrNotExist      = errors.New("err not exist MemberShip")
)

type CreateMember struct {
	Username       string
	MembershipType string
}

type UpdateMember struct {
	ID             string
	Username       string
	MembershipType string
}

type Repository interface {
	Create(*CreateMember) (*Membership, error)
	Update(*UpdateMember) (*Membership, error)
	Get(id string) *Membership
	Delete(id string) error
}

type repository struct {
	data map[string]Membership
}

func NewRepository(data map[string]Membership) Repository {
	return &repository{data: data}
}

func (r *repository) Create(member *CreateMember) (*Membership, error) {
	if r.isExist(member.Username) {
		return nil, ErrDuplicateName
	}

	m := NewMembership(uuid.New().String(), member)
	r.data[m.ID] = *m
	return m, nil
}

func (r *repository) Update(member *UpdateMember) (*Membership, error) {
	if r.isExist(member.Username) {
		return nil, ErrDuplicateName
	}

	m := r.data[member.ID]
	m.Update(member)
	r.data[member.ID] = m
	return &m, nil
}

func (r *repository) Get(id string) *Membership {
	if m, ok := r.data[id]; ok {
		return &m
	}

	return nil
}

func (r *repository) Delete(id string) error {
	if _, ok := r.data[id]; !ok {
		return ErrNotExist
	}

	delete(r.data, id)
	return nil
}

func (r *repository) isExist(username string) bool {
	for _, m := range r.data {
		if m.IsEqualName(username) {
			return true
		}
	}
	return false
}
