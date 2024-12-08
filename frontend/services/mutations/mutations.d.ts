import { QueryKeys } from "../queries";

export type MutationConfig<T extends string> = {
  [key in T]: {
    url: string;
    method: "POST" | "PUT" | "DELETE" | "PATCH";
    refetchQueries?: QueryKeys[];
  };
};
