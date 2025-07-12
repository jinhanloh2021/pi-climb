export const API_ENDPOINTS = {
  // User endpoints
  USER_INFO: "/api/v0/myinfo",
  USER_PROFILE: (id: string) => `/api/v0/users/${id}`,

  // Post endpoints
  POSTS: "/api/v0/posts",
  POST_DETAIL: (id: string) => `/api/v0/posts/${id}`,

  // Feed endpoints
  FEED: "/api/v0/feed",
} as const;
