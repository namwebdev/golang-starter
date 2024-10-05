package persistent_objects

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"column:uuid;primary_key;type:char(36);not null"`
	Username string    `gorm:"column:username;type:varchar(100);not null"`
	IsActive bool      `gorm:"column:is_active;type:boolean;default:false"`
	Roles    []Role    `gorm:"many2many:user_roles"`
}

func (u *User) TableName() string {
	return "users"
}
