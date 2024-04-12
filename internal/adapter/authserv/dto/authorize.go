package dto

const authorizeUrl = "/api/v1/auth/validate-token"

type AuthorizationRequest struct {
	Token string `json:"token"`
}

type AuthorizationResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Role  string `json:"role"`
}

func (AuthorizationRequest) URL() string {
	return authorizeUrl
}
