import {
  userMutationConfig,
  type UserMutationKeys,
  type UserMutationReturn,
} from "./user";

export type MutationKeys = UserMutationKeys;
export const mutationConfig = {
  ...userMutationConfig,
};
export type MutationReturns = UserMutationReturn;
