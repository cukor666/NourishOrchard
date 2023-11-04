<template>
    <div>
        <el-dialog v-model="userStore.UserInfoVisible" title="个人信息" @close="closeDialog">
            <el-table :data="userStore.userGridData">
                <el-table-column property="ID" label="ID" width="50" />
                <el-table-column property="name" label="用户名" width="80" />
                <el-table-column property="gender" label="性别" width="60" />
                <el-table-column property="promise" label="权限" width="60" />
                <el-table-column property="birthday" label="生日" width="220" />
                <el-table-column property="phone" label="联系方式" width="120" />
                <el-table-column property="address" label="地址" width="200" />
            </el-table>
        </el-dialog>
    </div>
</template>

<script setup>
import { useUserStore } from '../stores/stores'
import router from '../router';

const userStore = useUserStore()

function closeDialog() {
    userStore.UserInfoVisible = false
    userStore.userGridData = [{
        ID: 0,
        name: "xxx",
        gender: "xxx",
        birthday: "xxx",
        phone: "xxx",
        address: "xxx",
        promise: "xxx"
    }]
    router.back()
}

</script>
<style scoped>
.el-button--text {
    margin-right: 15px;
}

.el-select {
    width: 300px;
}

.el-input {
    width: 300px;
}

.dialog-footer button:first-child {
    margin-right: 10px;
}
</style>