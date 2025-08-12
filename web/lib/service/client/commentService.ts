import { API_ENDPOINTS } from "@/lib/api/endpoints";
import { Comment, Like } from "@/lib/api/types";
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
}
