<template>
  <div class="container">
    <h1>订单</h1>
    <el-table :data="orderList" style="width: 100%">
      <el-table-column prop="createdAt" label="生成订单日期" width="180"/>
      <el-table-column prop="id" label="订单编号" width="120"/>
      <el-table-column prop="title" label="订单标题" width="220"/>
      <el-table-column prop="status" label="订单状态" width="80"/>
      <el-table-column prop="receiverName" label="收件人姓名" width="120"/>
      <el-table-column prop="receiverPhone" label="收件人电话" width="120"/>
      <el-table-column prop="address" label="收件地址" width="240"/>
    </el-table>
  </div>
</template>

<script setup>

import {onMounted, ref} from "vue";
import request from "@/axios/request.js";
import {OrderList} from "@/api/order-api.js";
import {ElMessage} from "element-plus";

const orderList = ref([])

onMounted(() => {
  request.get(OrderList, {
    params: {
      pageSize: 100,
      pageNum: 1
    }
  }).then(res => {
    orderList.value = res.data.orders
    // 处理时间格式
    orderList.value.forEach(item => {
      item.createdAt = new Date(item.createdAt).toLocaleString()
      switch (item.status) {
        case 1:
          item.status = '待付款'
          break
        case 2:
          item.status = '待发货'
          break
        case 3:
          item.status = '待收货'
          break
        case 4:
          item.status = '已完成'
          break
        case 5:
          item.status = '已取消'
          break
        default:
          item.status = '未知状态'
      }
    })
  }).catch(err => {
    console.log(err)
    ElMessage.error('获取订单列表失败')
  })
})

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