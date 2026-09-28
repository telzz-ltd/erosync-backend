from dataclasses import dataclass
from psycopg_pool import ConnectionPool
from psycopg.rows import class_row
from src.domain.otp import OTP


@dataclass
class OtpRepository:
    pool: ConnectionPool

    def save(self, otp: OTP):
        with self.pool.connection() as conn:
            conn.execute("""
                INSERT into otps (recipient, code_hash, purpose, channel, attempts, max_attempts, created_at, expires_at)
			    VALUES (?, ?, ?, ?, ?, ?, ?, ?)
			    ON CONFLICT (recipient, purpose, channel) DO UPDATE SET
                    code_hash=EXCLUDED.code_hash,
                    attempts=EXCLUDED.attempts,
                    max_attempts=EXCLUDED.max_attempts,
                    expires_at=EXCLUDED.expires_at;
                """,
                         (otp.recipient, otp.code_hash, otp.purpose, otp.channel,
                          otp.attempts, otp.max_attempts, otp.created_at, otp.expires_at)
                         )

    def find_one(self, recipient: str, purpose: str, channel: str) -> OTP | None:
        with self.pool.connection() as conn, conn.cursor(row_factory=class_row(OTP)) as cur:
            cur.execute(
                "SELECT * FROM otps WHERE recipient = ? AND channel = ? AND purpose = ? LIMIT 1;",
                (recipient, channel, purpose)
            )
            return cur.fetchone()

    def delete(self, otp: OTP):
        with self.pool.connection() as conn:
            conn.execute(
                "DELETE FROM otps WHERE recipient = ? AND channel = ? AND purpose = ?;",
                (otp.recipient, otp.channel, otp.purpose))
