from fastapi import Depends
from typing import Annotated

from src.service import UserService
from src.adapter.postgres import UserRepository, BrandRepository, OtpRepository

UserDeps = Annotated[UserService, Depends(UserRepository)]
