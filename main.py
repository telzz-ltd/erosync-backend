from fastapi import FastAPI

from internal.users import user_router

app = FastAPI()


@app.get("/health")
def health_check():
    return {"message": "App working fine"}


app.mount("/", user_router)
