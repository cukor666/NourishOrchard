import axios from "axios";

const request = axios.create({
  baseURL: "http://localhost:9000",
  timeout: 5000,
});

request.interceptors.request.use(
  (config) => {
    config.headers["Content-Type"] = "application/json";
    config.headers["Accept"] = "application/json";
    const token = localStorage.getItem("nourish-token");
    if (token) {
      config.headers["Authorization"] = `Bearer ${token}`;
    }
    config.headers["username"] = localStorage.getItem("nourish-account")
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

request.interceptors.response.use(
  (response) => {
    console.log('响应成功：',response);
    return response.data;
  },
  (error) => {
    console.log("Response Error:", error);
    return Promise.reject(error);
  }
);

export default request;
