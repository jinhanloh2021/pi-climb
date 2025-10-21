import { ONBOARDING_STEPS } from "@/app/(onboarding)/onboarding/page";
import { cn } from "@/lib/utils";
import React, { Dispatch, SetStateAction } from "react";
import { useForm } from "react-hook-form";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "./ui/card";
import { Popover, PopoverTrigger, PopoverContent } from "./ui/popover";
import { Button } from "./ui/button";
import { Textarea } from "./ui/textarea";
import { Label } from "./ui/label";
import { Checkbox } from "./ui/checkbox";
import { Calendar } from "./ui/calendar";
import { ArrowLeftIcon, ChevronDownIcon } from "lucide-react";
import { UserService } from "@/lib/service/client/userService";
import { UpdateUserRequest } from "@/lib/api/types";
import { useRouter } from "next/navigation";

type Inputs = {
  bio: string;
  is_public: boolean;
  date_of_birth: string;
};

type Props = {
  setStep: Dispatch<SetStateAction<ONBOARDING_STEPS>>;
};

export function OnboardingFormUserDetails({
  className,
  setStep,
  ...props
}: Props & React.ComponentPropsWithoutRef<"div">) {
  const [open, setOpen] = React.useState(false);

  const {
    register,
    handleSubmit,
    setValue,
    watch,
    formState: { errors, isSubmitting, isValidating },
  } = useForm<Inputs>({
    mode: "onBlur",
    defaultValues: {
      is_public: true,
    },
  });

  const dateOfBirthValue = watch("date_of_birth");
  const router = useRouter();

  const updateProfile = async (data: Inputs) => {
    console.dir(data);
    const reqBody: UpdateUserRequest = {
      ...data,
    };
    try {
      const user = await UserService.updateProfile(reqBody);
      console.dir(user);
    } catch (error) {
      console.error(`Error updating profile: ${error}`);
    } finally {
      router.push("/");
    }
  };

  return (
    <div className={cn("flex flex-col gap-6 my-auto", className)} {...props}>
      <Card>
        <CardHeader>
          <ArrowLeftIcon
            onClick={() => setStep(ONBOARDING_STEPS.USERNAME)}
            className="text-primary hover:text-primary/80 cursor-pointer"
          />
          <CardTitle>Welcome to Pi Climb</CardTitle>
          <CardDescription>Set up your account</CardDescription>
        </CardHeader>
        <CardContent>
          <form onSubmit={handleSubmit(updateProfile)}>
            <div className="flex flex-col gap-6">
              <div className="grid gap-2">
                <Label htmlFor="bio">Your bio</Label>
                <Textarea
                  {...register("bio", {
                    maxLength: {
                      value: 256,
                      message: "Must be 256 characters or less",
                    },
                  })}
                  placeholder="Type your bio here"
                  autoComplete="off"
                />
                {errors.bio && (
                  <p className="text-sm text-red-500">{errors.bio.message}</p>
                )}
              </div>
              <div className="flex items-center gap-3">
                <Checkbox
                  id="is_public"
                  {...register("is_public")}
                  defaultChecked
                />
                <Label htmlFor="is_public">Public profile</Label>
              </div>
              <div className="grid gap-2">
                <Label htmlFor="date" className="px-1">
                  Date of birth
                </Label>
                <Popover open={open} onOpenChange={setOpen}>
                  <PopoverTrigger asChild>
                    <Button
                      variant="outline"
                      id="date"
                      className="w-48 justify-between font-normal"
                    >
                      {dateOfBirthValue
                        ? new Date(dateOfBirthValue).toLocaleDateString()
                        : "Select date"}
                      <ChevronDownIcon />
                    </Button>
                  </PopoverTrigger>
                  <PopoverContent
                    className="w-auto overflow-hidden p-0"
                    align="start"
                  >
                    <Calendar
                      mode="single"
                      selected={
                        dateOfBirthValue
                          ? new Date(dateOfBirthValue)
                          : undefined
                      }
                      captionLayout="dropdown"
                      onSelect={(date) => {
                        if (date) {
                          setValue("date_of_birth", date.toISOString());
                        }
                        setOpen(false);
                      }}
                    />
                  </PopoverContent>
                </Popover>
              </div>
              <Button
                type="submit"
                className="w-full"
                disabled={isSubmitting || isValidating}
              >
                Submit
              </Button>
            </div>
          </form>
        </CardContent>
      </Card>
    </div>
  );
}
