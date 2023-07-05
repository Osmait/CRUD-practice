import Express from "express";
import { Request, Response } from "express";

const app = Express();

app.use(Express.json())


interface post {
  id: number
  title: string,
  conntent: string
}

let db: post[] = []


app.get("/", (req: Request, res: Response) => {
  res.status(200).json(db)
});

app.post("/", (req: Request, res: Response) => {
  const body: post = req.body;

  db.push(body);
  res.status(201).send("Created");
})

app.delete("/:id", (req: Request, res: Response) => {
  const { id } = req.params
  db = db.filter((p: post) => p.id != Number(id))
  res.status(200).json(db)

})

app.listen(8000, () => {
  console.log("server up");
});
