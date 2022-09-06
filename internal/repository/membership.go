package repository

type Membership struct {
	iD             string
	userName       string
	membershipType string
}

func CreateMembership(id, userName, membershipType string) *Membership {
	return &Membership{
		iD:             id,
		userName:       userName,
		membershipType: membershipType,
	}
}

func (m *Membership) Update(member *UpdateRequest) {
	m.userName = member.UserName
	m.membershipType = member.MembershipType
}

func (m *Membership) IsEqualName(username string) bool {
	return m.userName == username
}

func (m *Membership) Type() string {
	return m.membershipType
}

func (m *Membership) Username() string {
	return m.userName
}

func (m *Membership) ID() string {
	return m.iD
}
