
import os
from typing import Annotated
from fastapi import APIRouter,UploadFile,Form,File,Depends
import logging.config
from model import PostModel
from services_post import PostService
from auth import get_current_active_user

post_routes = APIRouter(prefix="/post")
post_service = PostService()
# setup loggers
logging.config.fileConfig('logging.conf', disable_existing_loggers=False)

# get root logger
logger = logging.getLogger(__name__)

@post_routes.get("/")
def get_post(_: str = Depends(get_current_active_user)):
    logging.info("call service Find all")
    post =  post_service.find_all()
    logging.info("returning post...")
    return post

@post_routes.post("/")
def create_post(post:PostModel,_: str = Depends(get_current_active_user)):
    logging.info("Creating post...")
    post_service.create(post)
    logging.info("Created!") 
    return "Created"


@post_routes.post("/upload")
def upload_file(file: Annotated[UploadFile, File()],_: str = Depends(get_current_active_user)):
    conten = file.file.read()
    try:
        if not os.path.exists("upload"):
            os.mkdir("upload")
        with open(f"upload/{file.filename}","wb")as f:
            f.write(conten)
    except Exception as e:
        return {"message":"Erro Uploading file"}
    finally:
        file.file.close()
    return "Uploaded!"