// lib/services/userService.ts
import { apiClient } from "@/lib/api/client";
import { API_ENDPOINTS } from "@/lib/api/endpoints";
import { User } from "@/lib/api/types";

export class UserService {
  static async getMyInfo(): Promise<User> {
    return apiClient.get<User>(API_ENDPOINTS.USER_INFO);
  }

  static async getProfile(userId: string): Promise<User> {
    return apiClient.get<User>(API_ENDPOINTS.USER_PROFILE(userId));
  }

  static async updateProfile(data: Partial<User>): Promise<User> {
    return apiClient.put<User>(API_ENDPOINTS.USER_INFO, data);
  }
}
