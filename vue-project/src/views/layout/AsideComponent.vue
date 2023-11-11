<template>
    <el-aside :style="{ width: asideStore.asideWidth + 'px' }" style="background: #A1C4C9;">
        <el-scrollbar>
            <el-menu style="width: auto;" :collapse="asideStore.isCollapse" unique-opened>
                <el-menu-item index="0" style="background-color: rgba(0, 0, 0, 0); 
                    letter-spacing: 30px; color: black; margin-top: 0; 
                    margin-left: 0;" @click="goHome">
                    <el-icon>
                        <HomeFilled />
                    </el-icon>
                    <span v-show="!asideStore.isCollapse">首页</span>
                </el-menu-item>
                <el-sub-menu index="1">
                    <template #title>
                        <el-icon>
                            <Cpu />
                        </el-icon>
                        <span v-show="!asideStore.isCollapse">系统管理</span>

                    </template>
                    <el-menu-item-group>
                        <!-- <template #title>Group 1</template> -->
                        <el-menu-item index="1-1">Option 1</el-menu-item>
                        <el-menu-item index="1-2">Option 2</el-menu-item>
                        <div style="height: 20px;"></div>
                    </el-menu-item-group>
                </el-sub-menu>
                <el-sub-menu index="2">
                    <template #title>
                        <el-icon>
                            <User />
                        </el-icon>
                        <!-- @click="userManage" -->
                        <span v-show="!asideStore.isCollapse">人员管理</span>
                    </template>
                    <el-menu-item-group>
                        <!-- <template #title>Group 1</template> -->
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
                        <div style="height: 20px;"></div>
                    </el-menu-item-group>
                </el-sub-menu>
                <el-sub-menu index="3">
                    <template #title>
                        <el-icon>
                            <ShoppingCartFull />
                        </el-icon>
                        <span v-show="!asideStore.isCollapse">订单管理</span>

                    </template>
                    <el-menu-item-group>
                        <!-- <template #title>Group 1</template> -->
                        <el-menu-item index="3-1">Option 1</el-menu-item>
                        <el-menu-item index="3-2">Option 2</el-menu-item>
                        <el-menu-item index="3-3">Option 3</el-menu-item>
                        <div style="height: 20px;"></div>

                    </el-menu-item-group>
                </el-sub-menu>
                <el-sub-menu index="4">
                    <template #title>
                        <el-icon>
                            <Grape />
                        </el-icon>
                        <span v-show="!asideStore.isCollapse">水果管理</span>

                    </template>
                    <el-menu-item-group>
                        <el-menu-item index="4-1">Option 1</el-menu-item>
                        <el-menu-item index="4-2">Option 2</el-menu-item>
                        <div style="height: 20px;"></div>

                    </el-menu-item-group>

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

}

</script>

<style>
/* .el-aside {
    background: linear-gradient(150deg, rgb(51, 137, 194), rgb(72, 209, 129));
} */

.layout-container-demo .el-aside {
    color: var(--el-text-color-primary);
    background: var(--el-color-primary-light-8);
}

.layout-container-demo .el-menu {
    border-right: none;
}

.el-menu {
    /* background-color: #A1C4C9; */
    background: linear-gradient(150deg, rgb(51, 137, 194), rgb(72, 209, 129));
}

/* 未激活的菜单项文本颜色 */
.el-menu-item {
    /* color: gray; */
    border-radius: 20px;
    background-color: #808AC9;
    /* margin: 2px 6px; */
    margin: 4px 10px 0 10px;

}

/* 悬停的菜单项文本颜色 */
.el-menu-item:hover {
    color: #FFF;
    background-color: #5BA9D1;
    border-radius: 20px;
    margin: 4px 10px 0 10px;
}

/* 激活的菜单项文本颜色 */
.el-menu-item.is-active {
    color: white;
    background-color: #FC803A;
    border-radius: 20px;
    margin: 4px 10px 0 10px;
}

.el-menu-item-group {
    background-color: #190625;
    color: #808AC9;
}

.el-sub-menu {
    background: linear-gradient(150deg, rgb(51, 137, 194), rgb(72, 209, 129));
}
</style>