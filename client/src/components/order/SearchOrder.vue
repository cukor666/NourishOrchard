<template>
  <el-dialog :model-value="props.searchDialogV" title="搜索用户信息" width="500" @close="closeDialog">
    <el-form label-width="100px">
      <el-form-item label="ID：">
        <el-input v-model.number="order.id" type="number" clearable placeholder="请输入id"/>
      </el-form-item>

      <el-form-item label="订单标题：">
        <el-input v-model="order.title" clearable placeholder="请输入标题"/>
      </el-form-item>

      <el-form-item label="订单状态：">
        <el-select v-model="order.status" placeholder="请选择状态">
          <el-option label="全部" value=""></el-option>
          <el-option label="待付款" value="1"></el-option>
          <el-option label="待发货" value="2"></el-option>
          <el-option label="待收货" value="3"></el-option>
          <el-option label="已完成" value="4"></el-option>
          <el-option label="已取消" value="5"></el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="商品ID：">
        <el-input v-model.number="order.commodityId" type="number" clearable placeholder="请输入商品id"/>
      </el-form-item>

      <el-form-item label="用户ID：">
        <el-input v-model.number="order.buyerId" type="number" clearable placeholder="请输入用户id"/>
      </el-form-item>

      <el-form-item label="管理员ID：">
        <el-input v-model.number="order.adminId" type="number" clearable placeholder="请输入管理员id"/>
      </el-form-item>

      <el-form-item label="仓库ID：">
        <el-input v-model.number="order.warehouseId" type="number" clearable placeholder="请输入仓库id"/>
      </el-form-item>

      <el-form-item label="收货人名：">
        <el-input v-model="order.receiverName" clearable placeholder="请输入收货人名"/>
      </el-form-item>

      <el-form-item label="联系电话：">
        <el-input v-model="order.receiverPhone" clearable placeholder="请输入联系电话"/>
      </el-form-item>

      <el-form-item label="地址：">
        <el-input v-model="order.address" clearable placeholder="请输入地址"/>
      </el-form-item>

      <el-form-item label="备注：">
        <el-input v-model="order.remark" type="textarea" clearable placeholder="请输入备注"/>
      </el-form-item>

      <el-form-item label="创建时间：">
        <el-date-picker v-model="order.createdAt" type="datetimerange" range-separator="至" start-placeholder="开始日期" end-placeholder="结束日期" placeholder="选择日期" />
      </el-form-item>

      <el-form-item label="更新时间：">
        <el-date-picker v-model="order.updatedAt" type="datetimerange" range-separator="至" start-placeholder="开始日期" end-placeholder="结束日期" placeholder="选择日期" />
      </el-form-item>
    </el-form>

    <template #footer>
      <div class="dialog-footer">
        <el-button @click="cancel">取消</el-button>
        <el-button type="primary" @click="search">搜索</el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>
import {useOrder} from "@/hooks/search/useOrder.js";

let props = defineProps({
  searchDialogV: Boolean
})

const emit = defineEmits(['closeDialog', 'search'])

const closeDialog = () => {
  emit('closeDialog')
}

const {order} = useOrder()

const cancel = () => {
  closeDialog()
}
const search = () => {
  emit('search', order.value)
}

</script>

<style scoped lang="scss"></style>