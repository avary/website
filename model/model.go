package model

type UserInfoResponse struct {
	Email    string `json:"email"`
	Username string `json:"nickname"` // Assuming 'nickname' is the Auth0's username
}

type AccountDetails struct {
	Name            string
	Username        string
	Email           string
	Country         string
	ShippingAddress string
}
