from dataclasses import dataclass
from typing import Any
import datetime as dt


@dataclass
class Category:
    id: str
    name: str
    description: str


@dataclass
class Brand:
    id: str
    name: str
    description: str
    logo_url: str
    contact_info: dict
    created_at: dt.datetime
    categories: list[Category]
