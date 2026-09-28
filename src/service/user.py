from uuid import uuid7

import bcrypt

from domain import User
from src.port.repository import UserRepository
from src.schema.user import RegisterRequest


class UserService:
    def __init__(self, store: UserRepository) -> None:
        self.store = store

    def create(self, dto: RegisterRequest) -> User:
        password_hash = bcrypt.hashpw(
            dto.password.encode(), bcrypt.gensalt()).decode()

        user = User.create(
            id=str(uuid7()), name=dto.name, email=dto.email, password_hash=password_hash
        )
        self.store.save(user)
        return user
