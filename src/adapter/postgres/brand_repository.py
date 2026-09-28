from dataclasses import dataclass
from psycopg_pool import ConnectionPool
from src.domain.brand import Brand, Category
from psycopg.rows import dict_row, class_row


@dataclass
class BrandRepository:
    pool: ConnectionPool

    def save(self, brand: Brand):
        with self.pool.connection() as conn, conn.cursor() as cur:
            cur.execute(
                """
                INSERT INTO brands (id, name, description, logo_url, contact_info, created_at)
		        VALUES (%s, %s, %s, %s, %s, %s)
		        ON CONFLICT (id) DO UPDATE SET
			        name=EXCLUDED.name,
			        description=EXCLUDED.description,
                    logo_url=EXCLUDED.logo_url,
                    contact_info=EXCLUDED.contact_info;
                """,
                (
                    brand.id, brand.name, brand.description,
                    brand.logo_url, brand.contact_info, brand.created_at,
                )
            )

            cur.execute(
                "DELETE FROM brand_category_pivot WHERE brand_id = ?", (brand.id,))

            cur.executemany(
                "INSERT INTO brand_category_pivot (brand_id, category_id) VALUES (%s, %s);",
                [(brand.id, cat.id) for cat in brand.categories],
            )

            conn.commit()

    def find_by_id(self, id: str) -> Brand | None:
        with self.pool.connection() as conn, conn.cursor(row_factory=dict_row) as cur:
            cur.execute("SELECT * FROM brands WHERE id = %s;", (id,))
            result = cur.fetchone()

            if result is not None:
                brand = Brand(**result)

                cur.execute(
                    """
                    SELECT * FROM brand_categories bc JOIN brand_category_pivot cp
                    ON bc.id = cp.category_id AND cp.brand_id = %s;
                    """,
                    (id,)
                )

                result = cur.fetchall()
                brand.categories = [Category(**item) for item in result]

                return brand

    def exist_by_name(self, name: str) -> bool:
        with self.pool.connection() as conn, conn.cursor() as cur:
            cur.execute(
                "SELECT count(*) FROM brands WHERE name ILIKE %s;",
                (f"%{name}%",),
            )

            result = cur.fetchone()
            if result is not None:
                return result[0] > 0

            return False

    def delete(self, ids: tuple[str]):
        with self.pool.connection() as conn, conn.cursor() as cur:
            cur.execute("DELETE FROM brands WHERE id IN(%s)", (ids,))

    def find(self, params: dict = {}) -> list[Brand]:
        with self.pool.connection() as conn, conn.cursor(row_factory=class_row(Brand)) as cur:
            cur.execute("SELECT * FROM brands;")
            return cur.fetchall()
