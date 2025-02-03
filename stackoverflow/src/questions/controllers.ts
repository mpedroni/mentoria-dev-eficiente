import { z } from "zod";
import QuestionsService from "./service";
import UsersService from "../users/service";
import BadRequestError from "../errors/bad-request.error";

export const CreateQuestionSchema = z.object({
  title: z.string().max(150),
  description: z.string(),
  tags: z.array(z.string().min(1)).max(5).default([]),
});

export type CreateQuestionRequest = z.infer<typeof CreateQuestionSchema>;

export default class QuestionsController {
  constructor(
    private questions: QuestionsService,
    private users: UsersService
  ) {}

  async create(request: CreateQuestionRequest, userId: number): Promise<void> {
    const user = await this.users.find(userId);
    if (!user) {
      throw new BadRequestError("User not found");
    }

    return this.questions.create({
      description: request.description,
      tags: request.tags,
      title: request.title,
      userId,
    });
  }
}
