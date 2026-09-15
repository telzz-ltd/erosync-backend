from psycopg.rows import class_row
from psycopg_pool import ConnectionPool

from internal.users.domain import User
from internal.users.protocols import FindUsersResult


class PGUserStore:
    def __init__(self, pool: ConnectionPool) -> None:
        self.pool = pool

    def save(self, user: User):
        with self.pool.connection() as conn, conn.cursor() as cur:
            cur.execute(
                """
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
                """,
                (
                    user.id,
                    user.name,
                    user.email,
                    user.password_hash,
                    user.status,
                    user.role,
                    user.created_at,
                    user.updated_at,
                    user.email_verified_at,
                ),
            )

    def delete_by_id(self, id: str):
        raise SyntaxError("method not implemented")

    def find_by_email(self, email: str) -> User | None:
        with (
            self.pool.connection() as conn,
            conn.cursor(row_factory=class_row(User)) as cur,
        ):
            return cur.execute(
                "SELECT * FROM users WHERE email = %s LIMIT 1;", (email,)
            ).fetchone()

    def find_by_id(self, id: str) -> User:
        raise SyntaxError("method not implemented")

    def find(self, query: dict) -> FindUsersResult:
        raise SyntaxError("method not implemented")
