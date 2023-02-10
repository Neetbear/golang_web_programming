package internal

import (
	"errors"
)

var ErrNotFoundMembership = errors.New("not found membership")

type Repository struct {
	data map[string]Membership
}

func NewRepository(data map[string]Membership) *Repository {
	return &Repository{data: data}
}

func (r *Repository) Exist(id string) bool {
	_, ok := r.data[id]
	return ok
}

func (r *Repository) ExistName(name string) bool {
	for _, value := range r.data {
		if value.UserName == name {
			return true
		}
	}
	return false
}

func (r *Repository) Create(membership Membership) {
	r.data[membership.ID] = membership
}

func (r *Repository) GetById(id string) (Membership, error) {
	for _, membership := range r.data {
		if membership.ID == id {
			return membership, nil
		}
	}
	return Membership{}, ErrNotFoundMembership
}

func (r *Repository) Update(membership Membership) (Membership, error) {
	for _, member := range r.data {
		if member.ID == membership.ID {
			member.MembershipType = membership.MembershipType
			member.UserName = membership.UserName
			return member, nil
		}
	}
	return Membership{}, ErrNotFoundMembership
}

func (r *Repository) DeleteById(id string) error {
	for _, member := range r.data {
		if member.ID == id {
			delete(r.data, id)
			return nil
		}
	}
	return ErrNotFoundMembership
}
