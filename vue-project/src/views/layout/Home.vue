<template>
    <el-container class="layout-container-demo" style="height: 100vh">
        <el-aside :style="{ width: asideWidth + 'px' }"
            style="background: linear-gradient(150deg, rgb(51, 137, 194), rgb(72, 209, 129));" @open="handleOpen"
            @close="handleClose">
            <el-scrollbar>
                <el-menu style="width: auto;" :collapse="isCollapse" unique-opened>
                    <el-sub-menu index="1">
                        <template #title>
                            <el-icon>
                                <Cpu />
                            </el-icon>
                            <span v-show="!isCollapse">系统管理</span>

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
                            <span v-show="!isCollapse">人员管理</span>
                        </template>
                        <el-menu-item-group>
                            <!-- <template #title>Group 1</template> -->
                            <el-menu-item index="2-1" @click="userList">
                                <el-icon>
                                    <UserFilled />
                                </el-icon>
                                用户列表
                            </el-menu-item>
                            <el-menu-item index="2-2"><el-icon>
                                    <Avatar />
                                </el-icon>
                                管理员列表
                            </el-menu-item>
                            <div style="height: 20px;"></div>
                        </el-menu-item-group>
                    </el-sub-menu>
                    <el-sub-menu index="3">
                        <template #title>
                            <el-icon>
                                <ShoppingCartFull />
                            </el-icon>
                            <span v-show="!isCollapse">订单管理</span>

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
                            <span v-show="!isCollapse">水果管理</span>

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

        <el-container>
            <el-header style="font-size: 12px">
                <div class="toolbar">
                    <el-button @click="toggleSidebar" style="background-color: #37A2EA;color: white;margin-right: 10px;">
                        <el-icon v-show="isCollapse">
                            <DArrowRight />
                        </el-icon>
                        <el-icon v-show="!isCollapse">
                            <DArrowLeft />
                        </el-icon>
                    </el-button>
                    <div class="element1" style="width: 100%;flex: 1;font-size: 30px;">
                        <p class="element2">Nourish Orchard</p>
                        <el-icon class="element1">
                            <Watermelon />
                        </el-icon>
                    </div>
                    <div class="element2" style="flex: 1;text-align: right;">
                        <el-dropdown>
                            <el-icon style="margin-right: 8px; margin-top: 1px">
                                <setting />
                            </el-icon>
                            <template #dropdown>
                                <el-dropdown-menu>
                                    <el-dropdown-item @click="showInfo">个人信息</el-dropdown-item>
                                    <el-dropdown-item @click="updateInfo">更新信息</el-dropdown-item>
                                    <el-dropdown-item @click="exitLogin">退出登录</el-dropdown-item>
                                </el-dropdown-menu>
                            </template>
                        </el-dropdown>
                        <!-- 从pinia状态中获取 -->
                        <span>{{ userStore.loginUserName }}</span>
                    </div>
                </div>
            </el-header>
            <el-main>
                <router-view></router-view>
            </el-main>
        </el-container>
    </el-container>
</template>
  
<script setup>
import { ref, onMounted } from 'vue'
// import { Menu as IconMenu, Message, Setting } from '@element-plus/icons-vue'
import { useUserStore } from '../../stores/stores'
import { useRouter } from 'vue-router'
import request from '../../request';
import { ElMessage, ElMessageBox } from 'element-plus'


// 获取状态
const userStore = useUserStore()

// 路由
const router = useRouter();

// token
// const token = ref('')

// 在当前组件被挂载的时候自动执行
onMounted(() => {
    userStore.loginUserName = localStorage.getItem('name')  // 即使页面刷新，右上角的用户名依旧是登录的用户名
})

function showInfo() {
    // alert('按下了个人信息' + userStore.loginUserName)
    try {
        request.get('/user?name=' + userStore.loginUserName).then(response => {
            // 希望出现一个模态对话框，而不是打印到控制台上
            userStore.UserInfoVisible = true
            var user = response.data
            console.log(userStore.userGridData[0]);

            userStore.userGridData[0].ID = user.ID
            userStore.userGridData[0].name = user.name
            userStore.userGridData[0].gender = user.gender
            userStore.userGridData[0].birthday = user.birthday
            userStore.userGridData[0].phone = user.phone
            userStore.userGridData[0].address = user.address

            console.log(userStore.userGridData[0]);

            router.push('/user-info')
        }).catch(err => {
            console.error(err);
        })
    } catch (error) {
        console.error('An error occurred: ' + error);
    }
}

function updateInfo() {
    // alert('修改个人信息')
    router.push('/update-user-info')
}

// 退出登录，如果成功退出则注销用户，把对应的token删除
function exitLogin() {
    ElMessageBox.confirm(
        '您确定要退出登录吗?',
        'Warning',
        {
            confirmButtonText: '是的',
            cancelButtonText: '取消',
            type: 'warning',
        }
    ).then(response => {
        localStorage.removeItem('name')
        localStorage.removeItem('token')
        router.push('/')
        ElMessage({
            type: 'success',
            message: '退出登录成功，用户已注销'
        })
    }).catch(() => {
        ElMessage({
            type: 'info',
            message: '取消退出登录',
        })
    })

}

const asideWidth = ref(200)

const isCollapse = ref(false)

// 切换菜单的收缩状态
function toggleSidebar() {
    isCollapse.value = !isCollapse.value;
    asideWidth.value = isCollapse.value ? 64 : 200
}

const handleOpen = (key, keyPath) => {
    console.log(key, keyPath)
}
const handleClose = (key, keyPath) => {
    console.log(key, keyPath)
}

function userList() {
    router.push('/user-list')
}

</script>

<style scoped>
.layout-container-demo .el-header {
    position: relative;
    /* background-color: var(--el-color-primary-light-7); */
    color: var(--el-text-color-primary);
    background: linear-gradient(150deg, rgb(49, 164, 235), rgb(229, 93, 211));

}

.layout-container-demo .el-aside {
    color: var(--el-text-color-primary);
    background: var(--el-color-primary-light-8);
}

.layout-container-demo .el-menu {
    border-right: none;
}

.layout-container-demo .el-main {
    padding: 0;
}

.layout-container-demo .toolbar {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
    right: 20px;
}

.element1,
.element2 {
    display: inline;
    /* 或 display: inline-block; */
}

.el-main {
    background: linear-gradient(60deg, rgb(73, 185, 213), rgb(216, 64, 181));
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

    /* margin: 2px 6px; */
    margin: 4px 10px 0 10px;

}

/* 激活的菜单项文本颜色 */
.el-menu-item.is-active {
    /* color: rgb(191, 15, 194); */
    color: white;

    /* background-color: #5BA9D1; */
    background-color: #FC803A;

    border-radius: 20px;

    /* margin: 2px 6px; */
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
