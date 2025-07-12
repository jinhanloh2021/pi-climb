export interface ApiResponse<T> {
  data: T;
  message?: string;
}

export interface PaginatedResponse<T> {
  data: T[];
  next_cursor?: string;
  has_more: boolean;
}

export interface User {
  id: string;
  email: string;
  username: string;
  bio: string;
  is_public: boolean;
  phone_number: string;
  date_of_birth: string;

  followers: any;
  follower_count: number;
  following: any;
  following_count: number;

  avatar: any;
  posts: any;
  likes: any;
  comments: any;
  media: any;

  created_at: string;
  updated_at: string;
}

export interface Post {
  id: string;
  user_id: string;
  caption?: string;
  hold_colour?: string;
  grade?: string;
  created_at: string;
  user: User;
  media: Media[];
  // Add other post fields
}

export interface Media {
  id: string;
  url: string;
  media_type: string;
  width?: number;
  height?: number;
  // Add other media fields
}
