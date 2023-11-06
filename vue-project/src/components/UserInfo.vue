<template>
    <div>
        <el-drawer v-model="userStore.UserInfoVisible" title="用户个人信息" :before-close="handleClose">
            <h4>ID：{{ form.ID }}</h4>
            <h4>用户名：{{ form.name }}</h4>
            <h4>昵称：{{ form.nickname }}</h4>
            <h4>性别：{{ form.gender }}</h4>
            <h4>权限：{{ promise }}</h4>
            <h4>生日：{{ form.birthday }}</h4>
            <h4>联系方式：{{ form.phone }}</h4>
            <h4>家庭地址：{{ form.address }}</h4>
        </el-drawer>
    </div>
</template>

<script setup>
import { computed } from 'vue'
import { useUserStore } from '../stores/stores'
import router from '../router';

// 状态
const userStore = useUserStore()

// 数据
const form = userStore.userGridData

// 权限转换文字
const promise = computed(() => {
    switch (form.promise) {
        case 1:
            return '普通用户'
        case 2:
            return '管理员'
        case 3:
            return '超级管理员'
    }
})

// 关闭时触发
function handleClose() {
    userStore.UserInfoVisible = false
    userStore.userGridData = [{
        ID: 0,
        name: "xxx",
        nickname: "xxx",
        gender: "xxx",
        birthday: "xxx",
        phone: "xxx",
        address: "xxx",
        promise: 0
    }]
    router.back()
}

</script>
<style scoped>

h4 {
    margin-bottom: 10px;
}

.el-drawer {
    width: 300px;
}
</style>