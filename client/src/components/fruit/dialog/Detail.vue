<template>
  <el-dialog :model-value="props.detailDialogV" title="水果详情" width="500" @close="closeDialog">
    <el-form label-width="100px">
      <el-form-item label="ID：">
        <el-input v-model.number="props.fruit.fId" disabled/>
      </el-form-item>

      <el-form-item label="水果名称：">
        <el-input v-model="props.fruit.fName" clearable placeholder="请输入水果名称，例如：砂糖橘..."/>
      </el-form-item>

      <el-form-item label="含水量：">
        <el-input v-model.number="props.fruit.fWater" clearable placeholder="请输入含水量，例如：74..."/>
      </el-form-item>

      <el-form-item label="含糖量：">
        <el-input v-model.number="props.fruit.fSugar" clearable placeholder="请输入含糖量，例如：74..."/>
      </el-form-item>

      <el-form-item label="保质期：">
        <el-input v-model.number="props.fruit.fShelfLife" clearable placeholder="请输入保质期，例如：365..."/>
      </el-form-item>

      <!--供应商信息部分-->
      <el-form-item label="供应商id：">
        <el-input v-model.number="props.fruit.sId" clearable placeholder="请输入供应商id，例如：74..."/>
      </el-form-item>

      <el-form-item label="供应商名称：">
        <el-input v-model="props.fruit.sName" disabled/>
      </el-form-item>

      <el-form-item label="供应商地址：">
        <el-input v-model="props.fruit.sAddress" disabled/>
      </el-form-item>

      <el-form-item label="供应商联系人：">
        <el-input v-model="props.fruit.sContactPerson" disabled/>
      </el-form-item>

      <el-form-item label="供应商电话：">
        <el-input v-model="props.fruit.sPhone" disabled/>
      </el-form-item>

      <el-form-item label="供应商信誉：">
        <el-input v-model="props.fruit.sReputation" disabled/>
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
import {FruitUpdate} from "@/api/fruit/fruit-api.js";

let props = defineProps({
  detailDialogV: Boolean,
  fruit: {
    type: Object,
    default: {
      fId: 0,
      fName: '',
      fWater: 0,
      fSugar: 0,
      fShelfLife: 0,
      fOrigin: '',
      sId: 0,
      sName: '',
      sAddress: '',
      sContactPerson: '',
      sPhone: '',
      sReputation: 0
    }
  }
})

const emit = defineEmits(['closeDialog', 'updateFruit'])

const closeDialog = () => {
  emit('closeDialog')
}

const cancel = () => {
  closeDialog()
}

const update = () => {
  console.log(props.fruit)
  ElMessageBox.confirm('您确定要更改该水果信息吗？', '更新水果', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    request.put(FruitUpdate, {
      id: props.fruit.fId,
      name: props.fruit.fName,
      water: props.fruit.fWater,
      sugar: props.fruit.fSugar,
      shelfLife: props.fruit.fShelfLife,
      origin: props.fruit.fOrigin,
      supplierId: props.fruit.sId
    }).then(res => {
      if (res.code === 200) {
        ElMessage({message: '更新成功', type: 'success'})
        emit('updateFruit')
      } else {
        ElMessage({message: '更新失败'+res.msg, type: 'error'})
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

</style>