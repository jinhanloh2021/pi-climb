import { createClient } from "@/lib/supabase/server";
import { Session, User as AuthUser } from "@supabase/supabase-js";
import { redirect } from "next/navigation";
import { UserService } from "@/lib/service/server/userService";
import { User as PublicUser } from "@/lib/api/types";

export async function requireAuth(): Promise<{
  authUser: AuthUser; // auth.user from supabase
  publicUser: PublicUser; // public.user
  session: Session;
  token: string;
}> {
  const supabase = await createClient();
  const {
    data: { user: authUser },
    error,
  } = await supabase.auth.getUser();

  if (error || !authUser) {
    redirect("/auth/login");
  }

  const {
    data: { session },
    error: sessionError,
  } = await supabase.auth.getSession();

  if (sessionError || !session?.access_token) {
    redirect("/auth/login");
  }

  // Get logged in User profile
  let publicUser = null;
  try {
    publicUser = await UserService.getMe();
  } catch (e) {
    console.error("Could not fetch user profile for authenticated user:", e);
    redirect("/auth/login");
  }

  return {
    authUser: authUser,
    publicUser: publicUser,
    session: session,
    token: session.access_token,
  };
}
