from pydantic import BaseModel, EmailStr, Field
from typing import Annotated


class RegisteRequest(BaseModel):
    name: Annotated[str, Field()]
    email: Annotated[EmailStr, Field()]
    password: Annotated[str, Field(min_length=8, max_length=50)]


class LoginRequest(BaseModel):
    email: Annotated[EmailStr, Field()]
    password: Annotated[EmailStr, Field(min_length=8, max_length=50)]


class VerifyEmailRequest(BaseModel):
    code: Annotated[str, Field(max_length=6, min_length=6)]
