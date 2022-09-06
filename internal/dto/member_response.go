package dto

type CreateResponse struct {
	ID             string `json:"ID"`
	MembershipType string `json:"MembershipType"`
}

type UpdateResponse struct {
	ID             string `json:"ID"`
	UserName       string `json:"UserName"`
	MembershipType string `json:"MembershipType"`
}

type GetResponse struct {
	ID             string `json:"ID"`
	UserName       string `json:"UserName"`
	MembershipType string `json:"MembershipType"`
}

func NewCreateResponse(id, membershipType string) CreateResponse {
	return CreateResponse{
		ID:             id,
		MembershipType: membershipType,
	}
}

func NewUpdateResponse(id, username, membershipType string) UpdateResponse {
	return UpdateResponse{
		ID:             id,
		UserName:       username,
		MembershipType: membershipType,
	}
}

func NewGetResponse(id, username, membershipType string) GetResponse {
	return GetResponse{
		ID:             id,
		UserName:       username,
		MembershipType: membershipType,
	}
}
