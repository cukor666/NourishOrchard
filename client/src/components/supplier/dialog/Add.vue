<template>
  <el-dialog :model-value="props.addSupplierDialogV" title="添加供应商" width="500" @close="closeDialog">
    <el-form label-width="100px">
      <el-form-item label="供应商名称：">
        <el-input v-model="supplier.name" clearable placeholder="请输入供应商名称，例如：xx公司"/>
      </el-form-item>

      <el-form-item label="供应商地址：">
        <el-input v-model="supplier.address" clearable placeholder="请输入供应商地址，例如：广西百色"/>
      </el-form-item>

      <el-form-item label="联系人名：">
        <el-input v-model="supplier.contactPerson" clearable placeholder="请输入供应商联系人名，例如：张帅"/>
      </el-form-item>

      <el-form-item label="联系人电话：">
        <el-input v-model="supplier.phone" clearable placeholder="请输入联系人电话，例如：159..."/>
      </el-form-item>

      <el-form-item label="供应商信誉：">
        <el-input v-model.number="supplier.reputation" clearable placeholder="请输入供应商信誉，例如：88"/>
      </el-form-item>
    </el-form>

    <template #footer>
      <el-button type="primary" @click="addSupplier">添加</el-button>
      <el-button @click="cancel">取消</el-button>
    </template>
  </el-dialog>
</template>

<script setup>
import request from "@/axios/request.js";
import {ElMessage} from "element-plus";
import {useSupplier} from "@/hooks/search/useSupplier.js";
import {SupplierAdd} from "@/api/supplier/sup-api.js";

let props = defineProps({
  addSupplierDialogV: Boolean
})

const { supplier } = useSupplier()

const emit = defineEmits(['closeAddDialog', 'cancelAddSupplier'])

const closeDialog = () => {
  emit('closeAddDialog')
}

const addSupplier = async () => {
  try {
    let res = await request.post(SupplierAdd, {
      ...supplier.value
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
  emit('cancelAddSupplier')
}

</script>

<style scoped lang="scss">

</style>