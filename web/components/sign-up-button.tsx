import { Button } from "@/components/ui/button";
import Link from "next/link";

export function SignUpButton() {
  return (
    <Button asChild>
      <Link href={"/auth/sign-up"}>Sign Up</Link>
    </Button>
  );
}
