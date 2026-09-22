import os

from psycopg_pool import ConnectionPool

pool: ConnectionPool = ConnectionPool(os.getenv("DB_URL", ""), open=False)

