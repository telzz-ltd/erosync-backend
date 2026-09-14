from uuid import uuid7

import bcrypt

from .domain import User
from .protocols import UserStore
from .schema import RegisterUser


class UserService:
    def __init__(self, store: UserStore) -> None:
        self.store = store

    def create(self, dto: RegisterUser) -> User:
        password_hash = bcrypt.hashpw(dto.password.encode(), bcrypt.gensalt()).decode()

        user = User(
            id=str(uuid7()), name=dto.name, email=dto.email, password_hash=password_hash
        )
        self.store.save(user)
        return user
