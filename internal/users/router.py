from fastapi import APIRouter, HTTPException

from internal.lib import jwt
from internal.lib.db import pool
from internal.pg_store.user_store import PGUserStore

from .schema import ForgotPassword, Login, RegisterUser, ResetPassword, UserResponse
from .service import UserService

router = APIRouter()

store = PGUserStore(pool)
service = UserService(store)


@router.post("/auth/register")
def register(dto: RegisterUser):
    user = service.create(dto)
    access_token = jwt.encode({"sub": user.id, "role": user.role})
    return {"accessToken": access_token, "user": UserResponse(user)}


@router.post("/auth/login")
def login(dto: Login):
    user = store.find_by_email(dto.email)
    if user is None:
        raise HTTPException(400, {"message": "user not found"})

    access_token = jwt.encode({"sub": user.id, "role": user.role})
    return {"accessToken": access_token, "user": UserResponse(user)}


@router.post("/auth/forgot-password")
def forgot_password(dto: ForgotPassword):
    return {"data": dto}


@router.post("/auth/reset-password")
def reset_password(dto: ResetPassword):
    return {"data": dto}


@router.post("/verification/email/send-otp")
def send_email_verification_code():
    return {"message": "email sent"}


@router.post("/verification/email/verify")
def verify_email():
    return {"message", "email verified"}
