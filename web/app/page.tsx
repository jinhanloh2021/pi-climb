import { LoginButton } from "@/components/login-button";
import { SignUpButton } from "@/components/sign-up-button";
import { ThemeSwitcher } from "@/components/theme-switcher";
import { createClient } from "@/lib/supabase/server";
import Link from "next/link";
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
      <nav className="w-full flex justify-center border-b border-b-foreground/10 h-16">
        <div className="w-full max-w-5xl flex justify-end items-center p-3 px-5 text-sm gap-5">
          <ThemeSwitcher />
          <SignUpButton />
          <LoginButton />
        </div>
      </nav>
      <div className="w-full flex flex-col flex-grow justify-center items-center pb-[10%]">
        <h1 className="text-7xl font-semibold">Pi Climb</h1>
        <p className="w-[80%] md:w-[60%] lg:w-[50%] text-slate-500 text-center mt-4">
          Site is under development. Visit GitHub{" "}
          <Link
            className="text-blue-400"
            target="_blank"
            href={"https://github.com/jinhanloh2021/pi-climb"}
          >
            repository
          </Link>{" "}
          for full project documentation
        </p>
      </div>
    </main>
  );
}
