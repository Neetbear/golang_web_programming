package internal

type Membership struct {
	ID             string
	UserName       string
	MembershipType string
}

func NewMembership(id string, member *CreateMember) *Membership {
	return &Membership{
		ID:             id,
		UserName:       member.Username,
		MembershipType: member.MembershipType,
	}
}

func (m *Membership) Update(member *UpdateMember) {
	m.UserName = member.Username
	m.MembershipType = member.MembershipType
}

func (m *Membership) ToCreateResponse() CreateResponse {
	return CreateResponse{
		ID:             m.ID,
		MembershipType: m.MembershipType,
	}
}

func (m *Membership) ToUpdateResponse() UpdateResponse {
	return UpdateResponse{
		ID:             m.ID,
		UserName:       m.UserName,
		MembershipType: m.MembershipType,
	}
}

func (m *Membership) ToGetResponse() GetResponse {
	return GetResponse{
		ID:             m.ID,
		UserName:       m.UserName,
		MembershipType: m.MembershipType,
	}
}

func (m *Membership) IsEqualName(username string) bool {
	return m.UserName == username
}
