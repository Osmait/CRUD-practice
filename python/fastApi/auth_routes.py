from datetime import timedelta
from fastapi import APIRouter,status
from fastapi.exception_handlers import http_exception_handler

from auth import LoginRequest, Token, authenticate_user, create_access_token,fake_users_db





ACCESS_TOKEN_EXPIRE_MINUTES = 30


auth = APIRouter()


@auth.post("/login", response_model=Token)
async def login_for_access_token(
    form_data:LoginRequest
):
    user = authenticate_user(fake_users_db, form_data.username, form_data.password)
    if not user:
        raise http_exception_handler(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail="Incorrect username or password",
            headers={"WWW-Authenticate": "Bearer"},
        )
    access_token_expires = timedelta(minutes=ACCESS_TOKEN_EXPIRE_MINUTES)
    access_token = create_access_token(
        data={"sub": user.username}, expires_delta=access_token_expires
    )
    return {"access_token": access_token, "token_type": "bearer"}


