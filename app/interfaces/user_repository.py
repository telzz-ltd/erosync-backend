from typing import Protocol

from models.user import User


class QueryUserParam:
    id: str
    name: str
    email: str


class QueryUserResult:
    users: list[User]
    count: int


class UserRepository(Protocol):
    def save(self, user: User):
        ...

    def findById(self, id: str) -> User | None:
        ...

    def findByEmail(self, email: str) -> User | None:
        ...

    def query(self, query: QueryUserParam) -> QueryUserResult:
        ...
