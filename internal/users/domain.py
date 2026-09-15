import datetime as dt

from .constants import Role, Status


class User:
    id: str
    name: str
    email: str
    password_hash: str
    role: Role
    status: Status
    created_at: dt.datetime
    updated_at: dt.datetime
    email_verified_at: dt.datetime | None = None

    def __init__(
        self,
        id: str,
        name: str,
        email: str,
        password_hash: str,
        role: Role = Role.USER,
        status: Status = Status.ACTIVE,
        created_at: dt.datetime | None = None,
        updated_at: dt.datetime | None = None,
        email_verified_at: dt.datetime | None = None,
    ) -> None:
        self.id = id
        self.name = name
        self.email = email
        self.password_hash = password_hash
        self.status = status
        self.role = role

        now = dt.datetime.now(dt.timezone.utc)
        self.created_at = created_at if created_at is not None else now
        self.updated_at = updated_at if updated_at is not None else now

    @classmethod
    def create(cls, id: str, name: str, email: str, password_hash: str):
        if not id or not name or not email or not password_hash:
            raise ValueError("all fields are required")

        return cls(
            id=id,
            name=name,
            email=email,
            password_hash=password_hash,
        )

    def change_password(self, password_hash: str):
        if password_hash:
            raise ValueError("password_hash must not be empty")
        self.password_hash = password_hash
        self.touch()

    def update_info(self, name: str, email: str):
        if not name or not email:
            raise ValueError("all fields are required")

        self.name = name
        self.email = email
        self.touch()

    def make_admin(self):
        self.role = Role.ADMIN
        self.touch()

    def touch(self):
        self.updated_at = dt.datetime.now(dt.timezone.utc)

    def deactivate(self):
        self.status = Status.INACTIVE
        self.touch()

    def activate(self):
        self.status = Status.ACTIVE
        self.touch()

    def suspend(self):
        self.status = Status.SUSPENDED
        self.touch()

    def make_moderator(self):
        self.role = Role.MODERATOR
        self.touch()

    def make_user(self):
        self.role = Role.USER
        self.touch()

    def email_verified(self) -> bool:
        return self.email_verified_at != None

    def verify_email(self):
        self.email_verified_at = dt.datetime.now(dt.timezone.utc)
        self.touch()
