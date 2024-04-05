<template>
  <el-dialog :model-value="props.detailDialogV" title="供应商详情" width="500" @close="closeDialog">
    <el-form label-width="100px">
      <el-form-item label="ID：">
        <el-input v-model.number="props.supplier.id" disabled/>
      </el-form-item>

      <el-form-item label="供应商名称：">
        <el-input v-model="props.supplier.name" clearable placeholder="请输入供应商名称，例如：xx公司..."/>
      </el-form-item>

      <el-form-item label="联系地址：">
        <el-input v-model="props.supplier.address" clearable placeholder="请输入联系地址，例如：广东省深圳市福田区..."/>
      </el-form-item>

      <el-form-item label="联系人名：">
        <el-input v-model="props.supplier.contactPerson" clearable placeholder="请输入联系人名，例如：王富贵"/>
      </el-form-item>

      <el-form-item label="联系电话：">
        <el-input v-model="props.supplier.phone" clearable placeholder="请输入联系电话，例如：13745968672"/>
      </el-form-item>

      <el-form-item label="信誉：">
        <el-input v-model.number="props.supplier.reputation" clearable placeholder="请输入供应商信誉，例如：80"/>
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
import {SupplierUpdate} from "@/api/supplier/sup-api.js";

let props = defineProps({
  detailDialogV: Boolean,
  supplier: {
    type: Object,
    default: {
      id: 0,
      name: "",
      address: "",
      contactPerson: "",
      phone: "",
      reputation: 0
    }
  }
})

const emit = defineEmits(['closeDialog', 'updateSupplier'])

const closeDialog = () => {
  emit('closeDialog')
}

const cancel = () => {
  closeDialog()
}

const update = () => {
  console.log(props.supplier)
  ElMessageBox.confirm('您确定要更改该供应商信息吗？', '更新供应商', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    request.put(SupplierUpdate, {
      ...props.supplier
    }).then(res => {
      if (res.code === 200) {
        ElMessage({message: '更新成功', type: 'success'})
        emit('updateSupplier')
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
@import "@/scss/common/ErrorWord.scss";

.left-space {
  margin-left: 100px;
  margin-top: -15px;
}
</style>