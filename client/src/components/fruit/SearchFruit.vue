<template>
  <el-dialog :model-value="props.searchDialogV" title="搜索水果信息" width="500" @close="closeDialog">
    <el-form label-width="100px">
      <el-form-item label="ID：">
        <el-input v-model.number="fruit.id" type="number" clearable placeholder="请输入id"/>
      </el-form-item>

      <el-form-item label="水果名称：">
        <el-input v-model="fruit.name" clearable placeholder="请输入水果名称，例如：砂糖橘..."/>
      </el-form-item>

      <el-form-item label="含水量：">
        <el-input v-model.number="fruit.water" clearable placeholder="请输入含水量，例如：75"/>
      </el-form-item>

      <el-form-item label="含糖量：">
        <el-input v-model.number="fruit.sugar" clearable placeholder="请输入含糖量，例如：75"/>
      </el-form-item>

      <el-form-item label="保质期：">
        <el-input v-model.number="fruit.shelfLife" clearable placeholder="请输入保质期，例如：30"/>
      </el-form-item>

      <el-form-item label="原产地：">
        <el-input v-model="fruit.origin" clearable placeholder="请输入原产地，例如：广西"/>
      </el-form-item>

      <el-form-item label="供应商id：">
        <el-input v-model="fruit.supplierId" clearable placeholder="请输入供应商id，例如：6"/>
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
import {useFruit} from "@/hooks/search/useFruit.js";

let props = defineProps({
  searchDialogV: Boolean
})

const emit = defineEmits(['closeDialog', 'search'])

const closeDialog = () => {
  emit('closeDialog')
}

const {fruit} = useFruit()


const cancel = () => {
  closeDialog()
}
const search = () => {
  emit('search', fruit.value)
}

</script>

<style scoped lang="scss"></style>