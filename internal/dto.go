package internal

const (
	naver = "naver"
	payco = "payco"
	toss  = "toss"
)

type CreateRequest struct {
	UserName       string
	MembershipType string
}

func (c *CreateRequest) ToCreateMember() *CreateMember {
	return &CreateMember{
		Username:       c.UserName,
		MembershipType: c.MembershipType,
	}
}

func (c *CreateRequest) IsValid() bool {
	if c.UserName == "" || c.MembershipType == "" {
		return false
	}
	t := c.MembershipType
	return isValidPlatform(t)
}

type CreateResponse struct {
	ID             string
	MembershipType string
}

type UpdateRequest struct {
	ID             string
	UserName       string
	MembershipType string
}

func (u *UpdateRequest) ToUpdateMember() *UpdateMember {
	return &UpdateMember{
		ID:             u.ID,
		Username:       u.UserName,
		MembershipType: u.MembershipType,
	}
}

func (r *UpdateRequest) IsValid() bool {
	return r.ID != "" && r.UserName != "" && isValidPlatform(r.MembershipType)
}

func isValidPlatform(t string) bool {
	return t == naver || t == payco || t == toss
}

type UpdateResponse struct {
	ID             string
	UserName       string
	MembershipType string
}
type GetResponse struct {
	ID             string
	UserName       string
	MembershipType string
}
