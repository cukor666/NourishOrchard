<template>
    <el-aside :style="{ width: asideWidth + 'px' }"
        style="border-radius: 100px; margin: 50px auto; background: #161f2f;transition:1s">
        <el-scrollbar>
            <el-menu style="width: auto;" :collapse="isCollapse" unique-opened>
                <el-menu-item index="0" @click="toggleSidebar">
                    <img v-if="isCollapse" src="../../assets/笑脸.svg" width="25" alt="gg">
                    <img v-else src="../../assets/笑脸2.svg" width="25" style="margin: 0 auto;" alt="">
                </el-menu-item>
                <el-menu-item index="1" style=" margin-top: 0; margin-left: 0;" @click="goHome">
                    <el-icon>
                        <HomeFilled />
                    </el-icon>
                    <span v-show="!isCollapse">首&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;页</span>
                </el-menu-item>
                <el-sub-menu index="2">
                    <template #title>
                        <el-icon>
                            <Cpu />
                        </el-icon>
                        <span v-show="!isCollapse">系统管理</span>
                    </template>
                    <el-menu-item index="2-1">Option 1</el-menu-item>
                    <el-menu-item index="2-2">Option 2</el-menu-item>
                </el-sub-menu>
                <el-sub-menu index="3">
                    <template #title>
                        <el-icon>
                            <User />
                        </el-icon>
                        <span v-show="!isCollapse">人员管理</span>
                    </template>
                    <el-menu-item index="3-1" @click="userList">
                        <el-icon>
                            <UserFilled />
                        </el-icon>
                        用户列表
                    </el-menu-item>
                    <el-menu-item index="3-2" @click="adminList">
                        <el-icon>
                            <Avatar />
                        </el-icon>
                        管理员列表
                    </el-menu-item>
                    <el-menu-item index="3-3" @click="addAdmin">
                        <el-icon>
                            <Avatar />
                        </el-icon>
                        添加管理员
                    </el-menu-item>
                </el-sub-menu>
                <el-sub-menu index="4">
                    <template #title>
                        <el-icon>
                            <ShoppingCartFull />
                        </el-icon>
                        <span v-show="!isCollapse">订单管理</span>
                    </template>
                    <el-menu-item index="4-1">
                        订单列表
                    </el-menu-item>
                    <el-menu-item index="4-2">
                        Option 2
                    </el-menu-item>
                    <el-menu-item index="4-3">
                        Option 3
                    </el-menu-item>
                </el-sub-menu>
                <el-sub-menu index="5">
                    <template #title>
                        <el-icon>
                            <Grape />
                        </el-icon>
                        <span v-show="!isCollapse">产品管理</span>
                    </template>
                    <el-menu-item index="5-1" @click="fruitList">水果信息</el-menu-item>
                    <el-menu-item index="5-2">画廊</el-menu-item>
                </el-sub-menu>
                <el-sub-menu index="6">
                    <template #title>
                        <el-icon>
                            <Van />
                        </el-icon>
                        <span v-show="!isCollapse">采购管理</span>
                    </template>
                    <el-menu-item index="6-1">公司列表</el-menu-item>
                    <el-menu-item index="6-2">画廊</el-menu-item>
                </el-sub-menu>
            </el-menu>
        </el-scrollbar>
    </el-aside>
</template>

<script setup>
import { ref } from 'vue'
// import { useAsideStore } from '../../stores/aside'
import { useRouter } from 'vue-router'

const isCollapse = ref(true);
// 侧边栏宽度
const asideWidth = ref(64);

// 切换菜单的收缩状态
function toggleSidebar() {
    isCollapse.value = !isCollapse.value;
    asideWidth.value = isCollapse.value ? 64 : 200
}

// 侧边栏状态
// const asideStore = useAsideStore()

// 路由
const router = useRouter();

function goHome() {
    router.push('/')
}

function userList() {
    router.push('/user/list')
}

function adminList() {
    router.push('/admin/list')
}

function addAdmin() {
    router.push('/admin/add')
}

function fruitList() {
    router.push('/fruit/list')
}

</script>

<style scoped>
.el-aside {
    margin-right: 5px;
}

.layout-container-demo .el-aside {
    color: var(--el-text-color-primary);
    background: var(--el-color-primary-light-8);
}

.layout-container-demo .el-menu {
    border-right: none;
}

span {
    color: #f4cccc;
}

.el-icon {
    color: #f4cccc;
}

/* .el-menu {
    color: red;
} */

/* 未激活的菜单项文本颜色 */
.el-menu-item {
    color: #f4cccc;
    background-color: #161F2F;
}

/* 悬停的菜单项文本颜色 */
.el-menu-item:hover {
    background-color: #e06666;
}

/* 激活的菜单项文本颜色 */
.el-menu-item.is-active {
    color: #ffd966;
    background-color: #e06666;
}

.el-sub-menu {
    background: #0d0e0f;
}

.el-sub-menu:hover {
    background-color: #E06666;
}

/* .el-sub-menu__title {
    background-color: red !important;
}

.el-menu .el-menu-item:hover {
    background-color: aqua !important;
}

.el-menu .el-sub-menu {
    background: #0d0e0f !important;
}

.el-menu .el-sub-menu:hover {
    background: #ffd966 !important;
} */
</style>