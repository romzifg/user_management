package rolemodule

import "time"

type UpdateRoleDto struct {
	RoleName  string `json:"role_name"`
	UpdatedAt time.Time `json:"updated_at" gorm:"timestamp"`
}