<template>
    <div class="container">
        <div class="path__content">
            <span style="font-size: 14px; font-weight: bold;">首页</span>
        </div>
        <div class="setting-bar">
            <el-dropdown :hide-on-click="false">
                <div class="avatar"></div>
                <template #dropdown>
                    <el-dropdown-menu>
                        <el-dropdown-item @click="selfInfo">个人信息</el-dropdown-item>
                        <el-dropdown-item @click="exitDialogVisiable = true">退出登录</el-dropdown-item>
                    </el-dropdown-menu>
                </template>
            </el-dropdown>
        </div>
    </div>
    <div class="tags">
        <el-tag v-for="(tag, index) in tags" :key="tag.id" closable type="primary" @close="removeTag(index)">
            {{ tag.name }}
        </el-tag>
    </div>

    <el-dialog v-model="selfInfoDialogVisiable" title="个人信息" width="500" align-center @close="infoDialogCancel">
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
        <template #footer>
            <div class="dialog-footer">
                <el-button @click="infoDialogCancel">取消</el-button>
                <el-button type="primary" @click="updateUser">
                    更新
                </el-button>
            </div>
        </template>
    </el-dialog>

    <el-dialog v-model="exitDialogVisiable" title="退出" width="300" align-center>
        <span>你确定要退出吗？</span>
        <template #footer>
            <div class="dialog-footer">
                <el-button @click="exitDialogVisiable = false">取消</el-button>
                <el-button type="primary" @click="exit">确定</el-button>
            </div>
        </template>
    </el-dialog>
</template>

<script setup>
import { reactive, ref } from "vue";
import { useRouter } from 'vue-router'
import { useUserInfo } from '@/hooks/header/useUserInfo'
import { useValid } from '@/hooks/common/useValid'

import request from '@/axios/request'
import { ElMessage } from "element-plus";
const router = useRouter()

const tags = ref([
    {
        id: 1,
        name: '首页',
        path: '/'
    },
    {
        id: 2,
        name: '用户列表',
        path: '/user/list'
    },
    {
        id: 3,
        name: '管理员列表',
        path: '/admin/list'
    },
    {
        id: 4,
        name: '水果列表',
        path: '/fruit/list'
    }
])

const removeTag = (index) => {
    tags.value.splice(index, 1)
}

const user = reactive({
    username: "",
    name: "",
    gender: "",
    phone: "",
    address: "",
    birthday: "",
});

const errWord = reactive({
    name: '',
    gender: '',
    phone: '',
    address: '',
    birthday: ''
})

const { validName, validGender, validPhone, validAddress, validBirthday } = useValid(errWord)

const valid = () => {
    validName(user.name)
    validGender(user.gender)
    validPhone(user.phone)
    validAddress(user.address)
    validBirthday(user.birthday)
}

const { selfInfoDialogVisiable, gender, updateUserInfo } = useUserInfo(user)

const infoDialogCancel = () => {
    for (const key in errWord) {
        errWord[key] = ''
    }
    // console.log(errWord);
    selfInfoDialogVisiable.value = false
}

const updateUser = () => {
    valid()
    for (const key in errWord) {
        if (errWord[key] !== '') {
            ElMessage({
                message: '用户信息不正确',
                type: 'error'
            })
            return
        }
    }
    updateUserInfo()
}

const selfInfo = () => {
    // 向后端请求数据
    request.get('/account/get').then(res => {

        if (res.code === 200) {
            user.username = res.data.username
            user.name = res.data.name
            user.gender = res.data.gender
            user.phone = res.data.phone
            user.address = res.data.address
            user.birthday = res.data.birthday
            gender.value = user.gender === '男'
        } else {
            ElMessage({
                message: '请求失败，请退出重新登录',
                type: 'error'
            })
            localStorage.removeItem('nourish-token')
            localStorage.removeItem('nourish-account')
            router.replace({ name: 'Login' })
        }
    }).catch(err => {
        console.error('请求失败：', err)
    })

    selfInfoDialogVisiable.value = true
}

const exitDialogVisiable = ref(false)

const exit = () => {
    localStorage.removeItem('nourish-token')
    localStorage.removeItem('nourish-account')
    exitDialogVisiable.value = false
    router.push({ name: "Login" })
}


</script>

<style lang="scss" scoped>
@import "@/scss/home/header/header.scss";
@import "@/scss/common/ErrorWord.scss";

.left-space {
    margin-left: 100px;
    margin-top: -18px;
}
</style>