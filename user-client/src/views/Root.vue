<template>
  <div class="top">
    <div class="welcome">我的水果 你的爱~</div>
    <div class="link">
      <router-link to="/login">请登录</router-link>
      <router-link to="/register">注册</router-link>
      <span>|</span>
      <router-link to="/order">我的订单</router-link>
      <span>|</span>
      <router-link to="/cart">购物车 ({{ count }})</router-link>
    </div>
  </div>
  <div class="bar">
    <img style="width: 100px; height: 100px;"
         src="https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEj-NxuRo4oORl0C-daZ9m4nvncYMQtM9yv5dA-Pciwi-_FoYCfq83ZrjKpO17ixaUFpnfUTkEXJErJW-6TfhRXPfwr1jBgtOtndAGqU0s4bjucYq991TWOZbPZPWwL7C0qcHSEs4AZSJHs/s1600/Sweet-Love-Logo-CRsm.jpg"
         alt="logo">
    <div class="tab">
      <router-link to="/home">首页</router-link>
      <router-link to="#">南方水果</router-link>
      <router-link to="#">北方水果</router-link>
    </div>
    <input class="search" placeholder="搜索..." v-model="search"  />
  </div>
  <div class="content">
    <router-view></router-view>
  </div>
</template>

<script setup>

import {ref} from "vue";

const count = ref(5)
const search = ref('')
</script>

<style scoped lang="scss">
.top {
  width: 70%;
  margin: 20px auto 20px auto;
  font-size: 14px;
  color: #ccc;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

a {
  text-decoration: none;
}

.link {
  & > a {
    color: #ccc;
    margin-right: 10px
  }

  & > a:hover {
    color: #fa8038;
  }

  & > span {
    margin-right: 10px;
  }
}

.bar {
  display: flex;
  align-items: center;
  justify-content: center;
}

.tab {
  display: flex;
  justify-content: space-between;
  margin-left: 50px;
  margin-right: 50px;

  & > a {
    color: #727070;
  }

  & > a:hover {
    color: #fa8038;
  }

  & > * {
    margin-left: 20px;
    margin-right: 20px;
  }
}

.search {
  width: 200px;
  height: 30px;
  border-radius: 20px;
  padding: 10px;
  border: 1px solid #ccc;
}

.search:focus {
  outline: none;
  border-color: #fa8038;
}

.content {
  width: 80%;
  margin: 20px auto 20px auto;
}

</style>