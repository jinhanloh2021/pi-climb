import { useState, useEffect } from "react";

// NOTE: ideally use env variable NEXT_PUBLIC_APP_ENV for each env
//
/**
 * Hook to check if the app is running in dev env
 * Checks hostname is 'localhost' or starts with 'dev.'.
 * Only client side
 */
export const useIsDev = () => {
  const [isDev, setIsDev] = useState(false);

  useEffect(() => {
    const hostname = window.location.hostname;
    if (hostname.includes("localhost") || hostname.startsWith("dev.")) {
      setIsDev(true);
    }
  }, []);
  return isDev;
};
