from fastapi import APIRouter, HTTPException, Response
from psycopg.errors import UniqueViolation

import bcrypt

from internal.lib import jwt
from internal.lib.db import pool
from internal.pg_store.user_store import PGUserStore

from .schema import ForgotPassword, Login, RegisterUser, ResetPassword, UserResponse
from .service import UserService
from internal.lib.email import load_template, send_mail

router = APIRouter()

store = PGUserStore(pool)
service = UserService(store)


@router.post("/auth/register")
def register(dto: RegisterUser):
    try:
        user = service.create(dto)
        access_token = jwt.encode({"sub": user.id, "role": user.role})
        return {"accessToken": access_token, "user": UserResponse(user)}
    except UniqueViolation as e:
        raise HTTPException(400, {"message": e.args[0]})


@router.post("/auth/login")
def login(dto: Login):
    user = store.find_by_email(dto.email)
    if user is None:
        raise HTTPException(400, {"message": "invalid credentials"})

    if not bcrypt.checkpw(dto.password.encode(), user.password_hash.encode()):
        raise HTTPException(400, {"message": "invalid credentials"})

    access_token = jwt.encode({"sub": user.id, "role": user.role})
    return {"accessToken": access_token, "user": UserResponse(user)}


@router.post("/auth/forgot-password")
def forgot_password(dto: ForgotPassword):
    return {"data": dto}


@router.post("/auth/reset-password")
def reset_password(dto: ResetPassword):
    return {"data": dto}


@router.post("/verification/email/send-otp", )
def send_email_verification_code():
    send_mail(
        template="welcome",
        to="baba@test.com",
        subject="Welcome to Erosync",
        template_args={
            "app_url": "http://localhost:8080",
            "app_name": "Erosync LTD.",
            "name": "Usman",
            "expire_min": 6,
            "year": 2026,
            "code": "556770"
        }
    )
    return {"message": "email sent"}


@router.post("/verification/email/verify")
def verify_email():
    return {"message", "email verified"}
