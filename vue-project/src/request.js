// 对axios进行二次封装
import axios from "axios";
import router from "./router";
// 创建一个自定义 Axios 实例并命名为 'request'
const request = axios.create({
  // 前端处理跨域 设置基础URL，在前端层面已经设置了代理，但要求后端的路由要加上 /api 作为前缀
  // 后端处理跨域 不需要加/api前缀
  baseURL: "http://localhost:9000",
  timeout: 5000, // 设置请求超时时间
  headers: {
    "Content-Type": "application/json;charset=UTF-8;",
  },
});

// 请求拦截器
// request.interceptors.request.use(
//   (config) => {
//     // 可以在这里添加请求头或进行其他请求前的处理
//     return config;
//   },
//   (error) => {
//     return Promise.reject(error);
//   }
// );

// // 响应拦截器
// request.interceptors.response.use(
//   (response) => {
//     const { code, msg, data } = response.data;
//     if (code === 200) {
//       console.log("请求成功");
//       console.log("msg: " + msg);
//       return data;
//     } else if (code === 400) {
//       console.log("请求失败");
//       console.log("msg: " + msg);
//     }
//   }
// );

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

// // 用于发送 HTTP 请求
// function requestMethods(options) {
//     options.method = options.method || 'get'
//     if (options.method.toLowerCase() === 'get') {
//         // 这是为了将请求参数放在 URL 上，以便发送 GET 请求。
//         options.params = options.data
//     }
//     return request(options) // 返回发送请求的结果。
// }

// ['get', 'post', 'put', 'delete'].forEach(item => {
//     requestMethods[item] = (url, data) => {
//         return request({
//             url,
//             data,
//             method: item
//         })
//     }
// });

export default request; // 导出自定义 Axios 实例并命名为 'request'
