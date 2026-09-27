package domain

type ProductStatus string

var (
	ProductStatusActive   ProductStatus = "active"
	ProductStatusDraft    ProductStatus = "draft"
	ProductStatusArchived ProductStatus = "archived"
)

type Product struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Category    ProductCategory `json:"category"`
	Description string          `json:"description"`
	Details     string          `json:"details"`
	Price       int             `json:"price"`
	Discount    int             `json:"discount"`
	Status      ProductStatus   `json:"status"`
}
