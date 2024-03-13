<template>
  <el-form class="employee-info" label-width="100">
    <el-form-item label="账号：">
      <el-input v-model="employee.username" disabled/>
    </el-form-item>

    <el-form-item label="姓名：">
      <el-input v-model="employee.name" @blur="validName(employee.name)"/>
    </el-form-item>
    <div v-if="errWord.name" class="error-word left-space">{{ errWord.name }}</div>

    <el-form-item label="联系电话：">
      <el-input v-model="employee.phone" @blur="validPhone(employee.phone)"/>
    </el-form-item>
    <div v-if="errWord.phone" class="error-word left-space">{{ errWord.phone }}</div>

    <el-form-item label="职位：">
      <el-input v-model="employee.position" @blur="validPosition(employee.position)"/>
    </el-form-item>
    <div v-if="errWord.position" class="error-word left-space">{{ errWord.position }}</div>

    <el-form-item label="薪资：">
      <el-input v-model.number="employee.salary" @blur="validSalary(employee.salary)"/>
    </el-form-item>
    <div v-if="errWord.salary" class="error-word left-space">{{ errWord.salary }}</div>
  </el-form>
</template>

<script setup>
import {useEmployeeInfoStore} from '@/stores/employeeInfo'
import {storeToRefs} from "pinia";
import {reactive} from "vue";
import {useValid} from "@/hooks/common/useValid";

const employeeInfoStores = useEmployeeInfoStore()
const {employee} = storeToRefs(employeeInfoStores)

const errWord = reactive({
  name: '',
  phone: '',
  position: '',
  salary: 0
})

const {validName, validPhone, validPosition, validSalary} = useValid(errWord)

</script>

<style scoped lang="scss">
@import "@/scss/common/ErrorWord.scss";

.left-space {
  margin-top: -15px;
  margin-left: 100px;
}
</style>