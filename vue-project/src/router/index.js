import { createRouter, createWebHistory } from "vue-router";
import request from "../request";
// import { useUserStore } from "../stores/stores"

// const userStore = useUserStore()

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
              console.log('通过');
              next();
            } else {
              console.log('不通过');
              next(false);
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
  ],
});

export default router;
