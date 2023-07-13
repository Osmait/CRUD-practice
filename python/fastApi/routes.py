
from fastapi import APIRouter
import logging.config

from model import PostModel
from services_post import PostService

post_routes = APIRouter(prefix="/post")
post_service = PostService()
# setup loggers
logging.config.fileConfig('logging.conf', disable_existing_loggers=False)

# get root logger
logger = logging.getLogger(__name__)

@post_routes.get("/")
def get_post():
    logging.info("call service Find all")
    post =  post_service.find_all()
    logging.info("returning post...")
    return post

@post_routes.post("/")
def create_post(post:PostModel):
    logging.info("Creating post...")
    post_service.create(post)
    logging.info("Created!") 
    return "Created"


        