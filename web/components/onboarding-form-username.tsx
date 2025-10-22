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
import { Dispatch, SetStateAction } from "react";
import { useForm } from "react-hook-form";
import { User } from "@supabase/supabase-js";
import { UserService } from "@/lib/service/client/userService";
import { ApiError } from "@/lib/api/clientSideApiClient";
import { UpdateUserRequest } from "@/lib/api/types";
import { ONBOARDING_STEPS } from "@/lib/api/types";
import { Label } from "./ui/label";

type Inputs = {
  username: string;
};

type Props = {
  setStep: Dispatch<SetStateAction<ONBOARDING_STEPS>>;
  authUser: User | null;
};

export function OnboardingFormUsername({
  className,
  setStep,
  authUser,
  ...props
}: Props & React.ComponentPropsWithoutRef<"div">) {
  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting, isValidating },
  } = useForm<Inputs>({
    mode: "onBlur",
  });

  const onSubmit = async (data: Inputs) => {
    const reqBody: UpdateUserRequest = {
      username: data.username,
    };
    try {
      const user = await UserService.updateProfile(reqBody);
      setStep(ONBOARDING_STEPS.USER_DETAILS);
      console.dir(user);
    } catch (error) {
      console.error(`Error updating username: ${error}`);
    }
  };

  // TODO: Create endpoint just for username validation
  const validateUsername = async (username: string) => {
    try {
      const user = await UserService.getProfile(username);
      if (user) {
        return "Username is already taken";
      }
      return "Error finding username";
    } catch (error) {
      if ((error as ApiError).status == 404) {
        // Valid username
        return;
      } else {
        return "Error finding username";
      }
    }
  };

  return (
    <div className={cn("flex flex-col gap-6 my-auto", className)} {...props}>
      <Card>
        <CardHeader>
                    <CardTitle>Welcome to Pi Climb</CardTitle>
          <CardDescription>Set up your account</CardDescription>
        </CardHeader>
        <CardContent>
          <form onSubmit={handleSubmit(onSubmit)}>
            <div className="flex flex-col gap-6">
              <div className="grid gap-2">
                <Label htmlFor="email">Email</Label>
                <Input
                  placeholder="email"
                  autoComplete="off"
                  value={authUser?.email ?? ""}
                  disabled
                />
                <Label htmlFor="username">Username</Label>
                <Input
                  {...register("username", {
                    required: { value: true, message: "Username is required" },
                    minLength: {
                      value: 2,
                      message: "Must be at least 2 characters",
                    },
                    maxLength: {
                      value: 64,
                      message: "Must be 64 characters or less",
                    },
                    pattern: {
                      value: /^[a-zA-Z0-9_]+$/,
                      message: "Only a-z, A-Z, 0-9 and _ allowed",
                    },
                    validate: validateUsername,
                  })}
                  placeholder="username"
                  autoComplete="off"
                />
                {errors.username && (
                  <p className="text-sm text-red-500">
                    {errors.username.message}
                  </p>
                )}
              </div>
              <Button
                type="submit"
                className="w-full"
                disabled={isSubmitting || isValidating}
              >
                Next
              </Button>
            </div>
          </form>
        </CardContent>
      </Card>
    </div>
  );
}
