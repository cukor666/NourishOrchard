<template>
    <div class="container">
        用户列表
    </div>
</template>

<script setup>
import { onMounted } from "vue";
import { useUserInfoStore } from '@/stores/userInfo'
import { storeToRefs } from "pinia";
import router from "@/router";

const userInfoStore = useUserInfoStore()
const { promise } = storeToRefs(userInfoStore)

onMounted(() => {
    if (promise.value != 'admin') {
        router.push({ name: 'Root' })
    }
})

</script>

<style lang="scss" scoped></style>