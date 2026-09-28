from dataclasses import dataclass
import datetime as dt


class Inventory:
    product_id: str
    quantity: int
    low_stock_threshold: str
