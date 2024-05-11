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
          <button class="remove-btn" @click="removeCart(index)">移出购物车</button>
        </div>
      </div>
    </div>

    <div class="bottom">
      <div class="total"><span>总计:</span><span style="color: red">￥{{ orderTotal }}</span></div>
      <button class="submit-btn" :disabled="submitBtnDisabled" @click="submitOrder">提交订单</button>
    </div>
  </div>
</template>

<script setup>

import {useCartStore} from "@/stores/cart.js";
import {storeToRefs} from "pinia";
import {computed, onMounted, ref, watch} from "vue";
import {useOrderStore} from "@/stores/order.js";
import router from "@/router/index.js";

const {cartList} = storeToRefs(useCartStore())
const {orderList, money, selectedIndexes} = storeToRefs(useOrderStore())

const selectAll = ref(false)
const selectList = ref([])

watch(selectAll, (val) => {
  if (val) {
    for (let i = 0; i < selectList.value.length; i++) {
      selectList.value[i] = true
    }
  }
})

watch(selectList.value, (val) => {
  let count = 0
  for (let i = 0; i < val.length; i++) {
    if (selectList.value[i]) {
      count++
    }
  }
  selectAll.value = count === selectList.value.length
})

const submitBtnDisabled = computed(() => {
  let arr = selectList.value.filter(v => v === true)
  console.log(arr.length === 0)
  return arr.length === 0
})

onMounted(() => {
  for (let i = 0; i < cartList.value.length; i++) {
    selectList.value.push(false)
  }
})

const orderTotal = computed(() => {
  let sum = 0
  for (let i = 0; i < selectList.value.length; i++) {
    if (selectList.value[i]) {
      sum += cartList.value[i].fruit.price * cartList.value[i].quantity
    }
  }
  return sum
})


const removeCart = (index) => {
  // 删除cartList.value中下标为index的元素
  cartList.value.splice(index, 1)
  // 更新localStorage中的cartList
  localStorage.setItem('cart', JSON.stringify(cartList.value))
  // 更新selectList
  selectList.value.splice(index, 1)
  // 更新selectAll
  selectAll.value = selectList.value.length === cartList.value.length
}

const submitOrder = () => {
  // 获取登录用户信息，如果用户没有登录，则提醒用户登录后再提交订单
  // 从localStorage中获取用户信息
  let username = localStorage.getItem('nourish-orchard-user-name')
  let token = localStorage.getItem('nourish-orchard-user-token')
  if (!username || !token) {
    alert('请先登录')
    return
  }

  money.value = orderTotal.value
  // 存储订单信息到sessionStorage
  for (let i = 0; i < selectList.value.length; i++) {
    if (selectList.value[i]) {
      let item = cartList.value[i]
      let order = {
        title: item.fruit.name + '-' + item.fruit.desc.slice(0, 10) + '...',
        commodityId: item.fruit.id,
        quantity: item.quantity,
        receiverName: '',
        receiverPhone: '',
        address: '',
        remark: ''
      }
      orderList.value.push({order})
      selectedIndexes.value.push(i)
    }
  }

  // 跳转到支付页面
  router.push('/pay')
}

</script>

<style scoped lang="scss">
@import "@/sass/cart";

</style>