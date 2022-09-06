package dto

const (
	naver = "naver"
	payco = "payco"
	toss  = "toss"
)

type UpdateRequestBody struct {
	UserName       string
	MembershipType string
}
type UpdateRequest struct {
	ID string
	UpdateRequestBody
}

type CreateRequest struct {
	UserName       string
	MembershipType string
}

func (c *CreateRequest) IsValid() bool {
	if c.UserName == "" || c.MembershipType == "" || !isValidPlatform(c.MembershipType) {
		return false
	}
	t := c.MembershipType
	return isValidPlatform(t)
}
func (u *UpdateRequestBody) ToServiceDto(id string) UpdateRequest {
	return UpdateRequest{
		ID:                id,
		UpdateRequestBody: *u,
	}
}

func (r *UpdateRequestBody) IsValid() bool {
	return r.UserName != "" && isValidPlatform(r.MembershipType)
}

func isValidPlatform(t string) bool {
	return t == naver || t == payco || t == toss
}
