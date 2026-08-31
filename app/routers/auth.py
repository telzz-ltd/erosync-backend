from fastapi import APIRouter
from schema.auth import RegisteRequest
from config.db import get_connection
from services import UserService
from store.users import UserStore


router = APIRouter(prefix="/auth")


@router.post("/register")
def register_handler(req: RegisteRequest):
    with get_connection() as conn:
        users = UserService(UserStore(conn))
        users.create(req)
    return req
