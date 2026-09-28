from src.port.repository import BrandRepository


class BrandService:
    def __init__(self, repo: BrandRepository) -> None:
        self.repo = repo
