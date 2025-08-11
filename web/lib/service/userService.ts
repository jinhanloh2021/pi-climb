import { API_ENDPOINTS } from "@/lib/api/endpoints";
import { User } from "@/lib/api/types";
import { serverSideApiClient } from "../api/serverSideApiClient";

export class UserService {
  static async getMe(): Promise<User> {
    return serverSideApiClient.get<User>(API_ENDPOINTS.USER_INFO);
  }

  static async getProfile(username: string): Promise<User> {
    return serverSideApiClient.get<User>(API_ENDPOINTS.USER_PROFILE(username));
  }
}
