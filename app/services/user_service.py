from uuid import uuid7
import bcrypt

from interfaces.user_repository import UserRepository
from models import User
from schema.auth import RegisteRequest


class UserService:
    def __init__(self, repo: UserRepository) -> None:
        self._repo = repo

    def create(self, req: RegisteRequest) -> User:
        password_hash = bcrypt.hashpw(req.password.encode(), bcrypt.gensalt())
        user = User(
            id=str(uuid7()),
            name=req.name,
            email=req.email,
            password_hash=str(password_hash)
        )

        self._repo.save(user)
        return user
