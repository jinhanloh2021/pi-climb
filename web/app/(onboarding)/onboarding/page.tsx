"use client";
import { OnboardingFormUserDetails } from "@/components/onboarding-form-user-details";
import { OnboardingFormUsername } from "@/components/onboarding-form-username";
import { createClient } from "@/lib/supabase/client";
import { User } from "@supabase/supabase-js";
import { useEffect, useState } from "react";

export enum ONBOARDING_STEPS {
  USERNAME,
  USER_DETAILS,
}

export default function OnboardingPage() {
  const [step, setStep] = useState(ONBOARDING_STEPS.USERNAME);
  const [authUser, setAuthUser] = useState<User | null>(null);
  const supabase = createClient();
  useEffect(() => {
    async function fetchAndSetAuthUser() {
      const {
        data: { user },
        error,
      } = await supabase.auth.getUser();
      if (error) {
      } else {
        setAuthUser(user);
      }
    }
    fetchAndSetAuthUser();
  }, [supabase]);

  return (
    <>
      {step == ONBOARDING_STEPS.USERNAME && (
        <OnboardingFormUsername setStep={setStep} authUser={authUser} />
      )}
      {step == ONBOARDING_STEPS.USER_DETAILS && (
        <OnboardingFormUserDetails setStep={setStep} />
      )}
    </>
  );
}
