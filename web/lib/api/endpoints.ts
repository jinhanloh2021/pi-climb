export const API_ENDPOINTS = {
  // User endpoints
  USER_INFO: "/api/v0/myinfo",
  USER_PROFILE: (username: string) => `/api/v0/users/username/${username}`,

  // Post endpoints
  POSTS: "/api/v0/posts",
  POST_DETAIL: (id: string) => `/api/v0/posts/${id}`,

  // Feed endpoints
  FEED: (followCursor?: string, trendCursor?: string, limit?: number) =>
    `/api/v0/feed?following-cursor=${followCursor}&trending-cursor=${trendCursor}&limit=${limit}`,
} as const;
