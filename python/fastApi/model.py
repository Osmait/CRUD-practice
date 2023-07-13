from pydantic import BaseModel,Field

class PostModel(BaseModel):

    title: str = Field()
    description:str = Field()

