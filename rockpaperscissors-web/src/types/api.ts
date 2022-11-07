import { AxiosError } from "axios";

/**
 * Callback function for api methods
 */
export type CallbackFunction<T> = (
  data: T | null,
  error: Error | AxiosError<T> | null
) => void;
