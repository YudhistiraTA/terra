"use client";
import useApiQuery from "@/hooks/useApiQuery";
import { Input } from "../ui/input";
import { Label } from "../ui/label";
import { PasswordInput } from "../ui/password-input";

export default function LoginForm() {
  useApiQuery({ key: "health" });
  return (
    <form>
      <div className="grid w-full items-center gap-4">
        <div className="flex flex-col space-y-1.5">
          <Label htmlFor="email">Email</Label>
          <Input id="email" placeholder="example@mail.com" type="email" />
        </div>
        <div className="flex flex-col space-y-1.5">
          <Label htmlFor="password">Password</Label>
          <PasswordInput id="password" placeholder="Enter your password" />
        </div>
      </div>
    </form>
  );
}
