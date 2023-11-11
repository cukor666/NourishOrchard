import { createRouter, createWebHistory } from "vue-router";
import request from "../request";
import { ElMessage } from "element-plus";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/login",
      name: "Login",
      component: () => import("../views/login/Login.vue"),
    },
    {
      path: "/register",
      name: "Register",
      component: () => import("../views/login/Register.vue"),
    },
    {
      path: "/",
      name: "Home",
      component: () => import("../views/layout/Home.vue"),
      beforeEnter(to, from, next) {
        request
          .get("/", {
            headers: {
              Authorization: "Bearer " + localStorage.getItem("token"),
            },
          })
          .then((response) => {
            if (response.code == 200) {
              next();
            } else {
              ElMessage({
                message: "请先登录",
                type: "error",
              });
              next("/login"); // 跳转到登录页面
            }
          })
          .catch((err) => {
            console.log(err);
          });
      },
      children: [
        {
          path: "/",
          name: "Welcome",
          component: () => import("../views/main/Welcome.vue"),
        },
        {
          path: "/user",
          name: "User",
          children: [
            {
              path: "/user/info",
              name: "UserInfo",
              component: () => import("../components/user/UserInfo.vue"),
            },
            {
              path: "/user/update",
              name: "UpdateUserInfo",
              component: () => import("../components/user/UpdateUserInfo.vue"),
            },
            {
              path: "/user/list",
              name: "/UserList",
              component: () => import("../views/main/UserView.vue"),
            },
            {
              path: "/user/edit-info",
              name: "/EditUserInfo",
              component: () => import("../components/user/EditUser.vue"),
            },
          ],
        },
        {
          path: "/admin",
          name: "Admin",
          children: [
            {
              path: "/admin/list",
              name: "/AdminList",
              component: () => import("../views/main/AdminView.vue"),
            },
            {
              path: "/admin/add",
              name: "AddAdmin",
              component: () => import("../components/admin/AddAdmin.vue"),
            },
          ],
        },
      ],
    },
    {
      path: "/see-use-info",
      name: "/SeeUseInfo",
      component: () => import("../components/user/SeeUseInfo.vue"),
    },
  ],
});

export default router;
