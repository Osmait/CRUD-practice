
import uuid


class Post:
    def __init__(self,title,description,id=None ) -> None:
       self.id = id or uuid.uuid4()
       self.title = title
       self.description = description 


