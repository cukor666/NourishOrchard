<template>
  <el-dialog :model-value="searchDialogV" title="搜索员工信息" width="500" @close="closeDialog">
    <el-form label-width="100px">
      <el-form-item label="ID：">
        <el-input v-model.number="employee.id" type="number" clearable placeholder="请输入id"/>
      </el-form-item>

      <el-form-item label="账号：">
        <el-input v-model="employee.username" clearable placeholder="请输入账号，例如：CZKJ10..."/>
      </el-form-item>

      <el-form-item label="姓名：">
        <el-input v-model="employee.name" clearable placeholder="请输入姓名，例如：吴宣仪..."/>
      </el-form-item>

      <el-form-item label="联系电话：">
        <el-input v-model="employee.phone" clearable placeholder="请输入联系电话，例如：13745968672"/>
      </el-form-item>

      <el-form-item label="职位：">
        <el-input v-model="employee.position" clearable placeholder="请输入职位，例如：仓库管理员"/>
      </el-form-item>

      <el-form-item label="薪资：">
        <el-input v-model.number="employee.salaryA" clearable placeholder="请输入最低薪资，例如：3000"/>
        -
        <el-input v-model.number="employee.salaryB" clearable placeholder="请输入最高薪资，例如：3000"/>
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
import {useEmployee} from "@/hooks/search/useEmployee.js";

let props = defineProps({
  searchDialogV: Boolean
})

const emit = defineEmits(['closeDialog', 'search'])

const closeDialog = () => {
  emit('closeDialog')
}

const {employee} = useEmployee()

const cancel = () => {
  closeDialog()
}

const search = () => {
  emit('search', employee.value)
}

</script>

<style scoped lang="scss">

</style>