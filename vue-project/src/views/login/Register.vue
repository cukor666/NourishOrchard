<template>
    <div class="container">
        <div class="register-box">
            <!-- 注册输入框 -->
            <div class="register-input">
                <p class="register-title gradient-text">注册</p>
                <el-form :model="user" :rules="rules" ref="ruleFormRef">
                    <el-form-item prop="name">
                        <el-input :prefix-icon="User" clearable placeholder="账号" v-model="user.name" />
                    </el-form-item>

                    <el-form-item prop="password">
                        <el-input :prefix-icon="Lock" clearable show-password placeholder="密码" v-model="user.password" />
                    </el-form-item>

                    <el-form-item prop="phone">
                        <el-input :prefix-icon="Phone" clearable placeholder="联系电话" v-model="user.phone" />
                    </el-form-item>

                    <el-form-item prop="birthday">
                        <el-date-picker style="width: 100%;padding: 4px 10px;" v-model="user.birthday" type="date"
                            placeholder="选择您的出生日期" />
                    </el-form-item>

                    <el-form-item prop="address" style="margin-bottom: 0px;">
                        <el-input :prefix-icon="Location" clearable placeholder="地址" v-model="user.address" />
                    </el-form-item>

                    <el-form-item style="margin: 0 25%;">
                        <el-radio-group v-model="user.gender">
                            <!-- 
                            label就是这个单选框的值，可以是数字类型也可以是字符串类型还可以是布尔值
                            如果是字符串类型，必须使用单引号和双引号交叉共同使用
                         -->
                            <el-radio :label='"男"'>男</el-radio>
                            <el-radio :label='"女"'>女</el-radio>
                        </el-radio-group>
                    </el-form-item>

                    <el-form-item>
                        <el-button type="primary" plain @click="joinOur" style="width: 100px;margin: 0 25%;">加入我们</el-button>
                    </el-form-item>
                </el-form>
            </div>
            <!-- 注册图片 -->
            <div class="register-image">
                <img src="../../assets/register.jpg" class="image" alt="">
            </div>
        </div>
    </div>
</template>

<script setup>
import { User, Lock, Location, Phone } from '@element-plus/icons-vue'
import { ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import request from "../../request";

// 路由
const router = useRouter();

// 表单模型
const user = ref({
    name: '',
    password: '',
    gender: '男',
    birthday: '',
    phone: '',
    address: ''
})

// 表单数据，最终会存放到这里
const ruleFormRef = ref(null);

// 校验规则
const rules = ref({
    name: [{ required: true, message: '请输入姓名', trigger: 'blur' }],
    password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
    birthday: [{ required: true, message: "请选择出生日期", trigger: 'blur' }],
    phone: [{ required: true, message: '请输入联系方式', trigger: 'blur' }],
    address: [{ required: true, message: '请输入地址', trigger: 'blur' }],
})

function joinOur() {
    // 表单验证
    // 当填完页面上的表单数据后，myForm就已经有值了
    myForm.value.validate((valid) => {
        console.log(valid);
        if (valid) {
            // 表单验证通过，可以在这里执行提交操作
            // console.log('表单验证通过，可以提交数据了', user.value);

            request.post('/addUser', user.value).then(response => {
                console.log(response.data);
                if (response.data.code === 200) {
                    alert('欢迎加入我们！！！')
                    router.push('/')
                } else {
                    alert('数据未填写完整，请检查')
                }
            }).catch(err => {
                console.log(err);
                alert('系统异常错误')
            })
        } else {
            console.log('表单验证失败');
        }
    })
}

</script>

<style scoped>
.container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    background: linear-gradient(150deg, rgb(204, 214, 135), rgb(72, 209, 129));
}

.register-box {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 400px;
    width: 600px;
    border-radius: 15px;
    background-color: rgba(79, 195, 208, 0.7);
    box-shadow: 66px 66px 100px 0px #e5aebc
}


.register-input {
    flex: 2;
    text-align: center;
}

.register-image {
    flex: 3;
}

.gradient-text {
    background-image: linear-gradient(to right, #e9ab2e, #df1acb);
    -webkit-background-clip: text;
    background-clip: text;
    color: transparent;
}

.image {
    width: 100%;
}

.register-title {
    font-size: 30px;
    font-weight: bold;
    width: 100%;
    margin-bottom: 10px;
    text-shadow: 3px 3px 4px rgba(13, 221, 48, 0.5);
}

.el-input {
    padding: 0 10px;
}

/* .el-radio {
    margin: 0 3px;
} */
</style>