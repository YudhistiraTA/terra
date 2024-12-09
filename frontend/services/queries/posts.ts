import type { QueryUrl } from "./queries";

export type PostQueryKeys = "posts";
export const postQueryUrl: QueryUrl<PostQueryKeys> = {
  posts: "/v1/posts/list",
};
export type PostQueryReturns = {
  posts: {
    posts: {
      id: string;
      title: string;
      content: string;
    }[];
    next_cursor?: string;
    previous_cursor?: string;
  };
  secondQuery: unknown;
};
