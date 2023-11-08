<template>
    <el-header style="font-size: 12px">
        <div class="toolbar">
            <!-- 收缩和展开侧边栏的按钮 -->
            <el-button @click="toggleSidebar" style="background-color: #B3CFDA;color: black;margin-right: 10px;">
                <el-icon v-show="asideStore.isCollapse">
                    <DArrowRight />
                </el-icon>
                <el-icon v-show="!asideStore.isCollapse">
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
                    <el-icon size="large" color="black" style="margin-right: 8px; margin-top: 3px">
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
                <span style="font-size: large;">欢迎 {{ userStore.loginUserName }}</span>
            </div>
        </div>
    </el-header>
</template>

<script setup>
import { useUserStore } from '../../stores/stores'
import { useAsideStore } from '../../stores/aside';
import { useRouter } from 'vue-router'
import request from '../../request';
import { ElMessage, ElMessageBox } from 'element-plus'

// 侧边栏状态
const asideStore = useAsideStore()
const userStore = useUserStore()
// 路由
const router = useRouter();

// 切换菜单的收缩状态
function toggleSidebar() {
    asideStore.isCollapse = !asideStore.isCollapse;
    asideStore.asideWidth = asideStore.isCollapse ? 64 : 200
}


function showInfo() {
    // alert('按下了个人信息' + userStore.loginUserName)
    try {
        request.get('/user?name=' + userStore.loginUserName).then(response => {
            if (response.code == 200) {
                // console.log(response.data);

                // 希望出现一个模态对话框，而不是打印到控制台上
                userStore.UserInfoVisible = true
                var user = response.data
                // console.log(userStore.userGridData);

                // 这个地方待优化
                userStore.userGridData.ID = user.ID
                userStore.userGridData.name = user.name
                userStore.userGridData.nickname = user.nickname
                userStore.userGridData.gender = user.gender
                userStore.userGridData.birthday = user.birthday
                userStore.userGridData.phone = user.phone
                userStore.userGridData.address = user.address
                userStore.userGridData.promise = user.promise

                // console.log(userStore.userGridData);

                router.push('/user-info')
            } else {
                alert('无此用户')
            }
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


</script>

<style>

.layout-container-demo .el-header {
    position: relative;
    /* background-color: var(--el-color-primary-light-7); */
    color: var(--el-text-color-primary);
    /* background: linear-gradient(150deg, rgb(49, 164, 235), rgb(229, 93, 211)); */
    background: linear-gradient(150deg, rgb(200, 211, 218), rgb(14, 175, 219), rgb(184, 44, 161));

}

.element1,
.element2 {
    display: inline;
    /* 或 display: inline-block; */
}

.layout-container-demo .toolbar {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
    right: 20px;
}

</style>