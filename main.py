from fastapi import FastAPI
from fastapi.concurrency import asynccontextmanager

from src.lib.db import pool
from src.router import user_router


@asynccontextmanager
async def lifespan(instance: FastAPI):
    pool.open()
    yield
    pool.close()


app = FastAPI(lifespan=lifespan)


@app.get("/health")
def health_check():
    return {"message": "App working fine"}


app.mount("/", user_router)
