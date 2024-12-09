import {
  postMutationConfig,
  type PostMutationKeys,
  type PostMutationReturn,
} from "./posts";
import {
  userMutationConfig,
  type UserMutationKeys,
  type UserMutationReturn,
} from "./user";

export type MutationKeys = UserMutationKeys | PostMutationKeys;
export const mutationConfig = {
  ...userMutationConfig,
  ...postMutationConfig,
};
export type MutationReturns = {
  [k in UserMutationKeys]: UserMutationReturn[k];
} & {
  [k in PostMutationKeys]: PostMutationReturn[k];
};
