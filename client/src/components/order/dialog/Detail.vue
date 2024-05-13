<template>
  <el-dialog :model-value="props.detailDialogV" title="用户详情" width="500" @close="closeDialog">
    <el-form label-width="100px">
      <el-form-item label="ID：">
        <el-input v-model.number="props.order.id" disabled/>
      </el-form-item>

      <el-form-item label="标题：">
        <el-input v-model="props.order.title" />
      </el-form-item>

      <el-form-item label="状态：">
        <el-select v-model="props.order.status" >
          <el-option label="待付款" value="1"></el-option>
          <el-option label="待发货" value="2"></el-option>
          <el-option label="待收货" value="3"></el-option>
          <el-option label="已完成" value="4"></el-option>
          <el-option label="已取消" value="5"></el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="商品ID：">
        <el-input v-model.number="props.order.commodityId" />
      </el-form-item>

      <el-form-item label="数量：">
        <el-input v-model.number="props.order.quantity" />
      </el-form-item>

      <el-form-item label="买家ID：">
        <el-input v-model.number="props.order.buyerId" disabled/>
      </el-form-item>

      <el-form-item label="管理员ID：">
        <el-input v-model.number="props.order.adminId" />
      </el-form-item>

      <el-form-item label="仓库ID：">
        <el-input v-model.number="props.order.warehouseId" />
      </el-form-item>

      <el-form-item label="收件人：">
        <el-input v-model="props.order.receiverName" />
      </el-form-item>

      <el-form-item label="收件电话：">
        <el-input v-model="props.order.receiverPhone" />
      </el-form-item>

      <el-form-item label="地址：">
        <el-input v-model="props.order.address" />
      </el-form-item>

      <el-form-item label="备注：">
        <el-input v-model="props.order.remark" />
      </el-form-item>

      <el-form-item label="创建时间：">
        <el-input v-model="props.order.createdAt" disabled/>
      </el-form-item>

      <el-form-item label="更新时间：">
        <el-input v-model="props.order.updatedAt" disabled/>
      </el-form-item>

    </el-form>
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="cancel">取消</el-button>
        <el-button type="primary" @click="update">更新</el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>
import {ElMessageBox, ElMessage} from "element-plus";
import request from "@/axios/request.js";
import {OrderUpdate} from "@/api/order/order-api.js";


let props = defineProps({
  detailDialogV: Boolean,
  order: {
    type: Object,
    default: {
      id: 0,
      title: "",
      status: 0,
      commodityId: 0,
      quantity: 0,
      buyerId: 0,
      adminId: 0,
      warehouseId: 0,
      receiverName: "",
      receiverPhone: "",
      address: "",
      remark: "",
      createdAt: "",
      updatedAt: ""
    }
  }
})

const emit = defineEmits(['closeDialog', 'updateOrder'])

const closeDialog = () => {
  emit('closeDialog')
}


const cancel = () => {
  closeDialog()
}

const update = () => {
  console.log(props.order)
  ElMessageBox.confirm('您确定要更改该订单信息吗？', '更新订单', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    request.put(OrderUpdate, {
      ...props.order
    }).then(res => {
      if (res.code === 200) {
        ElMessage({message: '更新成功', type: 'success'})
        emit('updateOrder')
      } else {
        ElMessage({message: '更新失败', type: 'error'})
      }
    }).catch(err => {
      console.log(err)
      ElMessage({message: '服务器错误', type: 'error'})
    })
  }).catch(() => {
    ElMessage({message: '取消', type: 'info'})
  })
}

</script>

<style scoped lang="scss">

</style>