import { MutationConfig } from "./mutations";

export type PostMutationKeys = "create-post";
export const postMutationConfig: MutationConfig<PostMutationKeys> = {
  "create-post": {
    url: "/v1/posts/create",
    method: "POST",
    refetchQueries: ["posts"],
  },
};
export type PostMutationReturn = {
  "create-post": null;
};
