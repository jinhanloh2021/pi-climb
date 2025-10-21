import { clientSideApiClient } from "@/lib/api/clientSideApiClient";
import { API_ENDPOINTS } from "@/lib/api/endpoints";
import { UpdateUserRequest, User } from "@/lib/api/types";

export class UserService {
  static async getMe(): Promise<User> {
    return clientSideApiClient.get<User>(API_ENDPOINTS.USER_INFO);
  }

  static async getProfile(username: string): Promise<User> {
    return clientSideApiClient.get<User>(API_ENDPOINTS.USER_PROFILE(username));
  }

  static async updateProfile(updates: UpdateUserRequest): Promise<User> {
    return clientSideApiClient.patch<User>(
      API_ENDPOINTS.UPDATE_PROFILE,
      updates,
    );
  }
}
