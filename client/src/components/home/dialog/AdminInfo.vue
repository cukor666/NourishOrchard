<template>
    <el-form class="admin-info" label-width="100">
        <el-form-item label="账号：">
            <el-input v-model="admin.username" disabled />
        </el-form-item>

        <el-form-item label="姓名：">
            <el-input v-model="admin.name" @blur="validName(admin.name)" />
        </el-form-item>
        <div v-if="errWord.name" class="error-word left-space">{{ errWord.name }}</div>

        <el-form-item label="邮箱：">
            <el-input v-model="admin.email" @blur="validEmail(admin.email)" />
        </el-form-item>
        <div v-if="errWord.email" class="error-word left-space">{{ errWord.email }}</div>
    </el-form>
</template>

<script setup>
import { useAdminInfoStore } from '@/stores/adminInfo'
import { storeToRefs } from 'pinia'
import { reactive } from 'vue'
import { useValid } from '@/hooks/common/useValid'

const adminInfoStore = useAdminInfoStore()
const { admin } = storeToRefs(adminInfoStore)

const errWord = reactive({
    name: '',
    email: ''
})

const { validName, validEmail } = useValid(errWord)

</script>

<style lang="scss" scoped>
@import "@/scss/common/ErrorWord.scss";

.left-space {
    margin-top: -15px;
    margin-left: 100px;
}
</style>
