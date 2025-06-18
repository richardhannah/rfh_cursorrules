package dto

type UserDTO struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password,omitempty"`
	Salt      string `json:"salt,omitempty"`
	Ipaddress string `json:"ipaddress,omitempty"`
	Enabled   bool   `json:"enabled"`
	Role      string `json:"role,omitempty"`
}
