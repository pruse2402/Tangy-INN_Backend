package models

// LoginDetail capture the login request
type LoginDetail struct {
	EmailID  string `json:"emailId"`
	Password string `json:"password"`
}
