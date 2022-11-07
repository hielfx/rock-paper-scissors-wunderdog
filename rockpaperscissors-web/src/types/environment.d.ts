export {};

declare global {
  namespace NodeJS {
    interface ProcessEnv {
      REACT_APP_API_URL: string;
      REACT_APP_WS_URL: string;
    }
  }
}
