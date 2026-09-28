from pydantic import BaseModel, EmailStr, Field, ConfigDict
from domain import UserRole, UserStatus
import datetime as dt
from typing import Annotated


class RegisterRequest(BaseModel):
    name: str
    email: EmailStr
    password: str = Field(min_length=8, max_length=50)


class LoginRequest(BaseModel):
    email: EmailStr
    password: str = Field(min_length=8, max_length=50)


class VerifyEmailRequest(BaseModel):
    otp_code: str = Field(alias="code")

    model_config = ConfigDict(populate_by_name=True)


class ForgotPassword(BaseModel):
    email: EmailStr


class ResetPassword(BaseModel):
    email: Annotated[EmailStr, Field()]
    code: Annotated[str, Field(min_length=6, max_length=6)]
    password: Annotated[str, Field(min_length=8)]


class AuthResponse(BaseModel):
    user: "UserResponse"
    access_token: str = Field(alias="accessToken")
    refresh_token: str = Field(alias="refreshToken")

    model_config = ConfigDict(populate_by_name=True)


class UserResponse(BaseModel):
    id: str
    name: str
    email: str
    role: UserRole
    status: UserStatus
    created_at: dt.datetime
    updated_at: dt.datetime
    email_verified_at: dt.datetime | None = None

    model_config = ConfigDict(from_attributes=True)
