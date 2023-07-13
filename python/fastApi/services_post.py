


import json

from domain import Post
from model import PostModel


class PostService:


    def find_all(self):
        post:list[Post] = []
        with open("db.json","r") as f:
            post_Json = json.loads(f.read()) 
            print(post_Json)
        
            for post_object in post_Json:
                instancia = Post(post_object["title"],post_object["description"],post_object["id"])
                post.append(instancia)

        return post

    def create(self,post:PostModel):
        post_list = self.find_all()
        post_new = Post(description=post.description,title=post.title)
        post_list.append(post_new)
        post_Json = []
        for post_object in post_list:
            post_Json.append({
                "title": post_object.title,
                "description": post_object.description,
                "id": str(post_object.id)
                })
        with open("db.json","w") as f:
            json.dump(post_Json,f,indent=4)

