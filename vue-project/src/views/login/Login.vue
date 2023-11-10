<template>
    <div class="container">
        <div class="login-box">
            <!-- 图片 -->
            <div class="login-image">
                <img src="../../assets/login.jpg" class="image" alt="">
            </div>
            <!-- 登录输入框 -->
            <div class="login-input">
                <p class="login-title gradient-text">Nourish Orchard</p>
                <!-- 
                    :model="user" 绑定表单的用户数据
                    :rules="formRules" 校验规则
                    ref="myForm" 把表单的内容写到 myForm 中
                 -->
                <el-form :model="user" :rules="formRules" ref="myForm">
                    <!-- 
                        prop设置的要和v-model绑定的一致，当然如果使用的是ref，那就得点出来
                     -->
                    <el-form-item prop="name">
                        <el-input :prefix-icon="User" clearable placeholder="账号" v-model="user.name" />
                    </el-form-item>

                    <el-form-item prop="password">
                        <el-input :prefix-icon="Lock" clearable show-password placeholder="密码" v-model="user.password" />
                    </el-form-item>

                    <!-- 验证码图片 -->
                    <el-form-item prop="captcha">
                        <div style="display: flex; justify-content: center; margin-right: 10px;">
                            <el-input placeholder="验证码" v-model="user.captcha" style="flex: 1;" />
                            <img :src="captchaURL" @click="updateCaptcha" style="flex: 1;">
                        </div>
                    </el-form-item>

                    <el-form-item>
                        <div class="btn">
                            <el-button type="success" round style="flex: 1;" @click="goHome">登录</el-button>
                            <el-button round style="flex: 1;" @click="goRegister">注册</el-button>
                        </div>
                    </el-form-item>

                </el-form>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref } from "vue";
import { User, Lock } from '@element-plus/icons-vue'
import { useRouter } from 'vue-router'
import request from "../../request";
import { ElMessage } from 'element-plus'

// 登录用户对象
const user = ref({
    name: '',
    password: '',
    captcha: '',
})

// 验证码的图片路径，从后端传过来
const captchaURL = ref('http://localhost:9000/captcha')

// 更新验证码
function updateCaptcha() {
    captchaURL.value = 'http://localhost:9000/captcha?v=' + Math.random()
}

// 表单校验规则
const formRules = ref({
    name: [{ required: true, message: '请输入姓名', trigger: 'blur' }],
    password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
    captcha: [{ required: true, message: '请输入验证码', trigger: 'blur' }],
})

// 表单对象
const myForm = ref(null);

// 路由对象
const router = useRouter();

// 跳转到注册页面
function goRegister() {
    router.push('/register');
}

// 进入主页面
function goHome() {
    // 打印一下用户输入的验证码
    // console.log(user.value.captcha);

    // 表单验证
    myForm.value.validate((valid) => {
        if (valid) {
            // 表单验证通过
            ElMessage({
                message: '表单校验通过',
                type: 'success',
            })

            // 校验验证码
            request.post('/captcha/verify', {
                code: user.value.captcha    // 将前端输入的验证码发送给后端
            }).then(response => {
                // console.log(response);
                // 如果验证码校验通过了，才往下走
                if (response.code == 200) {
                    ElMessage({
                        message: '验证码正确',
                        type: 'success',
                    })
                    // 登录
                    // post的参数除了可以直接传user.value，
                    // 还可以这样传：{"name": user.value.name, "password": user.value.password}
                    request.post('/user-login', user.value).then(response => {
                        // console.log(response);  // data返回的是token
                        // 只有状态码是200才能登录系统，否则就不能登录
                        if (response.code === 200) {
                            // alert('登录成功')
                            localStorage.setItem('token', response.data)    // 将Token存放到本地存储上，这样后续可以通过localStorage.GetItem获取
                            localStorage.setItem('name', user.value.name)
                            router.push('/')    // 路由跳转就是通过这种类似压栈的
                        } else {
                            // alert('用户名或密码错误')
                            ElMessage({
                                message: '用户名或密码错误',
                                type: 'error',
                            })
                        }
                    }).catch(err => {
                        console.log(err);
                        // alert('系统异常错误')
                        ElMessage({
                            message: '系统异常错误',
                            type: 'error',
                        })
                    })
                } else {
                    // console.log('验证码不通过');
                    ElMessage({
                        message: '验证码不通过',
                        type: 'error',
                    })
                }
            }).catch(err => {
                console.log(err);
            })
        } else {
            // console.log('表单验证失败');
            ElMessage({
                message: '表单校验失败',
                type: 'error',
            })
        }
    });
}

</script>

<style scoped>
.container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    background: linear-gradient(60deg, rgb(73, 185, 213), rgb(216, 64, 181));
}

.login-box {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 320px;
    width: 600px;
    border-radius: 15px;
    background-color: rgba(54, 231, 69, 0.4);
    box-shadow: 66px 66px 100px 0px #3daaee
}

.login-image {
    flex: 1;
}

.image {
    width: 100%;
}

.login-input {
    flex: 1;
    text-align: center;
}

.login-title {
    font-size: 30px;
    font-weight: bold;
    width: 100%;
    margin-bottom: 10px;
    text-shadow: 3px 3px 4px rgba(13, 221, 48, 0.5);
}

.gradient-text {
    background-image: linear-gradient(to right, #16db44, #13e6ce);
    -webkit-background-clip: text;
    background-clip: text;
    color: transparent;
}

.el-input {
    padding: 5px 10px;
}

.btn {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 100%;
    margin-top: 5px;
    margin-left: 10px;
    margin-right: 10px;
}
</style>