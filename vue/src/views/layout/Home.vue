<template>
    <el-container class="layout-container-demo">
        <!-- 侧边栏 -->
        <aside-component />
        <el-container class="header-main">
            <header-component />
            <el-main>
                <router-view />
            </el-main>
        </el-container>
    </el-container>
</template>
  
<script setup>
import { onMounted } from 'vue'
import { useUserStore } from '@/stores/stores'
import request from '@/request';
import AsideComponent from './AsideComponent.vue';
import HeaderComponent from './HeaderComponent.vue';

// 获取状态
const userStore = useUserStore()

// 在当前组件被挂载的时候自动执行
onMounted(() => {
    userStore.loginUser.name = localStorage.getItem('name')  // 即使页面刷新，右上角的用户名依旧是登录的用户名
    request.get('/promise/' + userStore.loginUser.name).then(response => {
        if (response.code == 200) {
            userStore.loginUser.promise = response.data
        } else {
            userStore.loginUser.promise = -1
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
</style>
