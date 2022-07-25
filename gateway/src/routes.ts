import { Options as ProxyOpts } from "http-proxy-middleware/dist/types";

interface IRoute {
  url: string;
  auth: boolean;
  proxy: ProxyOpts;
}

export const routes: IRoute[] = [
  {
    url: "/auth/login",
    auth: false,
    proxy: {
      target: `${process.env.AUTH_SVC_ADDRESS}/user/register`,
      changeOrigin: true,
      pathRewrite: {
        [`^/auth/login`]: "",
      },
    },
  },
  {
    url: "/auth/me",
    auth: true,
    proxy: {
      target: `${process.env.AUTH_SVC_ADDRESS}/user/who-am-i`,
      changeOrigin: true,
      pathRewrite: {
        [`^/auth/me`]: "",
      },
    },
  },
];
