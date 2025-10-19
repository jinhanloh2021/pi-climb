"use client";

import { createContext, useContext, ReactNode } from "react";
import { User } from "@/lib/api/types";

const PublicUserContext = createContext<User | null>(null);

export const PublicUserProvider = ({
  children,
  publicUser,
}: {
  children: ReactNode;
  publicUser: User;
}) => {
  return (
    <PublicUserContext.Provider value={publicUser}>
      {children}
    </PublicUserContext.Provider>
  );
};

export const usePublicUser = (): User => {
  const context = useContext(PublicUserContext);
  if (!context) {
    throw new Error("usePublicUser must be used within a PublicUserProvider");
  }
  return context;
};
