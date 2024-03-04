<template>
    <div class="container">
        <div class="logo">
            <div class="img" @click="changeAside"></div>
            <div class="title" v-if="showItemContent">滋养果园</div>
        </div>
        <el-menu class="menu" :collapse="isCollapse" style="border: 0px;">
            <el-menu-item index="0" @click="router.push({ name: 'Home' })">
                <template #title>
                    <el-icon>
                        <House />
                    </el-icon>
                    <span v-if="showItemContent">首页</span>
                </template>
                <el-icon v-if="!showItemContent">
                    <House />
                </el-icon>
            </el-menu-item>
            <el-sub-menu index="1" v-if="promise === 'admin'">
                <template #title>
                    <el-icon>
                        <User />
                    </el-icon>
                    <span v-if="showItemContent">人员管理</span>
                </template>
                <el-menu-item index="1-1" @click="router.push({ name: 'UserList' })">
                    <template #title>
                        <span>用户列表</span>
                    </template>
                </el-menu-item>
                <el-menu-item index="1-2" @click="router.push({ name: 'AdminList' })">
                    <template #title>
                        <span>管理员列表</span>
                    </template>
                </el-menu-item>
            </el-sub-menu>
            <el-sub-menu index="2">
                <template #title>
                    <el-icon>
                        <Grape />
                    </el-icon>
                    <span v-if="showItemContent">水果管理</span>
                </template>
                <el-menu-item index="2-1" @click="router.push({ name: 'FruitList' })">
                    <template #title>
                        <span>水果列表</span>
                    </template>
                </el-menu-item>
            </el-sub-menu>
        </el-menu>
    </div>
</template>

<script setup>
import { computed, ref } from "vue";
import { useAsideStore } from '@/stores/aside'
import { useUserInfoStore } from '@/stores/userInfo'
import router from "@/router";
import { storeToRefs } from "pinia";

const asideStore = useAsideStore()
const userInfoStore = useUserInfoStore()
const { promise } = storeToRefs(userInfoStore)
const { setPromise } = userInfoStore

const showItemContent = computed(() => {
    return asideStore.asideWidth === 200
})

const isCollapse = ref(false)

const changeAside = () => {
    console.log('改变宽度');
    asideStore.changeAsideWidth()
    isCollapse.value = !isCollapse.value
}

</script>

<style lang="scss" scoped>
@import "@/scss/home/aside/aside.scss";
</style>