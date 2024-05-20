import axios from "axios";
import {baseURL_dev, baseURL_pro} from "@/config/api.js"

const request = axios.create({
    // baseURL: baseURL_dev,
    baseURL: baseURL_pro,
    timeout: 5000,
});

request.interceptors.request.use(
    (config) => {
        config.headers["Content-Type"] = "application/json";
        config.headers["Accept"] = "application/json";
        const token = localStorage.getItem("nourish-orchard-user-token");
        if (token) {
            config.headers["Authorization"] = `Bearer ${token}`;
        }
        config.headers["username"] = localStorage.getItem("nourish-orchard-user-name");
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

request.interceptors.response.use(
    (response) => {
        if (response.data.code === 200) {
            console.log("响应成功：", response.data);
        } else if (/^4[0-9]+$/.test(response.data.code)) {
            console.log('参数错误');
        } else if (/^5[0-9]+$/.test(response.data.code)) {
            console.log('服务器错误');
        }
        return response.data;
    },
    (error) => {
        console.log("Response Error:", error);
        return Promise.reject(error);
    }
);

export default request;
