// 对axios进行二次封装
import axios from "axios";
import router from "./router";
import { getServer } from "./config";

// 创建一个自定义 Axios 实例并命名为 'request'
const request = axios.create({
  // 前端处理跨域 设置基础URL，在前端层面已经设置了代理，但要求后端的路由要加上 /api 作为前缀
  // 后端处理跨域 不需要加/api前缀
  baseURL: getServer(),
  timeout: 5000, // 设置请求超时时间
  headers: {
    "Content-Type": "application/json;charset=UTF-8;",
  },
});

request.interceptors.response.use(
  (response) => {
    let res = response.data;
    if (typeof res === "string") {
      res = res ? JSON.parse(res) : res;
    }
    if (res.code === "400") {
      router.push("/");
    }
    return res;
  },
  (error) => {
    console.log("err" + error);
    return Promise.reject(error);
  }
);

export default request; // 导出自定义 Axios 实例并命名为 'request'
