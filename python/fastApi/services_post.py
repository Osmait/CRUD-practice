


import json

from domain import Post
from model import PostModel
from file_manager import FileManager

file_manager = FileManager()

class PostService:

    def find_all(self)-> list[Post]:
       return file_manager.read_file()
    
        
    def create(self,post:PostModel)->None:
        post_new = Post(description=post.description,title=post.title)
        file_manager.write_file(post_new) 

