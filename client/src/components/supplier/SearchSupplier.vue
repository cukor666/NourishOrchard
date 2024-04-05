<template>
  <el-dialog :model-value="props.searchDialogV" title="搜索用户信息" width="500" @close="closeDialog">
    <el-form label-width="100px">
      <el-form-item label="ID：">
        <el-input v-model.number="supplier.id" type="number" clearable placeholder="请输入id"/>
      </el-form-item>

      <el-form-item label="供应商名称：">
        <el-input v-model="supplier.name" clearable placeholder="请输入姓名，例如：xx公司..."/>
      </el-form-item>

      <el-form-item label="供应商地址：">
        <el-input v-model="supplier.address" clearable placeholder="请输入供应商地址，例如：广东省深圳市福田区..."/>
      </el-form-item>

      <el-form-item label="联系人名：">
        <el-input v-model="supplier.contactPerson" clearable placeholder="请输入联系人名，例如：王富贵..."/>
      </el-form-item>

      <el-form-item label="联系电话：">
        <el-input v-model="supplier.phone" clearable placeholder="请输入联系电话，例如：137..."/>
      </el-form-item>

      <el-form-item label="信誉：">
        <el-input v-model="supplier.reputation" clearable placeholder="请输入供应商信誉，例如：80"/>
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
import {useSupplier} from "@/hooks/search/useSupplier.js";

let props = defineProps({
  searchDialogV: Boolean
})

const emit = defineEmits(['closeDialog', 'search'])

const closeDialog = () => {
  emit('closeDialog')
}

const {supplier} = useSupplier()

const cancel = () => {
  closeDialog()
}
const search = () => {
  emit('search', supplier.value)
}

</script>

<style scoped lang="scss"></style>