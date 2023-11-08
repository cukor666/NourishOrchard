<template>
    <div>
        <el-dialog v-model="dialogVisible" title="用户详细信息与编辑" width="45%" draggable @close="closeDialog">
            <template #footer>
                <el-form :model="form" label-width="100px" :rules="formRules" ref="myForm">
                    <el-form-item label="ID">
                        <el-input v-model="form.ID" disabled />
                    </el-form-item>

                    <!-- 
                        使用计算属性
                        1：普通用户
                        2：管理员
                        3；超级管理员
                     -->
                    <el-form-item label="权限">
                        <el-input v-model="userPromise" disabled />
                    </el-form-item>

                    <el-form-item label="用户名">
                        <el-input v-model="form.name" disabled />
                    </el-form-item>

                    <el-form-item label="昵称">
                        <el-input v-model="form.nickname" :disabled="isDisabled" />
                    </el-form-item>

                    <el-form-item label="性别">
                        <el-radio-group v-model="form.gender">
                            <el-radio label="男" />
                            <el-radio label="女" />
                        </el-radio-group>
                    </el-form-item>

                    <el-form-item label="出生日期" prop="birthday">
                        <el-col :span="11">
                            <el-date-picker v-model="form.birthday" type="date" placeholder="请选择您的出生日期"
                                style="width: 100%" />
                        </el-col>
                    </el-form-item>

                    <el-form-item label="联系电话" prop="phone">
                        <el-input v-model="form.phone" placeholder="请输入您的联系电话" />
                    </el-form-item>

                    <el-form-item label="地址" prop="address">
                        <el-input v-model="form.address" placeholder="请输入您的家庭住址" />
                    </el-form-item>
                </el-form>

                <el-button type="primary" :plain="true" :disabled="isDisabled" @click="submitDialog">
                    确认更改
                </el-button>
            </template>
        </el-dialog>
    </div>
</template>
  
<script setup>
import { ref, reactive, computed } from 'vue'
import { ElMessage } from 'element-plus'
import { useUserStore } from '../../stores/stores'
import router from '../../router';
import request from '../../request';

// 状态
const userStore = useUserStore()

// 对话框显示
const dialogVisible = ref(true)

// 表单赋值，内容从pinia状态中获取
const form = reactive({
    ID: userStore.userRow.ID,
    name: userStore.userRow.name,
    nickname: userStore.userRow.nickname,
    gender: userStore.userRow.gender,
    birthday: userStore.userRow.birthday,
    phone: userStore.userRow.phone,
    address: userStore.userRow.address,
    CreatedAt: userStore.userRow.CreatedAt,
    UpdatedAt: userStore.userRow.UpdatedAt,
    promise: userStore.userRow.promise
})


// 默认提交按钮不可点击
const isDisabled = computed(() => {
    switch (userStore.loginUserPromise) {   // 判断当前登录的用户的权限
        case 1: // 普通用户
            return !(form.name === userStore.loginUserName)
        case 2: // 管理员
        case 3: // 超级管理员
            return false
    }
})

// 用户权限
const userPromise = computed(() => {
    switch (form.promise) {
        case 1:
            return '普通用户'
        case 2:
            return '管理员'
        case 3:
            return '超级管理员'
    }
})


// 表单校验规则
const formRules = ref({
    birthday: [{ required: true, message: '出生日期必填', trigger: 'blur' }],
    phone: [{ required: true, message: '联系电话必填', trigger: 'blur' },
    { pattern: /^\d{6,20}$/, message: '联系电话格式不正确', trigger: 'blur' }],
    address: [{ required: true, message: '家庭地址必填', trigger: 'blur' }],
})

// 表单填写完成后会将数据保存在这里
const myForm = ref(null);

// 对话框关闭的时候会自动调用
function closeDialog() {
    dialogVisible.value = false
    router.back()
}

// 当提交表单时触发
function submitDialog() {
    // 提交表单
    // console.log(form);
    // console.log(userStore.userRow);
    myForm.value.validate((valid) => {
        if (valid) {
            try {
                request.put('/update-user-info', form).then(response => {
                    // console.log(response);
                    ElMessage({
                        message: '更新用户信息成功',
                        type: 'success',
                    })
                })
            } catch (error) {
                ElMessage({
                    message: '更新失败，出现异常',
                    type: 'error',
                })
                console.log(error);
            }
            dialogVisible.value = false
        } else {
            console.log('表单验证失败');
        }
    })
}


</script>
  