import { PrismaClient, User } from "@prisma/client";

export default class UsersService {
  constructor(private prisma: PrismaClient) {}

  async find(id: number): Promise<User | null> {
    return this.prisma.user.findUnique({
      where: {
        id,
      },
    });
  }
}
