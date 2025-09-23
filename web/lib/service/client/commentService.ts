import { API_ENDPOINTS } from "@/lib/api/endpoints";
import { Comment } from "@/lib/api/types";
import { clientSideApiClient } from "@/lib/api/clientSideApiClient";

export class CommentService {
  static async createComment(
    body: { text: string },
    postID: string,
  ): Promise<Comment> {
    return clientSideApiClient.post<Comment>(
      API_ENDPOINTS.POST_COMMMENT(postID),
      body,
    );
  }

  static async getComments(postID: string): Promise<Comment[]> {
    return clientSideApiClient.get<Comment[]>(
      API_ENDPOINTS.POST_COMMMENT(postID),
    );
  }
}
