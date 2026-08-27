package model

import (
	"errors"
	"net/mail"
	"regexp"
	"strings"
	"time"
)

type Status string
type Role string

var (
	StatusActive    Status = "ACTIVE"
	StatusInactive  Status = "INACTIVE"
	StatusSuspended Status = "SUSPENDED"

	RoleUser      Role = "USER"
	RoleAdmin     Role = "ADMIN"
	RoleModerator Role = "MODERATOR"
)

type User struct {
	ID              string     `json:"id" db:"id"`
	Name            string     `json:"name" db:"name"`
	Email           string     `json:"email" db:"email"`
	PasswordHash    string     `json:"-" db:"password_hash"`
	Role            Role       `json:"role" db:"role"`
	Status          Status     `json:"status" db:"status"`
	CreatedAt       time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt       time.Time  `json:"updatedAt" db:"updated_at"`
	EmailVerifiedAt *time.Time `json:"emailVerifiedAt" db:"email_verified_at"`
}

func NewUser(id, name, email, passwordHash string) (*User, error) {
	id = strings.TrimSpace(id)
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)
	passwordHash = strings.TrimSpace(passwordHash)

	if id == "" || name == "" || email == "" || passwordHash == "" {
		return nil, errors.New("all fields must not be empty")
	}

	if len(passwordHash) < 8 {
		return nil, errors.New("password must be 8 or more chars")
	}

	emailAddr, err := mail.ParseAddress(email)
	if err != nil {
		return nil, err
	}

	matched, err := regexp.MatchString("^[a-zA-Z]{3,}(?: [a-zA-Z]{3,}){1,2}$", name)
	if !matched || err != nil {
		return nil, errors.New("invalid name")
	}

	return &User{
		ID:           id,
		Name:         name,
		Email:        emailAddr.Address,
		PasswordHash: passwordHash,
		Role:         RoleUser,
		Status:       StatusActive,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, nil
}

func (u *User) MakeAdmin() {
	u.Role = RoleAdmin
	u.UpdatedAt = time.Now()
}

func (u *User) Deactivate() {
	u.Status = StatusInactive
	u.UpdatedAt = time.Now()
}

func (u *User) Activate() {
	u.Status = StatusActive
	u.UpdatedAt = time.Now()
}

func (u *User) Suspend() {
	u.Status = StatusSuspended
	u.UpdatedAt = time.Now()
}

func (u *User) SetModerator() {
	u.Role = RoleModerator
	u.UpdatedAt = time.Now()
}

func (u *User) MakeModerator() {
	u.Role = RoleModerator
	u.UpdatedAt = time.Now()
}

func (u *User) MakeUser() {
	u.Role = RoleUser
	u.UpdatedAt = time.Now()
}

func (u *User) VerifyEmail() {
	u.EmailVerifiedAt = new(time.Now())
	u.UpdatedAt = time.Now()
}

func (u *User) EmailVerified() bool {
	return u.EmailVerifiedAt != nil
}
