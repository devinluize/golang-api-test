package auth

type RegisterRequest struct {
	UserName   string `gorm:"column:user_name" json:"user_name"`
	UserEmail  string `gorm:"column:user_email" json:"user_email"`
	Password   string `gorm:"column:password" json:"password"`
	UserRoleId int    `gorm:"column:user_role" json:"user_role_id"`
}

func (RegisterRequest) TableName() string {
	return "api.GoTestUsers"
}

type LoginRequest struct {
	UserName string `gorm:"column:user_name" json:"username"`
	Password string `gorm:"column:password" json:"password"`
}
type AuthResponses struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
