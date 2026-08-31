from fastapi import FastAPI

from routers import auth


app = FastAPI()


@app.get("/health")
def health():
    return {"message": "success"}


app.include_router(auth.router)
