<template>
    <div>
        <el-dialog v-model="dialogVisible" title="更新个人信息" width="45%" draggable @close="closeDialog">
            <template #footer>
                <el-form :model="form" label-width="100px" :rules="formRules" ref="myForm">
                    <el-form-item label="用户名">
                        <el-input v-model="form.name" disabled />
                    </el-form-item>

                    <el-form-item label="昵称">
                        <el-input v-model="form.nickname" placeholder="默认不修改昵称" />
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

                <el-button type="primary" @click="submitDialog">
                    确认提交
                </el-button>
            </template>
        </el-dialog>
    </div>
</template>
  
<script setup>
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'
import { useUserStore } from '../../stores/stores'
import router from '../../router';
import request from '../../request';

const userStore = useUserStore()

const dialogVisible = ref(true)

// 表单数据
const form = reactive({
    name: userStore.loginUserName,
    nickname: "",
    gender: "男",
    birthday: "",
    phone: "",
    address: "",
})


// 表单校验规则
const formRules = ref({
    birthday: [{ required: true, message: '出生日期必填', trigger: 'blur' }],
    phone: [{ required: true, message: '联系电话必填', trigger: 'blur' },
    { pattern: /^\d{6,20}$/, message: '联系电话格式不正确', trigger: 'blur' }],
    address: [{ required: true, message: '家庭地址必填', trigger: 'blur' }],
})

// 表单填写完成后，内容会存放到这里
const myForm = ref(null);

// 当对话框关闭时会触发
function closeDialog() {
    dialogVisible.value = false
    router.back()
}

// 当提交表单时会触发
function submitDialog() {
    // 提交表单
    // console.log(form);
    myForm.value.validate((valid) => {
        if (valid) {
            try {
                request.put('/user/update', form).then(response => {
                    if (response.code == 200) {
                        ElMessage({
                            message: '更新用户信息成功',
                            type: 'success',
                        })
                    } else {
                        console.log(response);
                        ElMessage({
                            message: '更新用户信息失败',
                            type: 'error',
                        })
                    }
                })
            } catch (error) {
                ElMessage({
                    message: '更新用户信息成功',
                    type: 'error',
                })
                console.log(error);
            }
            dialogVisible.value = false
            // router.back()
            router.push('/home')
        } else {
            console.log('表单验证失败');
        }
    })
}


</script>
  