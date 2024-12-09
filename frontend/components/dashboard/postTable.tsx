"use client";
import useApiQuery from "@/hooks/useApiQuery";
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "../ui/accordion";
import { LoaderCircle } from "lucide-react";
import { Button } from "../ui/button";
import { useQueryState } from "nuqs";
import { Input } from "../ui/input";

import { useState, useEffect } from "react";

export default function PostTable() {
  const [cursor, setCursor] = useQueryState("cursor", { history: "push" });
  const [search, setSearch] = useQueryState("search", {
    history: "push",
    throttleMs: 800,
  });
  const [debouncedSearch, setDebouncedSearch] = useState(search);

  useEffect(() => {
    const handler = setTimeout(() => {
      setDebouncedSearch(search);
      setCursor(null);
    }, 800);

    return () => {
      clearTimeout(handler);
    };
  }, [search, setCursor]);

  const { data, isLoading } = useApiQuery({
    key: "posts",
    fetchOptions: { params: { cursor, search: debouncedSearch } },
  });

  if (isLoading) {
    return <LoaderCircle className="animate-spin self-center w-full" />;
  }
  if (data?.data?.posts?.length === 0) {
    return (
      <>
        <Input
          value={search ?? ""}
          onChange={(e) => setSearch(e.target.value)}
          placeholder="Search"
          className="mb-2"
        />
        <p className="self-center w-full">No posts found</p>
      </>
    );
  }
  return (
    <>
      <Input
        value={search ?? ""}
        onChange={(e) => setSearch(e.target.value)}
        placeholder="Search"
        className="mb-2"
      />
      <Accordion type="single" collapsible>
        {data?.data?.posts?.map((post) => (
          <AccordionItem className="border-b-0" key={post.id} value={post.id}>
            <AccordionTrigger className="bg-stone-700 rounded my-1 px-2">
              {post.title}
            </AccordionTrigger>
            <AccordionContent className="bg-stone-500 p-4 rounded ml-3">
              {post.content}
            </AccordionContent>
          </AccordionItem>
        ))}
      </Accordion>
      <div className="flex w-full justify-between mt-2">
        <Button
          disabled={!data?.data?.previous_cursor}
          onClick={() => setCursor(data?.data?.previous_cursor ?? null)}
        >
          Previous Page
        </Button>
        <Button
          disabled={!data?.data?.next_cursor}
          onClick={() => setCursor(data?.data?.next_cursor ?? null)}
        >
          Next Page
        </Button>
      </div>
    </>
  );
}
