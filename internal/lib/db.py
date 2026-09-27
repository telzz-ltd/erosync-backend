import os
from psycopg_pool import ConnectionPool

pool: ConnectionPool = ConnectionPool(os.getenv("DATABASE_URL", ""), open=False)

