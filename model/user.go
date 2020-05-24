package model

// RegistrationBody contains the body params for the /register endpoint
type RegistrationBody struct {
	PubKey string `json:"public_key"`
}

// RegistrationResponse contains the response for the /register endpoint
type RegistrationResponse struct {
	PubKey       string `json:"public_key"`
	ID           string `json:"id"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

// UserDataType is the model for rethinkdb directory table
type UserDataType struct {
	PubKey string `rethinkdb:"public_key" json:"public_key"`
	ID     string `rethinkdb:"id" json:"id"`
}
