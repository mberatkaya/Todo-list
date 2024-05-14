package user

type CreateUserDto struct {
	Nickname string `json:"nickname"`
	FullName string `json:"fullName"`
	Password string `json:"password"`
}

type UpdateUserDto struct {
	Nickname *string `json:"nickname,omitempty"`
	FullName *string `json:"fullName,omitempty"`
	Password *string `json:"password,omitempty"`
}
