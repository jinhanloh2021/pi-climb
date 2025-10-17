"use client";

import { Button } from "@/components/ui/button";
import Link from "next/link";

export function LoginButton() {
  return (
    <Button variant={"outline"}>
      <Link href={"/auth/login"}>Login</Link>
    </Button>
  );
}
