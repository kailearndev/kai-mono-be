package user

type RegisterUserDTO struct {
	Username  string `json:"username" binding:"required,min=3,max=50"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
	FullName  string `json:"name" binding:"required,min=2,max=100"`
	Role      string `json:"role" binding:"omitempty,oneof=admin user"`
	AvatarURL string `json:"avatar_url" binding:"omitempty,url"`
}

type LoginUserDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
