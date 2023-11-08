<template>
    <el-container class="layout-container-demo">
        <!-- 侧边栏 -->
        <aside-component />

        <el-container class="header-main">
            <HeaderComponent />
            <el-main>
                <router-view />
            </el-main>
        </el-container>
    </el-container>
</template>
  
<script setup>
import { onMounted } from 'vue'
import { useUserStore } from '../../stores/stores'
import request from '../../request';
import AsideComponent from './AsideComponent.vue';
import HeaderComponent from './HeaderComponent.vue';

// 获取状态
const userStore = useUserStore()

// 在当前组件被挂载的时候自动执行
onMounted(() => {
    userStore.loginUserName = localStorage.getItem('name')  // 即使页面刷新，右上角的用户名依旧是登录的用户名
    request.get('/user-promise/' + userStore.loginUserName).then(response => {
        // console.log(response);
        if (response.code == 200) {
            userStore.loginUserPromise = response.data
        } else {
            userStore.loginUserPromise = -1
            console.log('后端错误');
        }
    }).catch(err => {
        console.log(err);
    })
})

</script>

<style scoped>
.header-main {
    display: flex;
    flex-direction: column;
}

.layout-container-demo {
    height: 100vh
}


.layout-container-demo .el-main {
    padding: 0;
}

.el-main {
    background: linear-gradient(60deg, rgb(73, 185, 213), rgb(216, 64, 181));
}
</style>
