package user

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `string:"password"`
}
