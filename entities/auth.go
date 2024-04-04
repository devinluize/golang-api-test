package entities

type User struct {
	UserName   string `gorm:"column:user_name" json:"user_name"`
	UserEmail  string `gorm:"column:user_email" json:"user_email"`
	Password   string `gorm:"column:password" json:"password"`
	UserRoleId int    `gorm:"column:user_role" json:"user_role_id"`
}

func (User) TableName() string {
	return "api.GoTestUsers"
}
