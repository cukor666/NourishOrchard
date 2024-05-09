<template>
  <div class="top">
    <div class="welcome">我的水果 你的爱~</div>
    <div class="link">
      <router-link to="/login">{{username === '' ? '请登录' : username}}</router-link>
      <a :href="admin_url_dev+'/register'">注册</a>
      <span>|</span>
      <router-link to="/order">我的订单</router-link>
      <span>|</span>
      <router-link to="/cart">购物车 ({{ count }})</router-link>
    </div>
  </div>
  <div class="bar">
    <img style="width: 100px; height: 100px;"
         src="@/assets/nourish_orchard.png"
         alt="logo">
    <div class="tab">
      <router-link to="/home">首页</router-link>
      <router-link to="/south">南方水果</router-link>
      <router-link to="/north">北方水果</router-link>
      <router-link to="/fruits">全部水果</router-link>
    </div>
    <input class="search" placeholder="搜索..." v-model="search"/>
  </div>
  <div class="content">
    <router-view></router-view>
  </div>

  <hr>

  <div class="footer">
    <div>
      Coyright © 2023 cukor.cn. All rights reserved. 桂ICP备19003596号-1
    </div>
    <div>
      版权所有 桂ICP备19003596号-1 京公网安备110108020201703
    </div>
  </div>
</template>

<script setup>

import {onMounted, ref} from "vue";
import {admin_url_dev} from "@/config/api.js"
import {useLoginUserStore} from "@/stores/loginUser.js";
import {storeToRefs} from "pinia";

const { username } = storeToRefs(useLoginUserStore())
const count = ref(0)
const search = ref('')

onMounted(() => {
  if (username.value === '') {
    username.value = localStorage.getItem('nourish-orchard-user-name')
  }
})

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

hr {
  margin-top: 50px;
  margin-bottom: 50px;
}

.footer {
  display: flex;
  flex-direction: column;
  align-items: center;
  color: #9f9c9c;

  div {
    margin-left: auto;
    margin-right: auto;
    margin-bottom: 10px;
  }
}

</style>