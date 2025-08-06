import { apiClient } from "@/lib/api/client";
import { API_ENDPOINTS } from "@/lib/api/endpoints";
import { FeedResponse } from "@/lib/api/types";

export class FeedService {
  static async getFeed(
    followCursor?: string,
    trendCursor?: string,
    limit?: number,
  ): Promise<FeedResponse> {
    return apiClient.get<FeedResponse>(
      API_ENDPOINTS.FEED(followCursor, trendCursor, limit),
    );
  }
}
