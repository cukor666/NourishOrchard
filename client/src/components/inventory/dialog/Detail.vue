<template>
  <el-dialog :model-value="props.detailDialogV" title="仓库详情" width="500" @close="closeDialog">
    <el-form label-width="100px">
      <el-form-item label="ID：">
        <el-input v-model.number="props.inventory.id" disabled/>
      </el-form-item>

      <el-form-item label="商品ID：">
        <el-input v-model.number="props.inventory.commodityId" />
      </el-form-item>

      <el-form-item label="数量：">
        <el-input v-model.number="props.inventory.quantity" />
      </el-form-item>

      <el-form-item label="员工ID：">
        <el-input v-model.number="props.inventory.employeeId" />
      </el-form-item>

      <el-form-item label="仓库ID：">
        <el-input v-model.number="props.inventory.warehouseId" />
      </el-form-item>

      <el-form-item label="创建时间：">
        <el-input v-model="props.inventory.createdAt" disabled/>
      </el-form-item>

      <el-form-item label="更新时间：">
        <el-input v-model="props.inventory.updatedAt" disabled/>
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
import {InventoryUpdate} from "@/api/inventory/inventory-api.js";

let props = defineProps({
  detailDialogV: Boolean,
  inventory: {
    type: Object,
    default: {
      id: 0,
      commodityId: 0,
      quantity: 0,
      employeeId: 0,
      warehouseId: 0,
      createdAt: '',
      updatedAt: ''
    }
  }
})

const emit = defineEmits(['closeDialog', 'updateInventory'])

const closeDialog = () => {
  emit('closeDialog')
}

const cancel = () => {
  closeDialog()
}

const update = () => {
  ElMessageBox.confirm('您确定要更改该库存信息吗？', '更新库存信息', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    request.put(InventoryUpdate, {...props.inventory}).then(res => {
      if (res.code === 200) {
        ElMessage({message: '更新成功', type: 'success'})
        emit('updateInventory')
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