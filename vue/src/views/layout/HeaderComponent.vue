<template>
    <el-header>
        <div class="title">
            <p>Nourish Orchard</p>
            <el-icon>
                <Watermelon />
            </el-icon>
        </div>
        <div class="user">
            <el-dropdown>
                <el-icon size="large" color="black">
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
            <span>欢迎 {{ userStore.loginUser.name }}</span>
        </div>
    </el-header>
</template>

<script setup>
import { useUserStore } from '@/stores/stores'
import { useRouter } from 'vue-router'
import request from '@/request';
import { ElMessage, ElMessageBox } from 'element-plus'

const userStore = useUserStore()
// 路由
const router = useRouter();

function showInfo() {
    try {
        request.get('/user/?name=' + userStore.loginUser.name).then(response => {
            if (response.code == 200) {
                console.log(response.data);
                // 希望出现一个模态对话框，而不是打印到控制台上
                userStore.UserInfoVisible = true
                userStore.tempUser = response.data
                router.push('/user/info')
            } else { ElMessage({ type: 'success', message: '无此用户' }) }
        }).catch(err => console.error(err))
    } catch (error) { console.error('An error occurred: ' + error); }
}

function updateInfo() { router.push('/user/update') }

// 退出登录，如果成功退出则注销用户，把对应的token删除
function exitLogin() {
    ElMessageBox.confirm(
        '您确定要退出登录吗?', 'Warning', {
        confirmButtonText: '是的',
        cancelButtonText: '取消',
        type: 'warning',
    }).then(response => {
        localStorage.removeItem('name')
        localStorage.removeItem('token')
        router.push('/login')
        ElMessage({ type: 'success', message: '退出登录成功，用户已注销' })
    }).catch(() => { ElMessage({ type: 'info', message: '取消退出登录' }) })
}
</script>

<style lang="scss" scoped>
.el-header {
    display: flex;
    align-items: center;
    font-size: 12px;
    margin-left: 0;
    padding-left: 0;

    .title {
        display: flex;
        align-items: center;
        flex: 1;
        width: 100%;
        font-size: 30px;
    }

    .user {
        display: flex;
        align-items: center;
        flex: 1;
        text-align: right;
        margin-right: auto;

        .el-dropdown {
            margin-top: 3px;
            margin-left: auto;
        }

        span {
            margin-left: 10px;
            font-size: large;
        }
    }
}
</style>