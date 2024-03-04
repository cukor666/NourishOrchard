import { createRouter, createWebHistory } from "vue-router";

// 路由
const routes = [
  {
    path: "/",
    name: "Root",
    meta: { requireAuth: true, requireAdmin: false },
    component: () => import("@/views/Root.vue"),
    children: [
      {
        path: "/home",
        name: "Home",
        meta: { requireAuth: true, requireAdmin: false },
        component: () => import("@/views/home/Home.vue"),
      },
      {
        path: "/user",
        name: "User",
        meta: { requireAuth: true, requireAdmin: true },
        children: [
          {
            path: "/user/list",
            name: "UserList",
            component: () => import("@/views/user/UserList.vue"),
          },
        ],
      },
      {
        path: "/admin",
        name: "Admin",
        meta: { requireAuth: true, requireAdmin: true },
        children: [
          {
            path: "/admin/list",
            name: "AdminList",
            component: () => import("@/views/admin/AdminList.vue"),
          },
        ],
      },
      {
        path: "/fruit",
        name: "Fruit",
        meta: { requireAuth: true, requireAdmin: false },
        children: [
          {
            path: "/fruit/list",
            name: "FruitList",
            component: () => import("@/views/fruit/FruitList.vue"),
          },
        ],
      },
    ],
  },
  {
    path: "/login",
    name: "Login",
    meta: { requireAuth: false, requireAdmin: false },
    component: () => import("@/views/login/LoginView.vue"),
  },
  {
    path: "/register",
    name: "Register",
    meta: { requireAuth: false, requireAdmin: false },
    component: () => import("@/views/register/RegisterView.vue"),
  },
  {
    path: "/forget-password",
    name: "ForgetPassword",
    meta: { requireAuth: false, requireAdmin: false },
    component: () => import("@/views/login/ForgetPasswordView.vue"),
  },
  {
    path: "/promise-missed",
    name: "PromiseMissed",
    meta: { requireAuth: false, requireAdmin: false },
    component: () => import("@/views/PromiseMiss.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

// 白名单
const whiteList = ["/login", "/register", "/forget-password"];

// 前置导航守卫
router.beforeEach((to, from, next) => {
  let exists = whiteList.includes(to.path);
  if (exists) {
    return next();
  }
  if (to.meta.requireAuth) {
    // 需要验证
    let token = localStorage.getItem("nourish-token");
    let regStr = /^[A-Za-z0-9-_]+\.[A-Za-z0-9-_]+\.[A-Za-z0-9-_]*$/;
    let promise = sessionStorage.getItem("nourish-promise");
    console.log(promise);
    if (!regStr.test(token)) {
      console.log("token格式不正确");
      return next({ name: "Login" });
    } else {
      // token格式正确
      if (to.meta.requireAdmin) {
        // 需要管理员身份
        return promise === "admin" ? next() : next({ name: "PromiseMissed" });
      } else {
        // 不需要管理员身份
        return next();
      }
    }
  }
  return next({ name: "Login" });
});

export default router;
