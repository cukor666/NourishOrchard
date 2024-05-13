import {createRouter, createWebHistory} from "vue-router";

// 路由
const routes = [
    {
        path: "/",
        name: "Root",
        meta: {requireAuth: true, requireAdmin: false},
        redirect: "/home",
        component: () => import("@/views/Root.vue"),
        children: [
            {
                path: "/home",
                name: "Home",
                meta: {requireAuth: true, requireAdmin: false},
                component: () => import("@/views/home/Home.vue"),
            },
            {
                path: "/about",
                name: "About",
                meta: {requireAuth: false, requireAdmin: false},
                component: () => import("@/views/about/About.vue")
            },
            {
                path: "/user",
                name: "User",
                meta: {requireAuth: true, requireAdmin: true},
                children: [
                    {
                        path: "/user/list",
                        name: "UserList",
                        component: () => import("@/views/user/UserList.vue"),
                    },
                    {
                        path: "/user/logout-list",
                        name: "LogoutUserList",
                        component: () => import("@/views/user/LogoutUserList.vue")
                    }
                ],
            },
            {
                path: "/admin",
                name: "Admin",
                meta: {requireAuth: true, requireAdmin: true},
                children: [
                    {
                        path: "/admin/list",
                        name: "AdminList",
                        component: () => import("@/views/admin/AdminList.vue"),
                    },
                ],
            },
            {
                path: "/employee",
                name: "Employee",
                meta: {requireAuth: true, requireAdmin: true},
                children: [
                    {
                        path: "/employee/list",
                        name: "EmployeeList",
                        component: () => import("@/views/employee/EmployeeList.vue")
                    }
                ]
            },
            {
                path: "/fruit",
                name: "Fruit",
                meta: {requireAuth: true, requireAdmin: false},
                children: [
                    {
                        path: "/fruit/list",
                        name: "FruitList",
                        component: () => import("@/views/fruit/FruitList.vue"),
                    },
                ],
            },
            {
                path: "/warehouse",
                name: "Warehouse",
                meta: {requireAuth: true, requireAdmin: false},
                children: [
                    {
                        path: "/warehouse/warehouse-list",
                        name: "WarehouseList",
                        component: () => import("@/views/warehouse/WarehouseList.vue")
                    }
                ]
            },
            {
              path: "/inventory",
              name: "Inventory",
              meta: {requireAuth: true, requireAdmin: false},
              children: [
                  {
                      path: "/inventory-list",
                      name: "InventoryList",
                      component: () => import("@/views/inventory/InventoryList.vue")
                  }
              ]
            },
            {
                path: "/suppliers",
                name: "Suppliers",
                meta: {requireAuth: true, requireAdmin: true},
                children: [
                    {
                        path: "/suppliers-list",
                        name: "SuppliersList",
                        component: () => import("@/views/suppliers/SuppliersList.vue")
                    }
                ]
            },
            {
                path: "/orders",
                name: "Orders",
                meta: {requireAuth: true, requireAdmin: false},
                children: [
                    {
                        path: "/orders/list",
                        name: "OrdersList",
                        component: () => import("@/views/orders/OrdersList.vue")
                    }
                ]
            }
        ],
    },
    {
        path: "/login",
        name: "Login",
        meta: {requireAuth: false, requireAdmin: false},
        component: () => import("@/views/login/LoginView.vue"),
    },
    {
        path: "/register",
        name: "Register",
        meta: {requireAuth: false, requireAdmin: false},
        component: () => import("@/views/register/RegisterView.vue"),
    },
    {
        path: "/forget-password",
        name: "ForgetPassword",
        meta: {requireAuth: false, requireAdmin: false},
        component: () => import("@/views/login/ForgetPasswordView.vue"),
    },
    {
        path: "/promise-missed",
        name: "PromiseMissed",
        meta: {requireAuth: false, requireAdmin: false},
        component: () => import("@/views/PromiseMiss.vue"),
    }
];

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes,
});

// 白名单
const whiteList = ["/login", "/register", "/forget-password", "/about"];

// 前置导航守卫
router.beforeEach((to, from, next) => {
    // 白名单 直接进入
    let exists = whiteList.includes(to.path);
    if (exists) {
        return next();
    }
    // 验证是否登录
    if (to.meta.requireAuth) {
        // 获取token
        let token = localStorage.getItem("nourish-token");
        // 验证token
        let regStr = /^[A-Za-z0-9-_]+\.[A-Za-z0-9-_]+\.[A-Za-z0-9-_]*$/;
        let promise = sessionStorage.getItem("nourish-promise") || localStorage.getItem("nourish-promise");
        console.log('当前账户的权限为：',promise);
        if (!regStr.test(token)) {
            console.log("token格式不正确");
            return next({name: "Login"});
        } else {
            console.log('token格式正确')
            // token格式正确
            if (to.meta.requireAdmin) {
                if (to.name === "SuppliersList") {
                    return next()
                }
                // 需要管理员身份
                return promise === "admin" ? next() : next({name: "PromiseMissed"});
            } else {
                // 不需要管理员身份
                return next();
            }
        }
    }
    return next({name: "Login"});
});

export default router;
