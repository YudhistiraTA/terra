"use client";
import { useForm } from "react-hook-form";
import { Input } from "../ui/input";
import { zodResolver } from "@hookform/resolvers/zod";
import { loginFormDefaultValue, loginFormSchema } from "./loginFormSchema";
import { z } from "zod";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "../ui/form";
import { CardContent, CardFooter } from "../ui/card";
import { Button } from "../ui/button";
import useApiMutation from "@/hooks/useApiMutation";
import { useRouter } from "next/navigation";
import { PasswordInput } from "../ui/password-input";
import { useState } from "react";

export default function LoginForm() {
  const router = useRouter();
  const [invalidLogin, setInvalidLogin] = useState(false);
  const form = useForm<z.infer<typeof loginFormSchema>>({
    resolver: zodResolver(loginFormSchema),
    defaultValues: loginFormDefaultValue,
  });
  const loginMutation = useApiMutation({
    key: "login",
  });
  function onSubmit(data: z.infer<typeof loginFormSchema>) {
    loginMutation
      .mutateAsync(data)
      .then(() => {
        router.push("/");
      })
      .catch(() => setInvalidLogin(true));
  }
  return (
    <>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)}>
          <CardContent>
            <div className="grid w-full items-center gap-4">
              <div className="flex flex-col space-y-1.5">
                {invalidLogin && (
                  <div className="text-red-500 text-sm -mt-5">
                    Invalid email or password.
                  </div>
                )}
                <FormField
                  control={form.control}
                  name="email"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel htmlFor="email">Email</FormLabel>
                      <FormControl>
                        <Input
                          id="email"
                          placeholder="example@mail.com"
                          type="email"
                          {...field}
                        />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </div>
              <div className="flex flex-col space-y-1.5">
                <FormField
                  control={form.control}
                  name="password"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel htmlFor="password">Password</FormLabel>
                      <FormControl>
                        <PasswordInput
                          id="password"
                          placeholder="Enter your password"
                          {...field}
                        />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </div>
            </div>
          </CardContent>
          <CardFooter className="flex justify-end">
            <Button>Login</Button>
          </CardFooter>
        </form>
      </Form>
    </>
  );
}
