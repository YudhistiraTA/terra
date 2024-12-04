"use client";
import generateSignature from "@/lib/generateSignature";
import {
  type QueryKeys,
  type QueryReturns,
  queryUrl,
} from "@/services/queries";
import { useQuery, UseQueryOptions } from "@tanstack/react-query";

export default function useApiQuery<K extends keyof QueryReturns>({
  key,
  fetchOptions,
  queryOptions,
}: {
  key: QueryKeys;
  fetchOptions?: RequestInit;
  queryOptions?: Omit<
    UseQueryOptions<QueryReturns[K], Error>,
    "queryKey" | "queryFn" | "staleTime"
  >;
}) {
  const url = queryUrl[key];
  const baseUrl = process.env.NEXT_PUBLIC_API_URL;
  const endpoint = new URL(url, baseUrl);

  const { signature, timestamp } = generateSignature({
    method: "GET",
    endpoint: url,
  });

  const result = useQuery<QueryReturns[K]>({
    queryKey: [key],
    queryFn: async () => {
      const res = await fetch(endpoint, {
        headers: {
          "X-SIGNATURE": signature,
          "X-TIMESTAMP": timestamp,
        },
        ...fetchOptions,
      });
      if (!res.ok) {
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
