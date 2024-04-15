<template>
  <el-dialog :model-value="props.addWarehouseDialogV" title="添加仓库" width="500" @close="closeDialog">
    <el-form label-width="100px">
      <el-form-item label="仓库地址：">
        <el-input v-model="warehouse.address" clearable placeholder="请输入仓库地址，如：北京市海淀区中科院物理所"/>
      </el-form-item>

      <el-form-item label="仓库容量：">
        <el-input-number v-model.number="warehouse.capacity" :min="0" :max="10000" placeholder="请输入仓库容量，单位：m³"/>
      </el-form-item>

      <el-form-item label="仓库状态：">
        <el-select v-model="warehouse.status" placeholder="请选择仓库状态">
          <el-option label="启用" value="0"></el-option>
          <el-option label="禁用" value="1"></el-option>
        </el-select>
      </el-form-item>

    </el-form>
    <template #footer>
      <el-button type="primary" @click="addWarehouse">添加</el-button>
      <el-button @click="cancel">取消</el-button>
    </template>
  </el-dialog>
</template>

<script setup>
import {ref} from "vue";
import request from "@/axios/request.js";
import {ElMessage} from "element-plus";
import {WarehouseAdd} from "@/api/warehouse/warehouse-api.js";

let props = defineProps({
  addWarehouseDialogV: Boolean
})

// 表单数据
const warehouse = ref({
  address: '',
  capacity: 0,
  status: 0
})

const emit = defineEmits(['closeAddDialog', 'cancelAddWarehouse'])

const closeDialog = () => {
  emit('closeAddDialog')
}

const addWarehouse = async () => {
  try {
    let res = await request.post(WarehouseAdd, {
      ...warehouse.value
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
  emit('cancelAddWarehouse')
}

</script>

<style scoped lang="scss">

</style>