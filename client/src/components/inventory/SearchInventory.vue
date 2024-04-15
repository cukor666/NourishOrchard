<template>
  <el-dialog :model-value="props.searchDialogV" title="搜索库存信息" width="500" @close="closeDialog">
    <el-form label-width="100px">
      <el-form-item label="ID：">
        <el-input v-model.number="inventory.id" type="number" clearable placeholder="请输入id"/>
      </el-form-item>
      <el-form-item label="商品ID：">
        <el-input v-model.number="inventory.commodityId" clearable placeholder="请输入商品id"/>
      </el-form-item>

      <el-form-item label="库存数量：">
        <el-input v-model.number="inventory.quantity" type="number" clearable placeholder="请输入库存数量"/>
      </el-form-item>

      <el-form-item label="员工ID：">
        <el-input v-model.number="inventory.employeeId" type="number" clearable placeholder="请输入员工id"/>
      </el-form-item>

      <el-form-item label="仓库ID：">
        <el-input v-model.number="inventory.warehouseId" type="number" clearable placeholder="请输入仓库id"/>
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
import {useInventory} from "@/hooks/search/useInventory.js";

let props = defineProps({
  searchDialogV: Boolean
})

const emit = defineEmits(['closeDialog', 'search'])

const closeDialog = () => {
  emit('closeDialog')
}

const {inventory} = useInventory()


const cancel = () => {
  closeDialog()
}
const search = () => {
  emit('search', inventory.value)
}

</script>

<style scoped lang="scss"></style>