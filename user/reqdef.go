package user

type CreateUserRequest struct {
	Nickname string `json:"nickname"`
	FullName string `json:"fullName"`
	Password string `json:"password"`
}
