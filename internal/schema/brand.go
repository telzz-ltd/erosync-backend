package schema

type CreateBrandRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=50"`
	Description string `json:"description" validate:"max=255"`
	LogoUrl     string `json:"logoUrl"`
	CategoryID  string `json:"categoryId" validate:"required"`
}

type UpdateBrandRequest CreateBrandRequest

type FindBrandQueryParams struct {
	Name string `json:"name"`
}

type CreateBrandCategoryRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=50"`
	Description string `json:"description" validate:"max=255"`
}
