from psycopg import Connection, connect
from os import getenv


def get_connection() -> Connection:
    [host, port, user, password] = [
        getenv("DB_HOST"),
        getenv("DB_PORT"),
        getenv("DB_USER"),
        getenv("DB_PASSWORD")
    ]

    return connect(f"host={host} port={port} user={user} password={password} sslmode=disable")
