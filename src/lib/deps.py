from fastapi import Depends
from psycopg_pool import ConnectionPool
import os

from src.adapter.postgres import UserRepository, BrandRepository, OtpRepository

pool = ConnectionPool(os.getenv("DATABASE_URL", ""), open=False)


def get_db():
    with pool.connection() as conn:
        yield conn


def get_user_repo(pool: ConnectionPool = Depends(get_db)):
    return UserRepository(pool)


def get_otp_repo(pool: ConnectionPool = Depends(get_db)):
    return OtpRepository(pool)
