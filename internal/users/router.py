from fastapi import APIRouter
from psycopg_pool import ConnectionPool

from internal.lib import jwt
from internal.pg_store.users import PGUserStore

from .schema import ForgotPassword, Login, RegisterUser, ResetPassword
from .service import UserService

router = APIRouter()


@router.post("/auth/register")
def register(dto: RegisterUser):
    with ConnectionPool() as pool:
        store = PGUserStore(pool)
        service = UserService(store)

        user = service.create(dto)
        access_token = jwt.encode({"sub": user.id, "role": user.role})
        return {"accessToken": access_token}


@router.post("/auth/login")
def login(dto: Login):
    return {"data": dto}


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
