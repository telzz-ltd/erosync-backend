package schema

type CreateBrandRequest struct {
	Name        string   `json:"name" validate:"required,min=3,max=50"`
	Description string   `json:"description" validate:"max=255"`
	LogoUrl     string   `json:"logoUrl"`
	CategoryIds []string `json:"categoryIds" validate:"required,min=1,unique,dive,required"`
}

type UpdateBrandRequest CreateBrandRequest

type FindBrandQueryParams struct {
	Name string `json:"name"`
}

type CreateBrandCategoryRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=50"`
	Description string `json:"description" validate:"max=255"`
}

type BulkCreateBrandCategoriesRequest struct {
	Data []CreateBrandCategoryRequest `json:"data" validate:"required,min=1,unique"`
}
