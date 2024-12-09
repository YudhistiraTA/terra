"use client";
import { useRouter } from "next/navigation";
import { Button } from "./button";
import cookies from "js-cookie";
import { useQueryClient } from "@tanstack/react-query";

export default function LogoutButton() {
  const query = useQueryClient();
  const router = useRouter();
  return (
    <Button
      variant="secondary"
      className="font-bold"
      onClick={() => {
        cookies.remove("sessionToken");
        cookies.remove("refreshToken");
        query.clear();
        router.push("/login");
      }}
    >
      Logout
    </Button>
  );
}
