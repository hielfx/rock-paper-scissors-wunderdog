import axios, { AxiosError } from "axios";
import Swal, { SweetAlertIcon } from "sweetalert2";
import { BASE_URL } from "../constants";

/**
 * Axios custom instance
 */
const instance = axios.create({
  baseURL: BASE_URL,
  // headers: {
  //   "Content-type": "application/json",
  //   Accept: "application/json",
  // },
});

/**
 * Helper utility to display a swal error
 */
export const apiErrorHandler = (error: Error | AxiosError, title?: string) => {
  console.log("Error: ", error);
  let _error = error;
  let icon: SweetAlertIcon = "error";
  let _title = "Application error";

  if (axios.isAxiosError(error)) {
    console.log("is axiosError");
    _error = error.response?.data;
    const _status = error.response?.status as number;
    if (_status >= 500) {
      icon = "error";
      _title = "Application error";
    } else if (_status >= 400) {
      //TODO: Handle custom error messages from API
      _title = "Bad request";
      icon = "warning";
    } else if (_status >= 300) {
      _title = "Redirection error";
      icon = "info";
    }
  }
  Swal.fire({
    title: title || _title,
    text: _error?.message,
    icon,
  });
};

export default instance;
