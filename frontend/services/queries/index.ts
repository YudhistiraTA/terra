import {
  type HealthQueryKey,
  type HealthQueryReturn,
  healthQueryUrl,
} from "./health";
import {
  postQueryUrl,
  type PostQueryKeys,
  type PostQueryReturns,
} from "./posts";

export type QueryKeys = HealthQueryKey | PostQueryKeys;
export const queryUrl = {
  ...healthQueryUrl,
  ...postQueryUrl,
};

export type QueryReturns = {
  [K in HealthQueryKey]: HealthQueryReturn[K];
} & {
  [K in PostQueryKeys]: PostQueryReturns[K];
};
