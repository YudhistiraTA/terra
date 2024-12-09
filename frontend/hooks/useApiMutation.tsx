"use client";
import {
  type MutationReturns,
  type MutationKeys,
  mutationConfig,
} from "@/services/mutations";
import { QueryKeys } from "@/services/queries";
import {
  useMutation,
  UseMutationOptions,
  useQueryClient,
} from "@tanstack/react-query";
import cookies from "js-cookie";
import { useRouter } from "next/navigation";

export default function useApiMutation<K extends keyof MutationReturns>({
  key,
  fetchOptions,
  mutationOptions,
  withoutAuth = false,
}: {
  key: MutationKeys;
  fetchOptions?: RequestInit;
  mutationOptions?: UseMutationOptions<MutationReturns[K], Error, unknown>;
  withoutAuth?: boolean;
}) {
  const router = useRouter();
  const { url, method, refetchQueries } = mutationConfig[key];
  const queryClient = useQueryClient();
  const baseUrl = process.env.NEXT_PUBLIC_API_URL;
  const endpoint = new URL(url, baseUrl);

  const invalidQueries = (refetchQueries: QueryKeys[]) => {
    if (!refetchQueries || refetchQueries.length === 0) {
      return;
    }
    const regexPattern = refetchQueries
      .map((query) => `\\b${query}\\b`)
      .join("|");
    const regex = new RegExp(regexPattern, "i");
    queryClient.refetchQueries({
      predicate: (query) =>
        query.queryKey.some(
          (key) => typeof key === "string" && regex.test(key)
        ),
    });
  };

  const headers = new Headers();
  if (!withoutAuth) {
    headers.append("Authorization", "Bearer " + cookies.get("sessionToken"));
  }

  const mutation = useMutation({
    mutationKey: [key],
    mutationFn: async (
      body?: unknown
    ): Promise<{ message: string; data: MutationReturns[K] }> => {
      const res = await fetch(endpoint, {
        headers,
        method,
        credentials: "include",
        body: JSON.stringify(body),
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
    onSuccess() {
      invalidQueries(refetchQueries || []);
    },
    ...mutationOptions,
  });

  return mutation;
}
