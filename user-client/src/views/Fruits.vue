<template>
  <div class="bgimg"></div>
  <div class="current-path">当前位置 > 首页 > 全部水果</div>
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

  <el-dialog :title="detailObj.name" v-model="dialogVisible" width="800px" :before-close="handleClose">
    <div class="detail-content">
      <div class="detail-img" :style="{ backgroundImage: `url(${detailObj.imgs[currentImgIndex]})` }">
        <div class="left" @click="leftImg">
          <el-icon>
            <Back/>
          </el-icon>
        </div>
        <div class="right" @click="rightImg">
          <el-icon>
            <Right/>
          </el-icon>
        </div>
        <div class="bottom">
          <div v-for="item in detailObj.imgs.length" :key="item"
               :style="item === currentImgIndex + 1? { backgroundColor: '#ffb168' } : {backgroundColor: '#fff' }"></div>
        </div>
      </div>
      <div class="detail-form">
        <h1 class="detail-name">{{ detailObj.name }}</h1>
        <div class="detail-price">
          <span>价格：</span>
          <span>￥{{ detailObj.price }}</span>
        </div>
        <div class="detail-quantity">
          <span>数量：</span>
          <el-input v-model.number="quantity" type="number" :min="0" :max="10"/>
          <span>单位：吨</span>

        </div>
        <div class="detail-desc">
          <span>描述：</span>
          <span>{{ detailObj.desc }}</span>
        </div>
        <div class="total-price">总价格：<span>￥{{ detailObj.price * quantity }}</span></div>
        <div class="add-cart" @click="addCart">加入购物车</div>
      </div>
    </div>
  </el-dialog>


</template>

<script setup>
import {useFruitBase} from "@/hooks/useFruitBase.js"
import {onMounted, ref, watch} from "vue";
import request from "@/axios/request.js";
import {FruitList} from "@/api/fruit-api.js";
import {ElMessage} from "element-plus";
import {useDetailDialog} from "@/hooks/useDetailDialog.js";
const cnt = ref(parseInt(sessionStorage.getItem('all-fruit-cnt') || 7));
const step = 7;

watch(cnt, (newVal) => {
  sessionStorage.setItem('all-fruit-cnt', newVal.toString());
})

const {fruitList} = useFruitBase()

onMounted(async () => {
  try {
    let res = await request.get(FruitList, {
      params: {
        flag: 0
      }
    })
    if (res.code === 200) {
      // fruitList.value = res.data.filter(item => item.region === '全部')
      fruitList.value = res.data.fruits
    }
  } catch (error) {
    console.log(error)
    ElMessage.error('获取全部水果数据失败')
  }
})

const showMore = () => {
  cnt.value += step;
}

const refresh = () => {
  cnt.value = 7;
  sessionStorage.removeItem('all-fruit-cnt')
}

const {dialogVisible, detailObj, currentImgIndex, quantity, addCart, handleClose, leftImg, rightImg, showDetail} = useDetailDialog()


</script>



<style scoped lang="scss">
@import "@/sass/fruit-base";

.bgimg {
  background-image: url('https://pic4.zhimg.com/v2-7ebb2a5b63c994fba1f85ae8b13beaff_b.jpg');
}

</style>