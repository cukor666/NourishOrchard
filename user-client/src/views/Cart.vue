<template>
  <div class="container">
    <h1 class="title"><span class="text">全部商品</span></h1>
    <div class="table">
      <div class="table-title">
        <div class="checkbox">
          <input type="checkbox" v-model="selectAll">
          <div>全选</div>
        </div>
        <span class="fruit-name">商品名称</span>
        <span class="fruit-count">数量</span>
        <span class="fruit-price">价格</span>
        <span class="fruit-operation">操作</span>
      </div>

      <div class="cart-item" v-for="(item, index) in cartList" :key="index">
        <input type="checkbox" v-model="selectList[index]"/>
        <div class="img" :style="{backgroundImage: item.fruit.imgs[0]}"></div>
        <div class="fruit-name">{{ item.fruit.name }}</div>
        <div class="fruit-count">{{ item.quantity }}</div>
        <div class="fruit-price">{{ item.fruit.price }}</div>
        <div class="fruit-operation">
          <button class="remove-btn" @click="removeCart()">移出购物车</button>
        </div>
      </div>
    </div>


    <div class="bottom">
      <div class="total"><span>总计:</span><span>{{ orderTotal }}</span></div>
      <button class="submit-btn" @click="submitOrder">提交订单</button>
    </div>
  </div>
</template>

<script setup>

import {useCartStore} from "@/stores/cart.js";
import {storeToRefs} from "pinia";
import {computed, onMounted, ref} from "vue";

const {cartList, total} = storeToRefs(useCartStore())
const selectAll = ref(false)
const selectList = ref([false, false, false])

onMounted(() => {
  console.log(cartList.value)
})

const orderTotal = computed(() => {
  let sum = 0
  for (let item of cartList.value) {
    sum += item.fruit.price * item.quantity
  }
  return sum
})

const removeCart = () => {
  console.log('removeCart')
}

const submitOrder = () => {
  console.log('submitOrder')
}

</script>

<style scoped lang="scss">

.container {
  width: 95%;
  min-height: 100vh;
  background-color: #fff;
  margin-left: auto;
  margin-right: auto;
  border-radius: 5px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.3);
  padding: 20px;
}

.title {
  font-size: 24px;
  color: #f8a546;

  .text {
    margin-top: 10px;
    border-bottom: 6px solid #f8a546;
  }
}

.table {

  .table-title {
    display: flex;
    align-items: center;
    font-size: 14px;
    font-weight: bold;
    background-color: #F1F1F1;
    padding: 10px;
    position: relative;


    .checkbox {
      display: flex;
      align-items: center;
      margin-left: 10px;

      div {
        margin-left: 5px;
      }
    }

    .fruit-name {
      position: absolute;
      left: 120px;

    }

    .fruit-count {
      position: absolute;
      left: 400px;
    }

    .fruit-price {
      position: absolute;
      left: 500px;
    }

    .fruit-operation {
      position: absolute;
      left: 800px;
    }

    & > span:last-child {
      margin-right: 50px;
    }
  }
}

.cart-item {
  width: 100%;
  min-height: 100px;
  background-color: #FFF4E8;
  margin: 20px 0;
  box-shadow: 0 -5px 10px rgba(0, 0, 0, 0.3);
  display: flex;
  align-items: center;
  padding: 10px;
  position: relative;

  input[type="checkbox"] {
    position: absolute;
    left: 20px;
  }

  .img {
    position: absolute;
    left: 40px;
    width: 50px;
    height: 50px;
    background-color: #ccc;
    background-image: url("https://mediabluk.cnr.cn/record/img/cnr/CNRCDP/2023/0602/803fbea5de9d4b168df04d9f7cb7d28610.jpg");
    background-size: cover;
    background-position: center;
  }

  .fruit-name {
    position: absolute;
    left: 120px;
  }

  .fruit-count {
    position: absolute;
    left: 400px;
  }

  .fruit-price {
    position: absolute;
    left: 500px;
  }

  .fruit-operation {
    position: absolute;
    left: 800px;

    .remove-btn {
      background-color: #fd5e58;
      border: 1px solid #faa8a5;
      border-radius: 5px;
      color: #fff;
      padding: 5px 10px;
      cursor: pointer;

      &:hover {
        transform: scale(1.1);
      }
    }
  }
}


.bottom {
  width: 100%;
  min-height: 50px;
  background-color: #F1F1F1;
  display: flex;
  align-items: center;

  .total {
    margin-left: 20px;
    font-size: 16px;
    font-weight: bold;
    color: #000;
  }

  .submit-btn {
    margin-left: auto;
    margin-right: 240px;
    padding: 10px;
    background-color: #f8a546;
    border: none;
    border-radius: 10px;
    color: #fff;
    font-size: 14px;
    cursor: pointer;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.3);

    &:hover {
      transform: scale(1.1);
    }
  }
}
</style>