

import json

from domain import Post


class FileManager:
    file_path ="db.json"

    def read_file(self)-> list[Post]:
        post:list[Post] = []
        with open(self.file_path,"r") as f:
            post_Json = json.loads(f.read()) 
          
        
            for post_object in post_Json:
                instancia = Post(post_object["title"],post_object["description"],post_object["id"])
                post.append(instancia)
        return post
    

    def write_file(self,post_new)-> None:
        post_list = self.read_file()
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

