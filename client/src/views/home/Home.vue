<template>
  <div class="container">
    <el-carousel indicator-position="outside" style="width: 100%">
      <el-carousel-item v-for="(item, index) in swiper" :key="index">
        <img :src="item.img" alt="无法显示"/>
      </el-carousel-item>
    </el-carousel>
    <div class="main-content">
      <div class="left">left</div>
      <div class="right">right</div>
    </div>
    <div class="footer">
      footer
    </div>
  </div>
</template>

<script setup>

import {ref} from "vue";

const swiper = ref([
  {
    id: 1,
    img: "https://pic.quanjing.com/u0/p9/QJ6660181509.jpg?x-oss-process=style/794wsy"
  },
  {
    id: 2,
    img: "https://pic.quanjing.com/bi/0c/QJ6375588182.jpg?x-oss-process=style/794wsy"
  },
  {
    id: 3,
    img: "https://pic.quanjing.com/gg/us/QJ6520051711.jpg?x-oss-process=style/794wsy"
  }
])

</script>

<style lang="scss" scoped>
.container {
  width: 100%;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.main-content {
  width: 100%;
  height: 260px;
  background-color: #00ff91;
  margin-top: 20px;
  display: flex;

  .left {
    flex: 1;
    height: 100%;
    background-color: #3c9fe7;
  }

  .right {
    flex: 1;
    height: 100%;
    background-color: #ff9100;
  }
}

.footer {
  width: 100%;
  height: 180px;
  background-color: #eabfbf;
  display: flex;
  justify-content: center;
  margin-top: 20px;
}

img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

</style>