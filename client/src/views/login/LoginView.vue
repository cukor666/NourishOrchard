<template>
  <div class="login-box">
    <div class="login-card">
      <h1>滋养果园</h1>
      <el-form>
        <el-form-item>
          <span class="label">账&nbsp;&nbsp;&nbsp;&nbsp;号：</span>
          <el-input v-model="user.account" placeholder="请输入账号" clearable @blur="validAccount(user.account)"/>
        </el-form-item>
        <div class="error-word left-space" v-if="errWord.account">{{ errWord.account }}</div>

        <el-form-item>
          <span class="label">密&nbsp;&nbsp;&nbsp;&nbsp;码：</span>
          <el-input v-model="user.password" type="password" placeholder="请输入密码" show-password
                    @blur="validPassword(user.password)"/>
        </el-form-item>
        <div class="error-word left-space" v-if="errWord.password">{{ errWord.password }}</div>
        <el-form-item>
          <router-link to="/forget-password" class="forget-password">忘记密码</router-link>
        </el-form-item>

      </el-form>
      <el-radio-group v-model="user.promise">
        <el-radio v-for="item in promiseList" :key="item.label" :label="item.label">
          {{ item.content }}
        </el-radio>
      </el-radio-group>
      <div class="button-group">
        <el-button type="primary" class="login-button" @click="login">登录</el-button>
        <el-button class="register-button" @click="gotoRegister">去注册</el-button>
      </div>

    </div>
  </div>
  <Vcode :show="isShow" @success="onSuccess" @close="isShow = false"/>
</template>

<script setup>
import {reactive, ref} from 'vue'
import {useValid} from '@/hooks/common/useValid'
import {usePromise} from '@/hooks/login/usePromise'
import {ElMessage} from 'element-plus'
import {useRouter} from 'vue-router'
import Vcode from "vue3-puzzle-vcode";
import request from '@/axios/request'

const router = useRouter()

const isClickLoginButton = ref(false)
const isShow = ref(false)

const user = reactive({
  account: '',
  password: '',
  promise: 'user'
})

const errWord = reactive({
  account: '',
  password: '',
  promise: ''
})

const {validAccount, validPassword, validPromise} = useValid(errWord)
const {promiseList} = usePromise()

const onSuccess = () => {
  console.log('验证成功');
  isShow.value = false

  // 清除数据空格
  user.account = user.account.trim()

  // 使用axios校验后端
  request.post('/login', {
    ...user,
    username: user.account
  }).then(res => {    // 后端校验通过则登录到主页
    if (res.code === 200) {
      let token = res.data
      localStorage.setItem('nourish-token', token)
      localStorage.setItem('nourish-account', user.account)
      localStorage.setItem('nourish-promise', user.promise)
      sessionStorage.setItem('nourish-promise', user.promise)
      ElMessage({
        message: '欢迎：' + user.account,
        type: 'success'
      })
      router.push({name: 'Root'})
    } else {
      ElMessage({
        message: '服务器响应失败',
        type: 'error'
      })
    }
  }).catch(error => {     // 后端校验不通过则不能登录
    console.log(error);
  })
}

const login = () => {
  if (isClickLoginButton.value) return;
  isClickLoginButton.value = true;
  validAccount(user.account);
  validPassword(user.password);
  validPromise(user.promise);
  for (const key in errWord) {
    if (Object.hasOwnProperty.call(errWord, key)) {
      if (errWord[key] !== '') {
        ElMessage({
          message: '用户信息不正确',
          type: 'error',
        })
        console.log('key: ', key);
        console.log('value: ', errWord[key]);
        setTimeout(() => {
          isClickLoginButton.value = false
        }, 2000);
        return;
      }
    }
  }
  // 弹出验证对话框
  isShow.value = true
  setTimeout(() => {
    isClickLoginButton.value = false
  }, 2000);
}

const gotoRegister = () => {
  router.push('/register')
}
</script>

<style lang="scss" scoped>
@import "@/scss/login/LoginView.scss";
@import "@/scss/common/ErrorWord.scss";

.left-space {
  margin-left: 70px;
}
</style>