import { createClient } from "@/lib/supabase/server";
import { Session, User } from "@supabase/supabase-js";
import { redirect } from "next/navigation";

export async function requireAuth(): Promise<{
  user: User;
  session: Session;
  token: string;
}> {
  const supabase = await createClient();
  const {
    data: { user },
    error,
  } = await supabase.auth.getUser();

  if (error || !user) {
    redirect("/auth/login");
  }

  const {
    data: { session },
    error: sessionError,
  } = await supabase.auth.getSession();

  if (sessionError || !session?.access_token) {
    redirect("/auth/login");
  }

  return {
    user,
    session,
    token: session.access_token,
  };
}

export async function getAuthToken(): Promise<string> {
  const supabase = await createClient();
  const {
    data: { session },
    error,
  } = await supabase.auth.getSession();

  if (error || !session?.access_token) {
    throw new Error("Authentication required");
  }

  return session.access_token;
}
