package persistent_objects

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID       uuid.UUID `gorm:"column:id;primary_key;type:char(36);auto_increment"`
	RoleName string    `gorm:"column:role_name;type:varchar(100);not null"`
	RoleNote string    `gorm:"column:role_note;type:varchar(100);not null"`
}

func (r *Role) TableName() string {
	return "roles"
}
