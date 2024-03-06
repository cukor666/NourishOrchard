<template>
    <el-form class="user-info" label-width="100">
        <el-form-item label="账号：">
            <el-input v-model="user.username" disabled />
        </el-form-item>
        <el-form-item label="姓名：">
            <el-input v-model="user.name" @blur="validName(user.name)" />
        </el-form-item>
        <div v-if="errWord.name" class="error-word left-space">{{ errWord.name }}</div>
        <el-form-item label="性别：">
            <el-switch v-model="gender" class="ml-2" inline-prompt
                style="--el-switch-on-color: #2394f0; --el-switch-off-color: #FF6699; margin-left: 5px;" active-text="男"
                inactive-text="女" />
        </el-form-item>
        <el-form-item label="联系电话：">
            <el-input v-model="user.phone" @blur="validPhone(user.phone)" />
        </el-form-item>
        <div v-if="errWord.phone" class="error-word left-space">{{ errWord.phone }}</div>

        <el-form-item label="联系地址：">
            <el-input v-model="user.address" @blur="validAddress(user.address)" />
        </el-form-item>
        <div v-if="errWord.address" class="error-word left-space">{{ errWord.address }}</div>

        <el-form-item label="出生日期：">
            <el-date-picker v-model="user.birthday" type="date" @blur="validBirthday(user.birthday)" />
        </el-form-item>
        <div v-if="errWord.birthday" class="error-word left-space">{{ errWord.birthday }}</div>
    </el-form>

</template>

<script setup>
import { useUserInfoStore } from "@/stores/userInfo"
import { storeToRefs } from "pinia"
import { useValid } from '@/hooks/common/useValid'
import { reactive } from "vue"

const errWord = reactive({
    name: '',
    gender: '',
    phone: '',
    address: '',
    birthday: '',
})

const userInfoStore = useUserInfoStore()
const { user, gender } = storeToRefs(userInfoStore)

const { validName, validPhone, validAddress, validBirthday } = useValid(errWord)

</script>

<style lang="scss" scoped>
@import "@/scss/common/ErrorWord.scss";

.left-space {
    margin-top: -15px;
    margin-left: 100px;
}
</style>
