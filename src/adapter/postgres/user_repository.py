from psycopg_pool import ConnectionPool
from psycopg.rows import class_row

from dataclasses import dataclass

from src.domain.user import User


@dataclass
class UserRepository:
    pool: ConnectionPool

    def save(self, user: User):
        with self.pool.connection() as conn:
            conn.execute("""
                INSERT into users (id, name, email, password_hash, status, role, created_at, updated_at, email_verified_at)
			    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
			    ON CONFLICT (id) DO UPDATE SET
                    name=EXCLUDED.name,
                    email=EXCLUDED.email,
                    password_hash=EXCLUDED.password_hash,
                    status=EXCLUDED.status,
                    role=EXCLUDED.role,
                    created_at=EXCLUDED.created_at,
                    updated_at=EXCLUDED.updated_at,
                    email_verified_at=EXCLUDED.email_verified_at;
            """,
                         (user.id, user.name, user.email, user.password_hash, user.status,
                          user.role, user.created_at, user.updated_at, user.email_verified_at)
                         )

    def find_by_email(self, email: str) -> User | None:
        with self.pool.connection() as conn, conn.cursor(row_factory=class_row(User)) as cur:
            cur.execute("SELECT * FROM users WHERE email = ?", (email, ))
            return cur.fetchone()

    def find_by_id(self, id: str) -> User | None:
        with self.pool.connection() as conn, conn.cursor(row_factory=class_row(User)) as cur:
            cur.execute("SELECT * FROM users WHERE id = ?", (id, ))
            return cur.fetchone()
