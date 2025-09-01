package types

type CreateUser struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUser struct {
	UserId string `json:"userId"`
	UserName string `json:"userName"`
}

type UpdateUserRequest struct {
	UserName string `json:"userName"`
}