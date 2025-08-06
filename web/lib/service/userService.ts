// lib/services/userService.ts
import { apiClient } from "@/lib/api/client";
import { API_ENDPOINTS } from "@/lib/api/endpoints";
import { User } from "@/lib/api/types";

export class UserService {
  static async getMe(): Promise<User> {
    return apiClient.get<User>(API_ENDPOINTS.USER_INFO);
  }

  static async getProfile(username: string): Promise<User> {
    return apiClient.get<User>(API_ENDPOINTS.USER_PROFILE(username));
  }
}
