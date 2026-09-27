package domain

import (
	"errors"
	"net/mail"
	"regexp"
	"strings"
	"time"
)

type UserStatus string
type UserRole string

var (
	UserStatusActive    UserStatus = "ACTIVE"
	UserStatusInactive  UserStatus = "INACTIVE"
	UserStatusSuspended UserStatus = "SUSPENDED"

	UserRoleUser      UserRole = "USER"
	UserRoleAdmin     UserRole = "ADMIN"
	UserRoleModerator UserRole = "MODERATOR"
)

type User struct {
	ID              string     `json:"id" db:"id"`
	Name            string     `json:"name" db:"name"`
	Email           string     `json:"email" db:"email"`
	PasswordHash    string     `json:"-" db:"password_hash"`
	Role            UserRole   `json:"role" db:"role"`
	Status          UserStatus `json:"status" db:"status"`
	CreatedAt       time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt       time.Time  `json:"updatedAt" db:"updated_at"`
	EmailVerifiedAt *time.Time `json:"emailVerifiedAt" db:"email_verified_at"`
}

func NewUser(id, name, email, passwordHash string) (User, error) {
	id = strings.TrimSpace(id)
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)
	passwordHash = strings.TrimSpace(passwordHash)

	if id == "" || name == "" || email == "" || passwordHash == "" {
		return User{}, errors.New("all fields must not be empty")
	}

	if len(passwordHash) < 8 {
		return User{}, errors.New("password must be 8 or more chars")
	}

	emailAddr, err := mail.ParseAddress(email)
	if err != nil {
		return User{}, err
	}

	matched, err := regexp.MatchString("^[a-zA-Z]{3,}(?: [a-zA-Z]{3,}){1,2}$", name)
	if !matched || err != nil {
		return User{}, errors.New("invalid name")
	}

	return User{
		ID:           id,
		Name:         name,
		Email:        emailAddr.Address,
		PasswordHash: passwordHash,
		Role:         UserRoleUser,
		Status:       UserStatusActive,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, nil
}

func (u *User) MakeAdmin() {
	u.Role = UserRoleAdmin
	u.UpdatedAt = time.Now()
}

func (u *User) Deactivate() {
	u.Status = UserStatusInactive
	u.UpdatedAt = time.Now()
}

func (u *User) Activate() {
	u.Status = UserStatusActive
	u.UpdatedAt = time.Now()
}

func (u *User) Suspend() {
	u.Status = UserStatusSuspended
	u.UpdatedAt = time.Now()
}

func (u *User) SetModerator() {
	u.Role = UserRoleModerator
	u.UpdatedAt = time.Now()
}

func (u *User) MakeModerator() {
	u.Role = UserRoleModerator
	u.UpdatedAt = time.Now()
}

func (u *User) MakeUser() {
	u.Role = UserRoleUser
	u.UpdatedAt = time.Now()
}

func (u *User) VerifyEmail() {
	u.EmailVerifiedAt = new(time.Now())
	u.UpdatedAt = time.Now()
}

func (u *User) EmailVerified() bool {
	return u.EmailVerifiedAt != nil
}
