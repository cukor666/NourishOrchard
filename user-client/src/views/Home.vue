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
<!--      <div class="detail" style="margin-left: auto" @click="recommendDetail(item)">查看详情</div>-->
      <div class="detail" style="margin-left: auto" @click="gotoList">前往水果列表--&gt;</div>
    </div>

    <!-- 右图左文    -->
    <div v-if="item.id % 2 === 0" class="info">
      <div class="num" style="text-align: left">0{{ item.id }}</div>
      <div class="name" style="text-align: left">{{ item.name }}</div>
      <div class="desc">{{ item.desc }}</div>
<!--      <div class="detail" @click="recommendDetail(item)">查看详情</div>-->
      <div class="detail" style="margin-right: auto" @click="gotoList">前往水果列表--&gt;</div>
    </div>
    <img v-if="item.id % 2 === 0" :src="item.bgImg" alt=""/>
  </div>

</template>

<script setup>

import {ref} from "vue";
import router from "@/router/index.js";

const swiper = ref([
  {
    id: 1,
    image: "https://image.healthjd.com/da/jfs/t1/175013/33/14234/1201959/60c3664eEdda4aa17/703b12235f04461b.jpg",
  },
  {
    id: 2,
    image: "https://as.chdev.tw/web/article/d/9/4/d3b7ff83-99d1-4903-8463-79806b53f7191677059699.jpg",
  },
  {
    id: 3,
    image: "https://pic4.zhimg.com/v2-7ebb2a5b63c994fba1f85ae8b13beaff_b.jpg",
  }
])

const recommend = ref([
  {
    id: 1,
    name: "椰子",
    bgImg: "https://as.chdev.tw/web/article/d/9/4/d3b7ff83-99d1-4903-8463-79806b53f7191677059699.jpg",
    desc: "椰子，是一种多汁且营养丰富的水果，营养价值较高，具有很强的食疗作用。椰子的营养成分主要是维生素C、B1、B2、B3、B5、B6、B12、钙、铁、锌、磷、钾、钠、胆固醇、淀粉、蛋白质、脂肪、核黄素、胡萝卜素、维生素A、维生素E、叶酸、葡萄糖、酒精度等。椰子的营养价值较高，具有很强的食疗作用。"
  },
  {
    id: 2,
    name: "杨桃",
    bgImg: "https://p4.itc.cn/images01/20220316/e5cd4ada228c4d80a764b1b72131e621.jpeg",
    desc: "杨桃也是一种热带水果，原产于东南亚，在中国南部、台湾、海南等地也有种植。杨桃果实呈五角星形，皮薄光滑，果肉脆嫩多汁，味甜酸爽。杨桃富含维生素C、B族维生素、钾、镁等矿物质，以及膳食纤维。它具有生津止渴、解暑利尿、健胃消食、美容养颜等功效。杨桃还可以用于药用，有止咳化痰、降血压、镇痛等作用。杨桃的食用方法多样，可以鲜食、榨汁、做沙拉、糖水、蜜饯等。杨桃也是一种重要的烹饪食材，可以用来制作菜肴、糕点、饮品等。"
  },
  {
    id: 3,
    name: "木瓜",
    bgImg: "https://www.gz.gov.cn/img/0/210/210959/8178181.png",
    desc: "木瓜是一种热带水果，原产于美洲，现已广泛种植于世界各地的热带和亚热带地区。木瓜树为小乔木，可高达10米，树叶大而掌状，花朵黄色，果实长圆形，果肉厚实，味甜多汁。木瓜富含维生素C、B族维生素、钾、镁等矿物质，以及膳食纤维。它具有生津止渴、健脾胃、助消化、美容养颜等功效。木瓜还可以用于药用，有抗炎、止痛、抗肿瘤等作用。木瓜的食用方法多样，可以鲜食、榨汁、做沙拉、糖水、蜜饯等。木瓜也是一种重要的烹饪食材，可以用来制作菜肴、糕点、饮品等。"
  },
  {
    id: 4,
    name: "龙眼",
    bgImg: "https://www.gz.gov.cn/img/0/210/210935/8178185.jpg",
    desc: "龙眼，又名桂圆，原产于中国南方和越南北部，是热带地区重要的经济作物。龙眼树为常绿乔木，可高达20米，树叶椭圆形，有光泽，花朵白色，果实球形或椭圆形，外壳黄褐色，果肉白色或淡黄色，味甜多汁。龙眼富含维生素C、B族维生素、钾、镁等矿物质，以及膳食纤维。它具有益气补血、滋阴润肺、健脾养胃、抗衰老等功效。龙眼还可以用于药用，有止咳化痰、明目、安神等作用。龙眼的食用方法多样，可以鲜食、干制、加工成罐头、蜜饯等。龙眼也是一种重要的烹饪食材，可以用来制作菜肴、糕点、饮品等。"
  }
])

const gotoList = () => {
  router.push('/fruits')
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
      width: 130px;
      height: 30px;
      background-color: #ff7012;
      display: flex;
      justify-content: center;
      align-items: center;
      border-radius: 5px;
      margin-top: 50px;
      color: #fff;
      cursor: pointer;
      padding: 10px;
      font-size: 14px;
      &:hover {
        transition: all 0.5s ease;
        transform: scale(1.1);
      }
    }
  }
}





</style>