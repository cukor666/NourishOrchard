<template>
  <el-dialog :model-value="props.detailDialogV" title="仓库详情" width="500" @close="closeDialog">
    <el-form label-width="100px">
      <el-form-item label="ID：">
        <el-input v-model.number="props.warehouse.id" disabled/>
      </el-form-item>
      <el-form-item label="地址：">
        <el-input v-model="props.warehouse.address" />
      </el-form-item>
      <el-form-item label="容量：">
        <el-input v-model.number="props.warehouse.capacity" />
      </el-form-item>
      <el-form-item label="状态：">
        <el-input v-model.number="props.warehouse.status" />
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
import {WarehouseUpdate} from "@/api/warehouse/warehouse-api.js";

let props = defineProps({
  detailDialogV: Boolean,
  warehouse: {
    type: Object,
    default: {
      id: 0,
      address: "",
      capacity: 0,
      status: 0
    }
  }
})

const emit = defineEmits(['closeDialog', 'updateWarehouse'])

const closeDialog = () => {
  emit('closeDialog')
}

const cancel = () => {
  closeDialog()
}

const update = () => {
  ElMessageBox.confirm('您确定要更改该仓库信息吗？', '更新仓库', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    request.put(WarehouseUpdate, {...props.warehouse}).then(res => {
      if (res.code === 200) {
        ElMessage({message: '更新成功', type: 'success'})
        emit('updateWarehouse')
      } else {
        ElMessage({message: '更新失败', type: 'error'})
      }
    }).catch(err => {
      log.error(err)
      ElMessage({message: '服务器错误', type: 'error'})
    })
  }).catch(() => {
    ElMessage({message: '取消', type: 'info'})
  })
}
</script>

<style scoped lang="scss">
</style>