from typing import Annotated

from pydantic import BaseModel, EmailStr, Field

from .domain import User


class RegisterUser(BaseModel):
    name: Annotated[str, Field(min_length=8)]
    email: Annotated[EmailStr, Field()]
    password: Annotated[str, Field(min_length=8)]


class Login(BaseModel):
    email: Annotated[EmailStr, Field()]
    password: Annotated[str, Field()]


class ForgotPassword(BaseModel):
    email: Annotated[EmailStr, Field()]


class ResetPassword(BaseModel):
    email: Annotated[EmailStr, Field()]
    code: Annotated[str, Field(min_length=6, max_length=6)]
    password: Annotated[str, Field(min_length=8)]


class UserResponse:
    def __init__(self, user: User) -> None:
        self.id = user.id
        self.name = user.name
        self.email = user.email
        self.status = user.status
        self.role = user.role
        self.createdAt = user.created_at
        self.updatedAt = user.updated_at
        self.emailVerifiedAt = user.email_verified_at
