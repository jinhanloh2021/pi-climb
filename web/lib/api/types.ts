export interface ApiResponse<T> {
  data: T;
  message?: string;
}

export interface PaginatedResponse<T> {
  data: T[];
  next_cursor?: string;
  has_more: boolean;
}

export interface FeedResponse {
  next_cursor: {
    following_cursor: string;
    trending_cursor: string;
  };
  posts: Post[];
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
  deleted_at?: string;
}

export interface Like {
  user_id: number;
  user?: User;

  post_id: number;
  post?: Post;

  created_at: string;
  updated_at: string;
  deleted_at?: string;
}

export interface Gym {
  id: number;
  name: string;
  grading_system: string[];
  google_place_id?: string;
  google_maps_uri?: string;
  address?: string;
  latitude?: number;
  longitude?: number;

  posts: Post[];
  media: Media[];

  created_at: string;
  updated_at: string;
  deleted_at?: string;
}

export interface Post {
  id: number;
  caption?: string;
  hold_colour?: string;
  grade?: string;

  user_id: string;
  user?: User;

  media: Media[];

  likes: Like[];
  like_count: number;

  views: number;

  comments: Comment[];
  comment_count: number;

  gym_id?: number;
  gym?: Gym;

  created_at: string;
  updated_at: string;
  deleted_at?: string;
}

export interface Media {
  id: number;
  storage_key: string;
  bucket: string;

  original_name: string;
  file_size: number;

  mime_type: string;
  order?: number;

  width?: number;
  height?: number;
  duration?: number;

  owner_id: number;
  owner_type: string;

  user_id: number;
  user?: User;
}
