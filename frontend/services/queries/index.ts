import {
  type HealthQueryKey,
  type HealthQueryReturn,
  healthQueryUrl,
} from "./health";

export type QueryKeys = HealthQueryKey;
export const queryUrl = {
  ...healthQueryUrl,
};
export type QueryReturns = HealthQueryReturn;
