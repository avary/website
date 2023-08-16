package model

type UserInfoResponse struct {
	Email    string `json:"email"`
	Username string `json:"nickname"` // Assuming 'nickname' is the Auth0's username
}

type AccountDetails struct {
	Name            string
	Country         string
	ShippingAddress string
}
