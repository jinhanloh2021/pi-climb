import { AuthButton } from "@/components/auth-button";
import { ThemeSwitcher } from "@/components/theme-switcher";
import { PublicUserProvider } from "@/context/PublicUserProvider";
import { requireAuth } from "@/lib/auth/utils";
import { redirect } from "next/navigation";

export default async function ProtectedLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const { publicUser } = await requireAuth();
  if (!publicUser.username) {
    redirect("/onboarding");
  }

  return (
    <PublicUserProvider publicUser={publicUser}>
      <main className="min-h-screen flex flex-col items-center">
        <div className="flex-1 w-full flex flex-col gap-20 items-center">
          <nav className="w-full flex justify-center border-b border-b-foreground/10 h-16">
            <div className="w-full max-w-5xl flex justify-end items-center p-3 px-5 text-sm gap-5">
              <div className="flex gap-5 items-center font-semibold"></div>
              <AuthButton />
              <ThemeSwitcher />
            </div>
          </nav>
          <>{children}</>
          <footer className="w-full flex items-center justify-center border-t mx-auto text-center text-xs gap-8 py-16"></footer>
        </div>
      </main>
    </PublicUserProvider>
  );
}
