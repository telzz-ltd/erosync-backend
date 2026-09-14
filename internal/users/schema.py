from typing import Annotated

from pydantic import BaseModel, EmailStr, Field


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
