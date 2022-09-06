package repository

type UpdateRequest struct {
	ID             string
	UserName       string
	MembershipType string
}

type Repository interface {
	Create(membership *Membership) (*Membership, error)
	Update(request *UpdateRequest) (*Membership, error)
	Get(id string) *Membership
	Delete(id string) error
	IsExistBlackList(username, membershipType string) bool
}

type repository struct {
	data      map[string]Membership
	blackList map[string]string
}

func NewRepository(data map[string]Membership) Repository {
	return &repository{data: data, blackList: map[string]string{}}
}

func (r *repository) Create(member *Membership) (*Membership, error) {
	if r.isExist(member.userName) {
		return nil, ErrDuplicateName
	}

	r.data[member.iD] = *member
	return member, nil
}

func (r *repository) IsExistBlackList(username, membershipType string) bool {
	if platform, ok := r.blackList[username]; ok && platform == membershipType {
		return true
	}
	return false
}

func (r *repository) Update(updateRequest *UpdateRequest) (*Membership, error) {
	if r.isExist(updateRequest.UserName) {
		return nil, ErrDuplicateName
	}

	m := r.data[updateRequest.ID]
	m.Update(updateRequest)
	r.data[updateRequest.ID] = m
	return &m, nil
}

func (r *repository) Get(id string) *Membership {
	if m, ok := r.data[id]; ok {
		return &m
	}

	return nil
}

func (r *repository) Delete(id string) error {
	if m, ok := r.data[id]; !ok {
		return ErrNotExist
	} else {
		r.blackList[m.userName] = m.membershipType
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
