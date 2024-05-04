<template>
  <div class="bgimg"></div>
  <div class="current-path">当前位置 > 首页 > 南方水果</div>
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

  <el-dialog :title="fruitDetail" v-model="dialogVisible" width="30%" :before-close="handleClose">
    <div class="detail-content">
      hahah
    </div>
  </el-dialog>


</template>

<script setup>

import {onMounted, ref, watch} from "vue";
import request from "@/axios/request.js";
import {FruitList} from "@/api/fruit-api.js";
import {ElMessage} from "element-plus";

const cnt = ref(parseInt(sessionStorage.getItem('south-fruit-cnt') || 7));
const step = 7;

watch(cnt, (newVal) => {
  sessionStorage.setItem('south-fruit-cnt', newVal.toString());
})

// 到时候这个从后端获取的数据
const fruitList = ref([{
  id: 0,
  name: '番石榴',
  price: 12.5,
  imgs: '',
  desc: '香甜可口，营养丰富，营养价值高。'
}])

onMounted(async () => {
  try {
    let res = await request.get(FruitList)
    if (res.code === 200) {
      // fruitList.value = res.data.filter(item => item.region === '南方水果')
      fruitList.value = res.data.fruits
    }

  } catch (error) {
    console.log(error)
    ElMessage.error('获取南方水果数据失败')
  }

})


const showMore = () => {
  cnt.value += step;
}

const refresh = () => {
  cnt.value = 7;
  sessionStorage.removeItem('south-fruit-cnt')
}

const dialogVisible = ref(false)
const fruitDetail = ref('')

const showDetail = (fruit) => {
  console.log(fruit)
  dialogVisible.value = true
  fruitDetail.value = fruit.name
}

const handleClose = () => {
  dialogVisible.value = false
}


</script>

<style scoped lang="scss">

.bgimg {
  width: 100%;
  height: 300px;
  background-image: url("https://cukor-resource-1318313222.cos.ap-guangzhou.myqcloud.com/img/番石榴.jpg");
  background-position: center;
  background-repeat: no-repeat;
  background-size: cover;
  border-radius: 15px;
  margin-bottom: 10px;
  box-shadow: 0 5px 10px rgba(0, 0, 0, 0.5);
  transition: transform 0.5s ease-in-out;
}

.bgimg:hover {
  transform: scale(1.01);
}

.current-path {
  font-size: 14px;
  color: #999;
  margin-bottom: 20px;
  margin-top: 20px;
}

.content {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 20px;

  .fruit {
    display: flex;
    flex-direction: column;
    align-items: center;

    .fruit-img {
      width: 150px;
      height: 200px;
      background-color: #ccc;
      background-position: center;
      background-repeat: no-repeat;
      background-size: cover;
      border-radius: 5px;
      overflow: hidden;
      position: relative;

      .detail {
        position: absolute;
        top: 50%;
        left: 100%;
        transform: translate(0, -50%);
        width: 150px;
        height: 30px;
        line-height: 30px;
        text-align: center;
        font-size: 12px;
        background-color: #fff;
        opacity: 0.8;
        cursor: pointer;
        transition: all 0.5s ease-in-out;
      }
    }

    .fruit-img:hover {
      transform: translateY(-5px) scale(1.1);

      .detail {
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
      }
    }

    .fruit-name {
      font-size: 16px;
      margin-top: 10px;
    }

    .fruit-price {
      font-size: 14px;
      font-weight: bold;
      color: #ff463f;
    }
  }

}

.btn-group {
  display: flex;
  justify-content: center;
  margin-top: 20px;
  margin-bottom: 20px;
  font-size: 0.8em;
}

.show-more,
.refresh {
  width: 100px;
  height: 20px;
  background-color: #f5a623;
  color: #fff;
  border: none;
  cursor: pointer;
  border-radius: 20px;
  text-align: center;
  line-height: 20px;
  margin-left: 5px;
  margin-right: 5px;
}

.show-more:hover,
.refresh:hover {
  transition: transform 0.3s ease-in-out;
  transform: scale(1.1);
}

</style>