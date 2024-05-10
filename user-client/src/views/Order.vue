<template>
  <div class="container">
    <h1>订单</h1>
    <el-table :data="orderList" style="width: 100%">
      <el-table-column prop="date" label="Date" width="180" />
      <el-table-column prop="name" label="Name" width="180" />
      <el-table-column prop="address" label="Address" />
    </el-table>
  </div>
</template>

<script setup>

const orderList = [
  // {
  //
  // }
]

</script>


<style scoped lang="scss">
.container {
  width: 95%;
  height: 100%;
  background-color: #fff;
  border-radius: 5px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
  padding: 20px;
}
</style>