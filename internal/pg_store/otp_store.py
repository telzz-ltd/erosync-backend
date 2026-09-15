from dataclasses import dataclass

from psycopg_pool import ConnectionPool

from internal.otps.domain import OTP


@dataclass
class PGOTPStore:
    pool: ConnectionPool

    def save(self, otp: OTP):
        raise ValueError("method not implemented")
