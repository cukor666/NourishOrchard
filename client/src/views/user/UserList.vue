<template>
  <div class="container">
    用户列表
  </div>
</template>

<script setup>
import {onMounted} from "vue";
import router from "@/router";

const promise = sessionStorage.getItem('nourish-promise') || localStorage.getItem('nourish-promise')

onMounted(() => {
  console.log(promise)
  if (promise !== 'admin') {
    router.push({name: 'Root'})
  }
})

</script>

<style lang="scss" scoped></style>