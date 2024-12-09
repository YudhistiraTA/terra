"use client";
import {
  type QueryKeys,
  type QueryReturns,
  queryUrl,
} from "@/services/queries";
import { useQuery, UseQueryOptions } from "@tanstack/react-query";
import cookies from "js-cookie";
import { useRouter } from "next/navigation";

export default function useApiQuery<K extends keyof QueryReturns>({
  key,
  fetchOptions,
  queryOptions,
  withoutAuth = false,
}: {
  key: QueryKeys;
  fetchOptions?: RequestInit;
  queryOptions?: Omit<
    UseQueryOptions<QueryReturns[K], Error>,
    "queryKey" | "queryFn" | "staleTime"
  >;
  withoutAuth?: boolean;
}) {
  const router = useRouter();
  const url = queryUrl[key];
  const baseUrl = process.env.NEXT_PUBLIC_API_URL;
  const endpoint = new URL(url, baseUrl);
  const sessionToken = cookies.get("sessionToken");

  const headers = new Headers();
  if (!withoutAuth) {
    headers.append("Authorization", "Bearer " + sessionToken);
  }

  const result = useQuery<QueryReturns[K]>({
    queryKey: [key],
    queryFn: async () => {
      const res = await fetch(endpoint, {
        headers,
        credentials: "include",
        ...fetchOptions,
      });
      if (!res.ok) {
        if (res.status === 401) {
          cookies.remove("sessionToken");
          cookies.remove("refreshToken");
          router.push("/login");
          throw new Error("Unauthorized");
        }
        const body = await res.json();
        throw new Error(body?.message || "Unknown error");
      }
      return await res.json();
    },
    retry: false,
    staleTime: 1000 * 60 * 5,
    ...queryOptions,
  });

  return result;
}
