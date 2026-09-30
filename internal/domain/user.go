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
	ID              string     `json:"id"`
	Name            string     `json:"name"`
	Email           string     `json:"email"`
	PasswordHash    string     `json:"-"`
	Status          UserStatus `json:"status"`
	Role            UserRole   `json:"role"`
	CreatedAt       time.Time  `json:"createdAt"`
	UpdatedAt       time.Time  `json:"updatedAt"`
	DeletedAt       *time.Time `json:"deletedAt"`
	EmailVerifiedAt *time.Time `json:"emailVerifiedAt"`
}

func NewUser(id, name, email, passwordHash string) User {
	return User{
		ID:           id,
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
		Status:       UserStatusActive,
		Role:         UserRoleRegular,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}
