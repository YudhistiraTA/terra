import type { QueryUrl } from "./queries";

export type HealthQueryKey = "health";
export const healthQueryUrl: QueryUrl<HealthQueryKey> = {
  health: "/v1/health",
};
export type HealthQueryReturn = {
  health: { message: string };
};
