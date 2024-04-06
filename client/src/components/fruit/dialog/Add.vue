<template>
  <el-dialog :model-value="props.addFruitDialogV" title="添加水果" width="500" @close="closeDialog">
    <el-form label-width="100px">
      <el-form-item label="水果名称：">
        <el-input v-model="fruit.name" clearable placeholder="请输入水果名称，例如：砂糖橘"/>
      </el-form-item>

      <el-form-item label="含水量：">
        <el-input v-model.number="fruit.water" clearable placeholder="请输入含水量，例如：80"/>
      </el-form-item>

      <el-form-item label="含糖量：">
        <el-input v-model.number="fruit.sugar" clearable placeholder="请输入含糖量，例如：30"/>
      </el-form-item>

      <el-form-item label="保质期：">
        <el-input v-model.number="fruit.shelfLife" clearable placeholder="请输入保质期，例如：7"/>
      </el-form-item>

      <el-form-item label="原产地：">
        <el-input v-model="fruit.origin" clearable placeholder="请输入原产地，例如：广西百色"/>
      </el-form-item>

      <el-form-item label="供应商id：">
        <el-input v-model.number="fruit.supplierId" clearable placeholder="请输入供应商id，例如：2"/>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button type="primary" @click="addFruit">添加</el-button>
      <el-button @click="cancel">取消</el-button>
    </template>
  </el-dialog>
</template>

<script setup>
import {ref} from "vue";
import request from "@/axios/request.js";
import {FruitAdd} from "@/api/fruit/fruit-api.js";
import {ElMessage} from "element-plus";

let props = defineProps({
  addFruitDialogV: Boolean
})

const fruit = ref({
  name: '',
  water: 0,
  sugar: 0,
  shelfLife: 0,
  origin: '',
  supplierId: 0
})

const emit = defineEmits(['closeAddDialog', 'cancelAddFruit'])

const closeDialog = () => {
  emit('closeAddDialog')
}

const addFruit = async () => {
  try {
    let res = await request.post(FruitAdd, {
      ...fruit.value
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
  emit('cancelAddFruit')
}

</script>

<style scoped lang="scss">

</style>