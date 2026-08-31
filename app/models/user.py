import datetime as dt
from enum import StrEnum


class UserStatus(StrEnum):
    ACTIVE = "ACTIVE"
    SUSPENDED = "SUSPENDED"
    BANNED = "BANNED"


class UserRole(StrEnum):
    USER = "USER"
    ADMIN = "ADMIN"
    MODERATOR = "MODERATOR"


class User:
    id: str
    name: str
    email: str
    password_hash: str
    status: UserStatus
    role: str
    created_at: dt.datetime
    updated_at: dt.datetime
    deleted_at: dt.datetime | None = None
    email_verified_at: dt.datetime | None = None

    def __init__(self, id: str, name: str, email: str, password_hash: str) -> None:
        if not id or not name or not email or not password_hash:
            raise ValueError("all fields must not be empty")

        self.id = id
        self.name = name
        self.email = email
        self.password_hash = password_hash
        self.status = UserStatus.ACTIVE
        self.role = UserRole.USER
        self.created_at = dt.datetime.now()
        self.updated_at = dt.datetime.now()

    def activate(self):
        self.status = UserStatus.ACTIVE
        self.updated_at = dt.datetime.now()

    def suspend(self):
        self.status = UserStatus.SUSPENDED
        self.updated_at = dt.datetime.now()

    def ban(self):
        self.status = UserStatus.BANNED
        self.updated_at = dt.datetime.now()

    def makeAdmin(self):
        self.role = UserRole.ADMIN
        self.updated_at = dt.datetime.now()

    def makeModerator(self):
        self.role = UserRole.MODERATOR
        self.updated_at = dt.datetime.now()

    def makeRegular(self):
        self.role = UserRole.USER
        self.updated_at = dt.datetime.now()

    def verifyEmail(self):
        self.email_verified_at = dt.datetime.now()
        self.updated_at = dt.datetime.now()

    def verified(self):
        return self.email_verified_at == None
