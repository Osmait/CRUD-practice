import Express from "express";
import { Request, Response } from "express";

const app = Express();

app.get("/", (req: Request, res: Response) => {
  res.send("hello word ");
});

app.listen(8000, () => {
  console.log("server up");
});
