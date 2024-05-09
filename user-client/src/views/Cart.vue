<template>
  <div class="container">
    <h1>total: {{total}}</h1>
    <ul>
      <li v-for="(item, index) in cartList" :key="index">
        <h2>{{ item.fruit.name }}</h2>
        <p>{{ item.quantity }}</p>
      </li>
    </ul>
  </div>
</template>

<script setup>

import {useCartStore} from "@/stores/cart.js";
import {storeToRefs} from "pinia";
import {onMounted} from "vue";

const { cartList, total } = storeToRefs(useCartStore())


onMounted(() => {
  console.log(cartList.value)
})

</script>

<style scoped lang="scss">

.container {
  width: 100vw;
  min-height: 100vh;
}

</style>