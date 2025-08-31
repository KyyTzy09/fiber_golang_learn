package types

type CreateUser struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
