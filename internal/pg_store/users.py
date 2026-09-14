from typing import Any

from psycopg_pool import ConnectionPool

from internal.users.domain import User
from internal.users.protocols import FindUsersResult


class PGUserStore:
    def __init__(self, pool: ConnectionPool[Any]) -> None:
        pass

    def save(self, user: User):
        raise SyntaxError("method not implemented")

    def delete_by_id(self, id: str):
        raise SyntaxError("method not implemented")

    def find_by_email(self, email: str) -> User:
        raise SyntaxError("method not implemented")

    def find_by_id(self, id: str) -> User:
        raise SyntaxError("method not implemented")

    def find(self, query: dict) -> FindUsersResult:
        raise SyntaxError("method not implemented")
