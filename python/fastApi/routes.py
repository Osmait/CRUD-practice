import json
from fastapi import APIRouter

from model import PostModel
from services_post import PostService

post_routes = APIRouter(prefix="/post")
post_service = PostService()


@post_routes.get("/")
def get_post():
    post =  post_service.find_all()


    return post


@post_routes.post("/")
def create_post(post:PostModel):
        post_service.create(post) 
        return "Created"


        