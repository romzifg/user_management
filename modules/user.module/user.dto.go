package usermodule

type CreateUserDataDto struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	RoleId    int64  `json:"role_id" binding:"required"`
}

type DataUserDto struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required"`
	RoleId    int64  `json:"role_id" binding:"required"`
}