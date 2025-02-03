import { PrismaClient } from "@prisma/client";
import express, { Request, Response } from "express";
import { z } from "zod";

import QuestionsController, {
  CreateQuestionSchema,
} from "./questions/controllers";
import QuestionsService from "./questions/service";
import UsersService from "./users/service";
import BadRequestError from "./errors/bad-request.error";

const app = express();

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

const prisma = new PrismaClient();
const questions = new QuestionsController(
  new QuestionsService(prisma),
  new UsersService(prisma)
);

type AuthorizedRequest = express.Request & {
  user: { id: number; name: string; isMember: boolean };
};

app.use((req, res, next) => {
  (req as AuthorizedRequest).user = {
    id: 2,
    name: "John Doe",
    isMember: true,
  };

  next();
});

app.post("/questions", async (req, res, next) => {
  try {
    CreateQuestionSchema.parse(req.body);
    const { user } = req as AuthorizedRequest;

    await questions.create(req.body, user.id);

    return res.status(201).send();
  } catch (error) {
    next(error);
  }
});

app.use((err: Error, req: Request, res: Response, next: any) => {
  if (err instanceof z.ZodError) {
    return res.status(400).json({
      message: "Bad Request",
      detail: "Invalid request data",
      errors: err.errors,
    });
  }

  if (err instanceof BadRequestError) {
    return res.status(400).json({
      message: "Bad Request",
      detail: err.message,
    });
  }

  next(err);
});

app.listen(3000, () => {
  console.log("Server is running on port 3000");
});
