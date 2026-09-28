from dataclasses import dataclass
from enum import StrEnum

from .brand import Category


class ProductStatus(StrEnum):
    ACTIVE = "active"
    DRAFT = "draft"
    ARCHIVED = "archived"


@dataclass
class Product:
    id: str
    name: str
    category: Category
    description: str
    details: str
    price: int
    discount: int
    status: ProductStatus
