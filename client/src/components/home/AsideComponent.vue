<template>
  <div class="container">
    <div class="logo">
      <div class="img" @click="changeAside"></div>
      <div class="title" v-if="showItemContent">滋养果园</div>
    </div>
    <el-menu class="menu" :default-active="asideStore.defaultActive" :collapse="isCollapse" style="border: 0;">
      <el-menu-item index="Home" @click="goHome">
        <template #title>
          <el-icon>
            <House/>
          </el-icon>
          <span v-if="showItemContent">首页</span>
        </template>
        <el-icon v-if="!showItemContent">
          <House/>
        </el-icon>
      </el-menu-item>
      <el-sub-menu index="User" v-if="promise === 'admin'">
        <template #title>
          <el-icon>
            <User/>
          </el-icon>
          <span v-if="showItemContent">人员管理</span>
        </template>
        <el-menu-item index="UserList" @click="gotoUserList">
          <template #title>
            <span>用户列表</span>
          </template>
        </el-menu-item>
        <el-menu-item index="AdminList" @click="gotoAdminList">
          <template #title>
            <span>管理员列表</span>
          </template>
        </el-menu-item>
      </el-sub-menu>
      <el-sub-menu index="Fruit">
        <template #title>
          <el-icon>
            <Grape/>
          </el-icon>
          <span v-if="showItemContent">水果管理</span>
        </template>
        <el-menu-item index="FruitList" @click="gotoFruitList">
          <template #title>
            <span>水果列表</span>
          </template>
        </el-menu-item>
      </el-sub-menu>
    </el-menu>
  </div>
</template>

<script setup>
import {computed, ref} from "vue";
import {useAsideStore} from '@/stores/aside'
import router from "@/router";
import {useHome} from "@/hooks/aside/useHome.js";
import {useUser} from "@/hooks/aside/useUser.js";
import {useAdmin} from "@/hooks/aside/useAdmin.js";
import {useFruit} from "@/hooks/aside/useFruit.js";

const {goHome} = useHome()
const {gotoUserList} = useUser()
const {gotoAdminList} =useAdmin()
const {gotoFruitList} = useFruit()

const asideStore = useAsideStore()

const promise = computed(() => {
  return sessionStorage.getItem('nourish-promise') || localStorage.getItem('nourish-promise')
})

const showItemContent = computed(() => {
  return asideStore.asideWidth === 200
})

const isCollapse = ref(false)

const changeAside = () => {
  asideStore.changeAsideWidth()
  isCollapse.value = !isCollapse.value
}

</script>

<style lang="scss" scoped>
@import "@/scss/home/aside/aside.scss";
</style>