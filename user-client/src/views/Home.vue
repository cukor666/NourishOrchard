<template>
  <!-- 轮播图 -->
  <el-carousel height="300px">
    <el-carousel-item v-for="item in swiper" :key="item.id">
      <img :src="item.image" alt="">
    </el-carousel-item>
  </el-carousel>
  <!-- 推荐商品 -->
  <div class="recommend" v-for="item in recommend" :key="item.id">
    <!-- 左图右文    -->
    <img v-if="item.id % 2 === 1" :src="item.bgImg" alt=""/>
    <div v-if="item.id % 2 === 1" class="info">
      <div class="num" style="text-align: right">0{{ item.id }}</div>
      <div class="name" style="text-align: right">{{ item.name }}</div>
      <div class="desc">{{ item.desc }}</div>
      <div class="detail" style="margin-left: auto" @click="recommendDetail(item)">查看详情</div>
    </div>

    <!-- 右图左文    -->
    <div v-if="item.id % 2 === 0" class="info">
      <div class="num" style="text-align: left">0{{ item.id }}</div>
      <div class="name" style="text-align: left">{{ item.name }}</div>
      <div class="desc">{{ item.desc }}</div>
      <div class="detail" @click="recommendDetail(item)">查看详情</div>
    </div>
    <img v-if="item.id % 2 === 0" :src="item.bgImg" alt=""/>
  </div>





</template>

<script setup>

import {ref} from "vue";

const swiper = ref([
  {
    id: 1,
    image: "https://picsum.photos/id/1018/1000/300",
  },
  {
    id: 2,
    image: "https://picsum.photos/id/1015/1000/300",
  },
  {
    id: 3,
    image: "https://picsum.photos/id/1019/1000/300",
  }
])

const recommend = ref([
  {
    id: 1,
    name: "商品1",
    bgImg: "https://picsum.photos/id/1018/1000/300",
    desc: "商品描述1，写点什么吧，写点什么吧，写点什么吧，写点什么吧，写点什么吧，写点什么吧，写点什么吧，写点什么吧，写点什么吧，写点什么吧，写点什么吧！"
  },
  {
    id: 2,
    name: "商品2",
    bgImg: "https://picsum.photos/id/1018/1000/300",
    desc: "商品描述2，写点什么吧，写点什么吧！"
  },
  {
    id: 3,
    name: "商品3",
    bgImg: "https://picsum.photos/id/1018/1000/300",
    desc: "商品描述3，写点什么吧，写点什么吧！"
  },
  {
    id: 4,
    name: "商品4",
    bgImg: "https://picsum.photos/id/1018/1000/300",
    desc: "商品描述4，写点什么吧，写点什么吧！"
  }
])

const recommendDetail = (item) => {
  console.log(item)
}

</script>

<style scoped lang="scss">
.el-carousel {
  .el-carousel__item {
    display: flex;
    justify-content: center;
    align-items: center;
  }
}

.recommend {
  display: flex;
  margin-top: 40px;

  img {
    flex: 1;
    width: 50%;
    height: 350px;
    object-fit: cover;
  }

  .info {
    flex: 1;

    .num {
      font-size: 60px;
      text-align: right;
      font-weight: bold;
      position: relative;
      color: #bbb4b4;
      border-bottom: 1px solid #ccc;
    }

    .name {
      font-size: 24px;
      font-weight: bold;
      margin-top: 20px;
      margin-bottom: 20px;
    }

    .desc {
      width: 95%;
      margin-left: auto;
      margin-right: auto;
    }

    .detail {
      width: 100px;
      height: 30px;
      background-color: #ff7012;
      display: flex;
      justify-content: center;
      align-items: center;
      border-radius: 5px;
      margin-top: 50px;
      color: #fff;
      cursor: pointer;
    }
  }
}





</style>