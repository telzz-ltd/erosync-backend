package domain

import (
	"encoding/json"
	"errors"
	"strings"
	"time"
)

type Brand struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	LogoUrl     *string          `json:"logoUrl"`
	ContactInfo *json.RawMessage `json:"contactInfo"`
	CreatedAt   time.Time        `json:"createdAt"`
	Categories  []BrandCategory  `json:"categories,omitempty" db:"-"`
}

func NewBrand(id, name string) (Brand, error) {
	if strings.TrimSpace(id) == "" || strings.TrimSpace(name) == "" {
		return Brand{}, errors.New("all fields are required")
	}

	return Brand{
		ID:        id,
		Name:      name,
		CreatedAt: time.Now(),
	}, nil
}

func (b *Brand) SetDescription(description string) {
	b.Description = description
}

func (b *Brand) SetLogoUrl(url string) {
	b.LogoUrl = new(url)
}

func (b *Brand) SetContactInfo(contactInfo json.RawMessage) {
	b.ContactInfo = &contactInfo
}
