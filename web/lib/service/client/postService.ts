import { API_ENDPOINTS } from "@/lib/api/endpoints";
import { Like } from "@/lib/api/types";
import { clientSideApiClient } from "@/lib/api/clientSideApiClient";

export class PostService {
  static async likePost(postID: string): Promise<Like> {
    return clientSideApiClient.post<Like>(API_ENDPOINTS.LIKE(postID));
  }

  static async unlikePost(postID: string): Promise<void> {
    return clientSideApiClient.delete<void>(API_ENDPOINTS.LIKE(postID));
  }
}
