<template>
    <div class="container">
        <div class="register-box">
            <!-- 注册输入框 -->
            <div class="register-input">
                <p class="register-title gradient-text">注册</p>
                <el-form :model="user" :rules="rules" ref="ruleFormRef">
                    <el-form-item prop="name">
                        <el-input :prefix-icon="User" clearable placeholder="用户名" v-model="user.name" />
                    </el-form-item>

                    <el-form-item prop="nickname">
                        <el-input :prefix-icon="Sugar" clearable placeholder="昵称" v-model="user.nickname" />
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

                    <el-form-item style="display: flex; justify-content: center; margin-left: 10px; margin-right: 10px;">
                        <el-button type="primary" plain round @click="joinOur" style="flex: 1;">
                            加入我们
                        </el-button>
                        <el-button style="flex: 1;" round @click="drawer = true">
                            帮助
                        </el-button>
                        <el-drawer v-model="drawer" title="I am the title" :with-header="false" style="text-align: left;">
                            <h2>注册表单帮助</h2>
                            <p>表单上的所有输入框都是<span>必填</span>的。</p>
                            <h3>用户名</h3>
                            <ol>
                                <li>用户名<span>不可以重复</span>。</li>
                                <li>用户名的格式必须是<span>数字、字母、下划线</span>组成。</li>
                                <li>必须以<span>字母开头</span>。</li>
                            </ol>
                            <h3>密码</h3>
                            <p>密码至少<span>6</span>位。</p>
                            <h3>联系电话</h3>
                            <p>必须是<span>数字</span>，并且在<span>6~20</span>位之间。</p>
                            <h3>家庭地址</h3>
                            <p><span>不能为空</span>，越详细我们更容易找到您。</p>
                            <h3>性别</h3>
                            <p><span>默认是男</span></p>
                            <h3>完成</h3>
                            <p>信息都填写完毕后就可以点击“加入我们”按钮完成注册。</p>
                        </el-drawer>
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
import { User, Lock, Location, Phone, Sugar } from '@element-plus/icons-vue'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import request from "../../request";
import { ElMessage } from 'element-plus'

// 路由
const router = useRouter();

// 注册帮助
const drawer = ref(false)

// 表单模型
const user = ref({
    name: '',
    nickname: '',
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
    name: [{ required: true, message: '请输入用户名', trigger: 'blur' },
    { pattern: /^[a-zA-Z]\w{1,19}$/, message: '用户名格式错误', trigger: 'blur' }],
    nickname: [{ required: true, message: '请输入昵称', trigger: 'blur' }],
    password: [{ required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, max: 20, message: '密码格式不正确', trigger: 'blur' }],
    birthday: [{ required: true, message: "请选择出生日期", trigger: 'blur' }],
    phone: [{ required: true, message: '请输入联系方式', trigger: 'blur' },
    { pattern: /^\d{6,20}$/, message: '联系电话格式不正确', trigger: 'blur' }],
    address: [{ required: true, message: '请输入地址', trigger: 'blur' }]
})

// 加入我们，注册成功
function joinOur() {
    // 表单验证
    // 当填完页面上的表单数据后，ruleFormRef就已经有值了
    ruleFormRef.value.validate((valid) => {
        // console.log(valid);
        if (valid) {
            // 表单验证通过，可以在这里执行提交操作
            ElMessage({
                message: '表单校验通过',
                type: 'success',
            })
            request.post('/user/add', user.value).then(response => {
                if (response.code === 200) {
                    ElMessage({
                        message: '欢迎加入我们！！！',
                        type: 'success',
                    })
                    router.push('/')
                } else {
                    ElMessage({
                        message: '数据未填写完整，请检查',
                        type: 'error',
                    })
                }
            }).catch(err => {
                console.log(err);
                ElMessage({
                    message: '系统异常错误',
                    type: 'error',
                })
            })
        } else {
            ElMessage({
                message: '表单验证失败',
                type: 'error',
            })
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
    height: 450px;
    width: 650px;
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

span {
    color: red;
}
</style>