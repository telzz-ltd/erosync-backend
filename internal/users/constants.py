from enum import StrEnum


class Status(StrEnum):
    ACTIVE = "ACTIVE"
    INACTIVE = "INACTIVE"
    SUSPENDED = "SUSPENDED"


class Role(StrEnum):
    USER = "USER"
    ADMIN = "ADMIN"
    MODERATOR = "MODERATOR"
