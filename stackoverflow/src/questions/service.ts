import { PrismaClient } from "@prisma/client";

type CreateQuestionInput = {
  title: string;
  userId: number;
  description: string;
  tags: string[];
};

export default class QuestionsService {
  constructor(private prisma: PrismaClient) {}

  async create({ title, userId, description, tags }: CreateQuestionInput) {
    const uniqueTags = [...new Set(tags)];

    await this.prisma.question.create({
      data: {
        title,
        description,
        tags: {
          connectOrCreate: uniqueTags.map((name) => ({
            create: { name },
            where: { name },
          })),
        },
        userId,
      },
    });
  }
}
