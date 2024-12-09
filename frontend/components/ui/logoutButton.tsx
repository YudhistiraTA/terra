"use client";
import { useRouter } from "next/navigation";
import { Button } from "./button";
import cookies from "js-cookie";

export default function LogoutButton() {
  const router = useRouter();
  return (
    <Button
      variant="secondary"
      className="font-bold"
      onClick={() => {
        cookies.remove("sessionToken");
        cookies.remove("refreshToken");
        router.push("/login");
      }}
    >
      Logout
    </Button>
  );
}
