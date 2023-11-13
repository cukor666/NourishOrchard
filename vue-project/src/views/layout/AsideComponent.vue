<template>
    <el-aside :style="{ width: asideStore.asideWidth + 'px' }"
        style="border-radius: 50px; margin: 50px auto; background: #161f2f;transition:1s">
        <el-scrollbar>
            <el-menu style="width: auto;" :collapse="asideStore.isCollapse" unique-opened>
                <el-menu-item index="0" style=" margin-top: 0; margin-left: 0;" @click="goHome">
                    <el-icon v-show="asideStore.isCollapse">
                        <HomeFilled />
                    </el-icon>
                    <span v-show="!asideStore.isCollapse"
                        style="margin: 0 auto;">首&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;页</span>
                </el-menu-item>
                <el-sub-menu index="1">
                    <template #title>
                        <el-icon>
                            <Cpu />
                        </el-icon>
                        <span v-show="!asideStore.isCollapse">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;系统管理</span>
                    </template>
                    <el-menu-item index="1-1">Option 1</el-menu-item>
                    <el-menu-item index="1-2">Option 2</el-menu-item>
                </el-sub-menu>
                <el-sub-menu index="2">
                    <template #title>
                        <el-icon>
                            <User />
                        </el-icon>
                        <!-- @click="userManage" -->
                        <span v-show="!asideStore.isCollapse">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;人员管理</span>
                    </template>
                    <el-menu-item index="2-1" @click="userList">
                        <el-icon>
                            <UserFilled />
                        </el-icon>
                        用户列表
                    </el-menu-item>
                    <el-menu-item index="2-2" @click="adminList">
                        <el-icon>
                            <Avatar />
                        </el-icon>
                        管理员列表
                    </el-menu-item>
                    <el-menu-item index="2-3" @click="addAdmin">
                        <el-icon>
                            <Avatar />
                        </el-icon>
                        添加管理员
                    </el-menu-item>
                </el-sub-menu>
                <el-sub-menu index="3">
                    <template #title>
                        <el-icon>
                            <ShoppingCartFull />
                        </el-icon>
                        <span v-show="!asideStore.isCollapse">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;订单管理</span>
                    </template>
                    <el-menu-item index="3-1">订单列表</el-menu-item>
                    <el-menu-item index="3-2">Option 2</el-menu-item>
                    <el-menu-item index="3-3">Option 3</el-menu-item>
                </el-sub-menu>
                <el-sub-menu index="4">
                    <template #title>
                        <el-icon>
                            <Grape />
                        </el-icon>
                        <span v-show="!asideStore.isCollapse">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;水果管理</span>
                    </template>
                    <el-menu-item index="4-1">水果信息</el-menu-item>
                    <el-menu-item index="4-2">画廊</el-menu-item>
                </el-sub-menu>
                <el-sub-menu index="5">
                    <template #title>
                        <el-icon>
                            <Van />
                        </el-icon>
                        <span v-show="!asideStore.isCollapse">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;采购管理</span>
                    </template>
                    <el-menu-item index="5-1">公司列表</el-menu-item>
                    <el-menu-item index="5-2">画廊</el-menu-item>
                </el-sub-menu>
            </el-menu>
        </el-scrollbar>
    </el-aside>
</template>

<script setup>
import { useAsideStore } from '../../stores/aside'
import { useRouter } from 'vue-router'

// 侧边栏状态
const asideStore = useAsideStore()

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

</script>

<style scoped>
.el-aside {
    /* background: linear-gradient(150deg, rgb(51, 137, 194), rgb(72, 209, 129)); */
    margin-right: 5px;
    /* margin-left: 5px; */
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

.el-menu {
    /* background-color: #A1C4C9; */
    /* background: linear-gradient(150deg, rgb(51, 137, 194), rgb(72, 209, 129)); */
    color: red;
}

/* 未激活的菜单项文本颜色 */
.el-menu-item {
    color: #f4cccc;
    /* border-radius: 20px; */
    background-color: #161F2F;
    /* margin: 2px 6px; */
    /* margin: 4px 10px 0 10px; */
}

/* 悬停的菜单项文本颜色 */
.el-menu-item:hover {
    /* color: #FFF; */
    background-color: #e06666;
    /* border-radius: 20px; */
    /* margin: 4px 10px 0 10px; */
}

/* 激活的菜单项文本颜色 */
.el-menu-item.is-active {
    color: #ffd966;
    background-color: #e06666;
    /* border-radius: 20px; */
    /* margin: 4px 10px 0 10px; */
}

.el-sub-menu {
    /* background: linear-gradient(150deg, rgb(51, 137, 194), rgb(72, 209, 129)); */
    background: #0d0e0f;
}
</style>