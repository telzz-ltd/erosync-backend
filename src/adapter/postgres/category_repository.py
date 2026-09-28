from typing import Any

from psycopg_pool import ConnectionPool
from psycopg import sql

from src.domain.brand import Category


class CategoryRepository:
    def __init__(self, tablename: str, pool: ConnectionPool):
        self.pool = pool
        self.tablename = tablename

    def save(self, category: Category) -> None:
        with self.pool.connection() as conn:
            query = sql.SQL(
                "INSERT INTO {} (id, name, description) VALUES (%s, %s, %s);"
            ).format(sql.Identifier(self.tablename))

            conn.execute(
                query,
                (category.id, category.name, category.description),
            )

    def find(self, params: dict[str, Any]) -> list[Category]:
        where: list[str] = []
        args: list[Any] = []

        if isinstance(name := params.get("name"), str):
            args.append(f"%{name}%")
            where.append(f"name ILIKE %s")

        if isinstance(ids := params.get("ids"), list):
            args.append(ids)
            where.append("id = ANY(%s)")

        sql = f"""
            SELECT id, name, description
            FROM {self.tablename}
        """

        if where:
            sql += " WHERE " + " AND ".join(where)

        with self.pool.connection() as conn:
            rows = conn.execute(sql, args).fetchall()  # type: ignore

        return [
            Category(
                id=row[0],
                name=row[1],
                description=row[2],
            )
            for row in rows
        ]

    def find_by_id(self, id: str) -> Category:
        with self.pool.connection() as conn:
            row = conn.execute(
                f"""
                SELECT id, name, description
                FROM {self.tablename}
                WHERE id = %s;
                """,  # type: ignore
                (id,),
            ).fetchone()

        if row is None:
            raise LookupError(f"{self.tablename} category {id!r} not found")

        return Category(
            id=row[0],
            name=row[1],
            description=row[2],
        )

    def delete(self, ids: list[str]) -> None:
        with self.pool.connection() as conn:
            conn.execute(
                f"""
                DELETE FROM {self.tablename}
                WHERE id = ANY(%s)
                """,  # type: ignore
                (ids,),
            )

    def bulk_insert(self, categories: list[Category]) -> None:
        if not categories:
            return

        with self.pool.connection() as conn:
            with conn.cursor() as cur:
                with cur.copy(
                    f"""
                    COPY {self.tablename} (id, name, description)
                    FROM STDIN
                    """  # type: ignore
                ) as copy:
                    for category in categories:
                        copy.write_row(
                            (
                                category.id,
                                category.name,
                                category.description,
                            )
                        )


class ProductCategoryRepository(CategoryRepository):
    def __init__(self, pool: ConnectionPool):
        super().__init__("product_categories", pool)


class BrandCategoryRepository(CategoryRepository):
    def __init__(self, pool: ConnectionPool):
        super().__init__("brand_categories", pool)
