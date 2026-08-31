from psycopg import Connection
from models.user import User
from interfaces.user_repository import QueryUserResult, QueryUserParam


class UserStore():
    def __init__(self, conn: Connection) -> None:
        self.conn = conn

    def save(self, user: User) -> None:
        sql = """
                INSERT into users 
                    (id, name, email, password_hash, status, role, created_at, updated_at, email_verified_at)
                VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s)
                ON CONFLICT (id) DO UPDATE SET
                    name=EXCLUDED.name,
                    email=EXCLUDED.email,
                    password_hash=EXCLUDED.password_hash,
                    status=EXCLUDED.status,
                    role=EXCLUDED.role,
                    created_at=EXCLUDED.created_at,
                    updated_at=EXCLUDED.updated_at,
                    email_verified_at=EXCLUDED.email_verified_at
                ;
            """

    def findById(self, id: str) -> User | None:
        ...

    def findByEmail(self, email: str) -> User | None:
        ...

    def query(self, query: QueryUserParam) -> QueryUserResult:
        return QueryUserResult()
