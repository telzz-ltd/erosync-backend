package domain

import "time"

type UserStatus string
type UserRole string

var (
	UserStatusActive    UserStatus = "ACTIVE"
	UserStatusInactive  UserStatus = "INACTIVE"
	UserStatusSuspended UserStatus = "SUSPENDED"

	UserRoleRegular UserRole = "REGULAR"
	UserRoleStaff   UserRole = "STAFF"
	UserRoleAdmin   UserRole = "ADMIN"
)

type User struct {
	ID              string
	Name            string
	Email           string
	PasswordHash    string
	Status          UserStatus
	Role            UserRole
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
	EmailVerifiedAt *time.Time
}
