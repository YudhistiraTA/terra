"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { postFormDefaultValue, postFormSchema } from "./postFormSchema";
import { useForm } from "react-hook-form";
import { useState } from "react";
import useApiMutation from "@/hooks/useApiMutation";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "../ui/form";
import { Input } from "../ui/input";
import { Button } from "../ui/button";
import { z } from "zod";
import { Textarea } from "../ui/textarea";

export default function PostForm() {
  const [hasError, setHasError] = useState(false);
  const form = useForm<z.infer<typeof postFormSchema>>({
    resolver: zodResolver(postFormSchema),
    defaultValues: postFormDefaultValue,
  });
  const postMutation = useApiMutation({
    key: "create-post",
  });
  function onSubmit(data: z.infer<typeof postFormSchema>) {
    postMutation.mutateAsync(data).catch(() => setHasError(true));
  }
  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)}>
        <div className="grid w-full items-center gap-4">
          <div className="flex flex-col space-y-1.5">
            {hasError && (
              <div className="text-red-500 text-sm -mt-5">
                Something went wrong. Please try again.
              </div>
            )}
            <FormField
              control={form.control}
              name="title"
              render={({ field }) => (
                <FormItem>
                  <FormLabel htmlFor="title">Title</FormLabel>
                  <FormControl>
                    <Input id="title" placeholder="Title" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>
          <div className="flex flex-col space-y-1.5">
            <FormField
              control={form.control}
              name="content"
              render={({ field }) => (
                <FormItem>
                  <FormLabel htmlFor="content">Content</FormLabel>
                  <FormControl>
                    <Textarea id="content" placeholder="Content" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>
        </div>
        <Button className="mt-2">Submit</Button>
      </form>
    </Form>
  );
}
