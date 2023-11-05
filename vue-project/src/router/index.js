import { createRouter, createWebHistory } from "vue-router";
import request from "../request";
import { ElMessage } from 'element-plus'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "Login",
      component: () => import("../views/login/Login.vue"),
    },
    {
      path: "/register",
      name: "Register",
      component: () => import("../views/login/Register.vue"),
    },
    {
      path: "/home",
      name: "Home",
      component: () => import("../views/layout/Home.vue"),
      beforeEnter(to, from, next) {
        request
          .get("/home", {
            headers: {
              Authorization: "Bearer " + localStorage.getItem("token"),
            },
          })
          .then((response) => {
            // console.log(response);
            if (response.code == 200) {
              // console.log("通过");
              ElMessage({
                message: "欢迎光临",
                type: "success",
              });
              next();
            } else {
              // console.log("不通过");
              ElMessage({
                message: "不通过",
                type: "error",
              });
              next('/');  // 跳转到登录页面
            }
          })
          .catch((err) => {
            console.log(err);
          });
      },
      children: [
        {
          path: "/user",
          name: "User",
          component: () => import("../views/main/UserView.vue"),
        },
        {
          path: "/admin",
          name: "Admin",
          component: () => import("../views/main/AdminView.vue"),
        },
        {
          path: "/user-info",
          name: "UserInfo",
          component: () => import("../components/UserInfo.vue"),
        },
        {
          path: "/update-user-info",
          name: "UpdateUserInfo",
          component: () => import("../components/UpdateUserInfo.vue"),
        },
        {
          path: "/user-list",
          name: "/UserList",
          component: () => import("../views/main/UserView.vue"),
        },
        {
          path: "/edit-user-info",
          name: "/EditUserInfo",
          component: () => import("../components/EditUser.vue"),
        },
      ],
    },
    {
      path: '/see-use-info',
      name: "/SeeUseInfo",
      component: () => import("../components/SeeUseInfo.vue")
    }
  ],
});

export default router;
