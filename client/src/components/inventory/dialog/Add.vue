<template>
  <el-dialog :model-value="props.addInventoryDialogV" title="添加仓库" width="500" @close="closeDialog">
    <el-form label-width="100px">
      <el-form-item label="商品ID：">
        <el-input v-model.number="inventory.commodityId" clearable placeholder="请输入商品ID，如：12"/>
      </el-form-item>

      <el-form-item label="数量：">
        <el-input v-model.number="inventory.quantity" clearable placeholder="请输入数量，如：100"/>
      </el-form-item>

      <el-form-item label="操作员ID：">
        <el-input v-model.number="inventory.employeeId" clearable placeholder="请输入操作员ID，如：12"/>
      </el-form-item>

      <el-form-item label="仓库ID：">
        <el-input v-model.number="inventory.warehouseId" clearable placeholder="请输入仓库ID，如：12"/>
      </el-form-item>

    </el-form>
    <template #footer>
      <el-button type="primary" @click="addInventory">添加</el-button>
      <el-button @click="cancel">取消</el-button>
    </template>
  </el-dialog>
</template>

<script setup>
import {ref} from "vue";
import request from "@/axios/request.js";
import {ElMessage} from "element-plus";
import {InventoryAdd} from "@/api/inventory/inventory-api.js";

let props = defineProps({
  addInventoryDialogV: Boolean
})

// 表单数据
const inventory = ref({
  commodityId: 0,
  quantity: 0,
  employeeId: 0,
  warehouseId: 0
})

const emit = defineEmits(['closeAddDialog', 'cancelAddInventory'])

const closeDialog = () => {
  emit('closeAddDialog')
  emit('cancelAddInventory')
}

const addInventory = async () => {
  try {
    let res = await request.post(InventoryAdd, {
      ...inventory.value
    })
    if (res.code === 200) {
      ElMessage({message: '添加成功', type: 'success'})
      emit('closeAddDialog')
    } else {
      ElMessage({message: '添加失败', type: 'error'})
    }
  } catch (e) {
    console.log(e)
  }
}

const cancel = () => {
  emit('cancelAddInventory')
}

</script>

<style scoped lang="scss">

</style>