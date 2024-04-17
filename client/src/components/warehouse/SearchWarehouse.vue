<template>
  <el-dialog :model-value="props.searchDialogV" title="搜索仓库信息" width="500" @close="closeDialog">
    <el-form label-width="100px">
      <el-form-item label="ID：">
        <el-input v-model.number="warehouse.id" type="number" clearable placeholder="请输入id"/>
      </el-form-item>

      <el-form-item label="地址：">
        <el-input v-model="warehouse.address" clearable placeholder="请输入地址"/>
      </el-form-item>

      <el-form-item label="容量：">
        <el-input v-model.number="warehouse.capacity" type="number" clearable placeholder="请输入容量"/>
      </el-form-item>

      <el-form-item label="仓库状态：">
        <el-select v-model="warehouse.status" clearable placeholder="请选择仓库状态">
          <el-option label="启用" value="1"></el-option>
          <el-option label="禁用" value="0"></el-option>
        </el-select>
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
import {useWarehouse} from "@/hooks/search/useWarehouse.js";

let props = defineProps({
  searchDialogV: Boolean
})

const emit = defineEmits(['closeDialog', 'search'])

const closeDialog = () => {
  emit('closeDialog')
}

const {warehouse} = useWarehouse()


const cancel = () => {
  closeDialog()
}
const search = () => {
  emit('search', warehouse.value)
}

</script>

<style scoped lang="scss"></style>