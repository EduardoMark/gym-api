package user

type CreateUserParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     Role   `json:"role"`

	Phone   string `json:"phone"`
	Gender  string `json:"gender"`
	Address string `json:"address"`
}
