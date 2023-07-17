
from  routes import post_routes

from datetime import datetime, timedelta
from typing import Annotated

from fastapi import Depends, FastAPI, HTTPException, status
from fastapi.security import OAuth2PasswordBearer
from jose import JWTError, jwt
from passlib.context import CryptContext
from pydantic import BaseModel
from auth_routes import auth


app = FastAPI()
app.include_router(post_routes)
app.include_router(auth)





# @app.get("/users/me/", response_model=User)
# async def read_users_me(
#     current_user: Annotated[User, Depends(get_current_active_user)]
# ):
#     return current_user


# @app.get("/users/me/items/")
# async def read_own_items(
#     current_user: Annotated[User, Depends(get_current_active_user)]
# ):
#     return [{"item_id": "Foo", "owner": current_user.username}]

# @app.get("/")
# def hello(current_user:Annotated[User,Depends(get_current_active_user)]):
#     return "Hello word"