package models

type User struct {
	ID       string `json:"id,omitempty" yaml:"Id,omitempty"`
	Name     string `json:"name,omitempty" yaml:"Name,omitempty"`
	Email    string `json:"email,omitempty" yaml:"Email,omitempty"`
	Password string `json:"password" yaml:"Password"`
}
