import { createClient } from "@/lib/supabase/server";
import { redirect } from "next/navigation";

export default async function Home() {
  const supabase = await createClient();

  const {
    data: { user },
  } = await supabase.auth.getUser();
  if (user) {
    redirect("/feed");
  }

  return (
    <main className="min-h-screen flex flex-col items-center">
      <div>Not logged in. Marketing page</div>
    </main>
  );
}
