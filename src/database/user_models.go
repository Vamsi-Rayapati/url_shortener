package database

import (
	"time"

	"github.com/google/uuid"
)

type UserRole string

type UserRolesStruct struct {
	User       UserRole
	Supervisor UserRole
}

var UserRoles = UserRolesStruct{
	User:       "USER",
	Supervisor: "SUPERVISOR",
}

type UserStatus string

const (
	Active UserStatus = "ACTIVE"

	InActive UserStatus = "INACTIVE"
)

type User struct {
	ID             uuid.UUID  `gorm:"type:char(36);primaryKey;default:(UUID())"`
	Username       string     `gorm:"type:varchar(255);not null;unique"`
	FirstName      string     `gorm:"type:varchar(255);not null"`
	LastName       string     `gorm:"type:varchar(255);not null"`
	PrimaryAddress string     `gorm:"type:varchar(255)"`
	Avatar         string     `gorm:"type:varchar(255)"`
	Mobile         string     `gorm:"type:varchar(15)"`
	Role           UserRole   `gorm:"type:enum('USER', 'SUPERVISOR');not null;default:'USER'"`
	Status         UserStatus `gorm:"type:enum('ACTIVE', 'INACTIVE');not null;default:'INACTIVE'"`
	CreatedAt      time.Time  `gorm:"autoCreateTime"`
	UpdatedAt      time.Time  `gorm:"autoUpdateTime"`
}
