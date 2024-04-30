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
          <a class="register-btn" style="margin-left: auto" href="http://localhost:5173/register">注册</a>
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

const user = ref({
  username: "",
  password: ""
})

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

const login = () => {
  console.log('登录')
  // 整体再校验一遍
  userRef.value.validate((valid) => {
    if (!valid) {
      console.log('表单校验失败')
      return false;
    }

    console.log('表单校验成功')
    // 登录逻辑

    // 进入主页
    router.push({name: "Root"})
  })
}

</script>

<style lang="scss" scoped>
.container {
  width: 100vw;
  min-height: 100vh;
  background: linear-gradient(to bottom right, #a0b6f6, #41a0e7);
  display: flex;
  justify-content: center;
  align-items: center;
}

.box {
  width: 500px;
  height: 400px;
  background: transparent;
  border-radius: 15px;
  box-shadow: 0 10px 10px rgba(0, 0, 0, 0.3);
  overflow: hidden;
  position: relative; /* 添加这一行，使z-index生效 */
  z-index: 2; /* 添加这一行，使box在circle1之上 */
  transition: box-shadow 1s ease-in-out, transform 1s ease-in-out;

  h1 {
    text-align: center;
    font-size: 50px;
    background: -webkit-linear-gradient(#eee, #dc8145 70%);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
  }
}

.box:hover {
  box-shadow: 0 20px 20px rgba(0, 0, 0, 0.3);
  transform: scale(1.01);
}

.circle {
  --r: 100px;
  width: var(--r);
  height: var(--r);
  background-color: #fff;
  border-radius: 50%;
  position: absolute;
  z-index: 1;
  box-shadow: 10px 10px 50px rgba(134, 36, 243, 0.7);
  animation: circleAnimate 2s ease-in-out infinite, circleAnimateFirst 2s ease-in-out 1;
}

@keyframes circleAnimateFirst {
  0% {
    transform: translate(240px, -100px) scale(0);
  }
  50% {
    transform: translate(240px, -100px) scale(1.5);
  }
  100% {
    transform: translate(240px, -100px) scale(1);
  }
}

@keyframes circleAnimate {
  0% {
    transform: translate(240px, -100px) scale(1);
    opacity: 1;
  }
  50% {
    transform: translate(240px, -100px) scale(1.2);
    opacity: 0.8;
  }
  100% {
    transform: translate(240px, -100px) scale(1);
    opacity: 1;
  }
}

.heart {
  --w: 50px;
  width: var(--w);
  height: var(--w);
  background-color: #fa2f2f;
  box-shadow: 0 0 50px rgb(252, 28, 28);
  position: relative;
  z-index: 3;
  animation: heartAnimation 2s ease-in-out infinite, heartAnimationFirst 2s ease-in-out 1;
}

.heart::before {
  --w: 50px;
  width: var(--w);
  height: var(--w);
  background-color: #fa2f2f;
  box-shadow: 0 0 10px rgb(252, 28, 28);
  position: absolute;
  content: "";
  border-radius: 50%;
  transform: translateX(-50%);
}

.heart::after {
  --w: 50px;
  width: var(--w);
  height: var(--w);
  background-color: #fa2f2f;
  box-shadow: 0 0 10px rgb(252, 28, 28);
  position: absolute;
  content: "";
  border-radius: 50%;
  transform: translateY(-50%);
}

@keyframes heartAnimationFirst {
  0% {
    transform: translate(500px, 160px) rotate(24deg) scale(1);
  }
  100% {
    transform: translate(-500px, -160px) rotate(24deg) scale(1);
  }
}

@keyframes heartAnimation {
  0% {
    transform: translate(-500px, -160px) rotate(24deg) scale(1);
  }
  30% {
    transform: translate(-500px, -160px) rotate(24deg) scale(1.2);
  }
  50% {
    transform: translate(-500px, -160px) rotate(24deg) scale(1.2);
  }
  70% {
    transform: translate(-500px, -160px) rotate(24deg) scale(1.2);
  }
  100% {
    transform: translate(-500px, -160px) rotate(24deg) scale(1);
  }
}

.rectangle {
  width: 100px;
  height: 100px;
  background-color: #fff;
  position: absolute;
  transform: translate(90px, 180px) rotate(45deg) skew(10deg, 20deg);
  box-shadow: 0 0 50px rgba(255, 211, 108, 0.7);
  animation: rectangleAnimation 2s ease-in-out infinite, rectangleAnimationFirst 2s ease-in-out 1;
}

@keyframes rectangleAnimationFirst {
  0% {
    transform: translate(90px, 400px) rotate(24deg) scale(1);
  }
  100% {
    transform: translate(90px, 180px) rotate(24deg) scale(1);
  }
}

@keyframes rectangleAnimation {
  0% {
    transform: translate(90px, 180px) rotate(45deg) skew(10deg, 20deg) rotate(0);
  }
  100% {
    transform: translate(90px, 180px) rotate(45deg) skew(10deg, 20deg) rotate(360deg);
  }
}

.el-form {
  width: 40%;
  margin-left: auto;
  margin-right: auto;
  margin-top: 60px;
}

.login-btn {
  background-color: #41a0e7;
  border: none;
  border-radius: 15px;
  transition: box-shadow 1s ease-in-out;
}

.login-btn:hover {
  background-color: #018eff;
  border: none;
  box-shadow: 0 0 10px rgba(3, 158, 255, 0.8), 0 0 20px rgba(3, 158, 255, 0.89), 0 0 30px rgb(3, 158, 255);
}

.register-btn {
  width: 60px;
  height: 30px;
  display: flex;
  justify-content: center;
  align-items: center;
  text-decoration: none;
  border-radius: 15px;
  background-color: #4dffa0;
  border: none;
  color: #fff;
  transition: box-shadow 1s ease-in-out;
}

.register-btn:hover {
  background-color: #4dffa0;
  border: none;
  color: #f1ed30;
  box-shadow: 0 0 10px rgba(62, 255, 156, 0.8), 0 0 20px rgba(76, 255, 165, 0.89), 0 0 30px rgb(3, 255, 142);
}

</style>