from fastapi import APIRouter

from .schema import ForgotPassword, Login, RegisterUser, ResetPassword

router = APIRouter()


@router.post("/auth/register")
def register(dto: RegisterUser):
    return {"data": dto}


@router.post("/auth/login")
def login(dto: Login):
    return {"data": dto}


@router.post("/auth/forgot-password")
def forgot_password(dto: ForgotPassword):
    return {"data": dto}


@router.post("/auth/reset-password")
def reset_password(dto: ResetPassword):
    return {"data": dto}
