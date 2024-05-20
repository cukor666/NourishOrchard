<template>
  <div class="container">
    <div class="box">
      <h1>滋养果园</h1>
      <el-form :model="user" :rules="rules" ref="userRef">
        <el-form-item prop="username">
          <el-input v-model="user.username" placeholder="请输入用户名" clearable></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input type="password" v-model="user.password" placeholder="请输入密码" clearable></el-input>
        </el-form-item>
        <el-form-item>
          <el-button class="login-btn" type="primary" @click="login">登录</el-button>
          <a class="register-btn" style="margin-left: auto" :href="register_url">注册</a>
        </el-form-item>
      </el-form>
    </div>

    <!--  来点花样  -->
    <div class="circle"></div>
    <div class="heart"></div>
    <div class="rectangle"></div>
  </div>
</template>


<script setup>

import {ref} from "vue";
import router from "@/router/index.js";
import request from "@/axios/request.js";
import {ElMessage} from "element-plus";
import { useLoginUserStore } from "@/stores/loginUser.js"
import {admin_url_pro} from "@/config/api.js";

const register_url = admin_url_pro + "/register"

const user = ref({
  username: "",
  password: ""
})

const { setUsername, setPromise } = useLoginUserStore()

const userRef = ref(null)

// 校验用户名格式
const validateUsername = (rule, value, callback) => {
  if (value === "") {
    callback(new Error("请输入用户名"));
  } else {
    // 正则表达式验证用户名格式
    const reg = /^CZKJ[0-9]{4,16}$/;
    if (!reg.test(value)) {
      callback(new Error("用户名格式不正确"));
    } else {
      callback()
    }
  }
}

// 校验密码格式
const validatePassword = (rule, value, callback) => {
  if (value === "") {
    callback(new Error("请输入密码"));
  } else {
    // 正则表达式验证密码格式
    const reg = /^[a-zA-Z0-9]{6,16}$/;
    if (!reg.test(value)) {
      callback(new Error("密码格式不正确"));
    } else {
      callback()
    }
  }
}

const rules = ref({
  username: [{validator: validateUsername, trigger: "blur"}],
  password: [{validator: validatePassword, trigger: "blur"}]
})

const login = async () => {
  console.log('登录')
  // 整体再校验一遍
  userRef.value.validate((valid) => {
        if (!valid) {
          console.log('表单校验失败')
          return false;
        }

        console.log('表单校验成功')
        // 登录逻辑
        request.post('/login', {
          username: user.value.username,
          password: user.value.password,
          promise: "user"
        }).then(res => {
          if (res.code === 200) {
            console.log(res.data)
            ElMessage.success({
              message: '登录成功',
              type: 'success'
            })
            // 保存用户信息
            localStorage.setItem('nourish-orchard-user-token', res.data)
            localStorage.setItem('nourish-orchard-user-name', user.value.username)
            // 保存用户信息到 store
            setUsername(user.value.username)
            setPromise("user")
            // 进入主页
            router.push('/')
          } else {
            ElMessage.error({
              message: '登录失败，请检查用户名或密码',
              type: 'error'
            })
          }
        }).catch(err => {
          console.log(err)
          ElMessage.error({
            message: '登录失败，请检查网络',
            type: 'error'
          })
        })
      }
  )
}

</script>

<style lang="scss" scoped>
@import "@/sass/login";

</style>