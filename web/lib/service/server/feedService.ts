import { API_ENDPOINTS } from "@/lib/api/endpoints";
import { FeedResponse } from "@/lib/api/types";
import { serverSideApiClient } from "@/lib/api/serverSideApiClient";

export class FeedService {
  static async getFeed(
    followCursor?: string,
    trendCursor?: string,
    limit?: number,
  ): Promise<FeedResponse> {
    return serverSideApiClient.get<FeedResponse>(
      API_ENDPOINTS.FEED(followCursor, trendCursor, limit),
    );
  }
}
