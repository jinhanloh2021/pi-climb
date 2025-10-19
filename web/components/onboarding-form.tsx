"use client";

import { cn } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { useRouter } from "next/navigation";
import { useState } from "react";
import { useForm } from "react-hook-form";

type Inputs = {
  username: string;
  bio?: string;
  is_public?: boolean;
  phone_number?: string;
  date_of_birth?: string;
};

export function OnboardingForm({
  className,
  ...props
}: React.ComponentPropsWithoutRef<"div">) {
  const [username, setUsername] = useState("");
  const [error, setError] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const router = useRouter();

  // TODO: multi stage form for username, dob, phone number, profile picture
  const { register, handleSubmit } = useForm<Inputs>();

  return (
    <div className={cn("flex flex-col gap-6 my-auto", className)} {...props}>
      <Card>
        <CardHeader>
          <CardTitle className="text-2xl flex flex-row justify-left items-center gap-2">
            Welcome to Pi Climb
          </CardTitle>
          <CardDescription>Set up your account</CardDescription>
        </CardHeader>
        <CardContent>
          <form onSubmit={() => {}}>
            <div className="flex flex-col gap-6">
              <div className="grid gap-2">
                <Input
                  {...register("username", { required: true })}
                  placeholder="username"
                  autoComplete="off"
                />
              </div>
              {error && <p className="text-sm text-red-500">{error}</p>}
              <Button type="submit" className="w-full" disabled={isLoading}>
                {isLoading ? "Submitting" : "Next"}
              </Button>
            </div>
          </form>
        </CardContent>
      </Card>
    </div>
  );
}
