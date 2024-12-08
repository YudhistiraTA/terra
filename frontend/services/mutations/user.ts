import { MutationConfig } from "./mutations";

export type UserMutationKeys = "login";
export const userMutationConfig: MutationConfig<UserMutationKeys> = {
  login: { url: "/v1/user/login", method: "POST" },
};
export type UserMutationReturn = {
  login: {
    message: string;
  };
};
