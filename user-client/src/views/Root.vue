<template>
  <div class="top">
    <div class="welcome">我的水果 你的爱~</div>
    <div class="link">
      <router-link to="/login" v-if=" !username ||username === ''">请登录</router-link>
      <span v-else>{{ username }}</span>
      <span v-if="username && username!== ''" class="logout" @click="logout">退出</span>
      <span>|</span>
      <a :href="admin_url_pro+'/register'">注册</a>
      <span>|</span>
      <router-link to="/order">我的订单</router-link>
      <span>|</span>
      <router-link to="/cart">购物车 ({{ total }})</router-link>
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
import {admin_url_dev, admin_url_pro} from "@/config/api.js"
import {useLoginUserStore} from "@/stores/loginUser.js";
import {storeToRefs} from "pinia";
import {useCartStore} from "@/stores/cart.js";
import {ElMessageBox} from "element-plus";

const {username} = storeToRefs(useLoginUserStore())
const {total} = storeToRefs(useCartStore())

const search = ref('')

onMounted(() => {
  if (username.value === '') {
    username.value = localStorage.getItem('nourish-orchard-user-name')
  }
})

const logout = () => {
  // 退出登录
  // 跳转到登录页面
  ElMessageBox.confirm('确认退出登录吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    localStorage.removeItem('nourish-orchard-user-name')
    localStorage.removeItem('nourish-orchard-user-token')
    window.location.href = '/login'
  }).catch(() => {
    // 取消操作

  })
}

</script>

<style scoped lang="scss">
@import "@/sass/root";

</style>