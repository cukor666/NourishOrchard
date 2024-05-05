<template>
  <div class="bgimg"></div>
  <div class="current-path">当前位置 > 首页 > 北方水果</div>
  <div class="content">
    <div class="fruit" v-for="fruit in fruitList.slice(0, cnt)" :key="fruit.id">
      <div class="fruit-img" :style="{ backgroundImage: `url(${fruit.imgs[0]})` }">
        <div class="detail" @click="showDetail(fruit)">详情</div>
      </div>
      <div class="fruit-name">{{ fruit.name }}</div>
      <div class="fruit-price">￥{{ fruit.price }}</div>
    </div>
  </div>
  <div class="btn-group">
    <div class="show-more" @click="showMore">展示更多</div>
    <div class="refresh" @click="refresh">收起</div>
  </div>

</template>


<script setup>
import {useFruitBase} from "@/hooks/useFruitBase.js"
import {onMounted, ref, watch} from "vue";
import request from "@/axios/request.js";
import {FruitList} from "@/api/fruit-api.js";
import {ElMessage} from "element-plus";
const cnt = ref(parseInt(sessionStorage.getItem('north-fruit-cnt') || 7));
const step = 7;

watch(cnt, (newVal) => {
  sessionStorage.setItem('north-fruit-cnt', newVal.toString());
})

const {fruitList} = useFruitBase()

onMounted(async () => {
  try {
    let res = await request.get(FruitList)
    if (res.code === 200) {
      // fruitList.value = res.data.filter(item => item.region === '南方水果')
      fruitList.value = res.data.fruits
    }
  } catch (error) {
    console.log(error)
    ElMessage.error('获取北方水果数据失败')
  }
})

const showMore = () => {
  cnt.value += step;
}

const refresh = () => {
  cnt.value = 7;
  sessionStorage.removeItem('north-fruit-cnt')
}

const showDetail = (fruit) => {
  console.log(fruit)
}

</script>



<style scoped lang="scss">
@import "@/sass/fruit-base";

.bgimg {
  background-image: url("https://m.guojiguoshu.com/sites/default/files/styles/large/public/field/image/cherries-598170_1280.jpg?itok=Yn9bCqQ_");
}

</style>