package pgcommonmodel

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CreatedAt time.Time
	Email     string  `gorm:"email; not null;type:varchar(100);unique:true"`
	Roles     []*Role `gorm:"many2many:user_roles;constraint:OnDelete:CASCADE"`
}

type Role struct {
	ID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name  string    `gorm:"type:varchar(100);primaryKey;;constraint:OnDelete:CASCADE"`
	Users []*User   `gorm:"many2many:user_roles;constraint:OnDelete:CASCADE"`
}

func (u *User) GetRolesCount() int {
	return len(u.Roles)
}

func (u *User) GetRolesFromUser() []*Role {
	return u.Roles
}

func (r *Role) GetUsersFromRole() []*User {
	return r.Users
}
