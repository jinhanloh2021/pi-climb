import { API_ENDPOINTS } from "@/lib/api/endpoints";
import { MyLikedPostResponse } from "@/lib/api/types";
import { serverSideApiClient } from "@/lib/api/serverSideApiClient";

export class PostService {
  static async myLikedPost(postID: string): Promise<MyLikedPostResponse> {
    return serverSideApiClient.get<MyLikedPostResponse>(
      API_ENDPOINTS.MY_LIKED_POST(postID),
    );
  }
}
