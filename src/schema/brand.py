from pydantic import BaseModel, Field, field_validator


class CreateBrandRequest(BaseModel):
    name: str = Field(min_length=3, max_length=50)
    description: str = Field(default="", max_length=255)
    logo_url: str = ""
    category_ids: list[str] = Field(min_length=1)

    @field_validator("category_ids")
    @classmethod
    def category_ids_must_be_unique(cls, value: list[str]) -> list[str]:
        if len(value) != len(set(value)):
            raise ValueError("category_ids must contain unique values")
        return value


class UpdateBrandRequest(CreateBrandRequest):
    pass


class FindBrandQueryParams(BaseModel):
    name: str = ""


class CreateCategoryRequest(BaseModel):
    name: str = Field(min_length=3, max_length=50)
    description: str = Field(default="", max_length=255)


class BulkCreateCategoriesRequest(BaseModel):
    data: list[CreateCategoryRequest] = Field(min_length=1)

    @field_validator("data")
    @classmethod
    def data_must_be_unique(cls, value: list[CreateCategoryRequest]):
        names = [item.name for item in value]

        if len(names) != len(set(names)):
            raise ValueError("data must contain unique categories")

        return value
