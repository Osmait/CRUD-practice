from fastapi import FastAPI
from  routes import post_routes




app = FastAPI()
app.include_router(post_routes)

@app.get("/")
def hello():
    return "Hello word"