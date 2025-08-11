import { API_ENDPOINTS } from "@/lib/api/endpoints";
import { Like } from "@/lib/api/types";
import { clientSideApiClient } from "../api/clientSideApiClient";

export class PostService {
  static async likePost(postID: string): Promise<Like> {
    return clientSideApiClient.post<Like>(API_ENDPOINTS.LIKE(postID));
  }
}
