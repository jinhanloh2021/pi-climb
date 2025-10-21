import { requireAuth } from "@/lib/auth/utils";
import { redirect } from "next/navigation";

export default async function OnboardingLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const { publicUser } = await requireAuth();
  if (publicUser.username) {
    redirect("/feed");
  }
  return (
    <main className="min-h-screen flex flex-col items-center">{children}</main>
  );
}
