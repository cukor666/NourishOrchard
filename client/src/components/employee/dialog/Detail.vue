<template>
  <el-dialog :model-value="props.detailDialogV" title="员工详情" width="500" @close="closeDialog">
    <el-form label-width="100px">
      <el-form-item label="ID：">
        <el-input v-model.number="props.employee.id" disabled/>
      </el-form-item>
      <el-form-item label="账号：">
        <el-input v-model="props.employee.username" disabled/>
      </el-form-item>
      <el-form-item label="姓名：">
        <el-input v-model="props.employee.name" clearable placeholder="请输入姓名，例如：吴宣仪..."
                  @blur="validName(props.employee.name)"/>
      </el-form-item>
      <div v-if="errWord.name" class="error-word left-space">{{ errWord.name }}</div>
      <el-form-item label="联系电话：">
        <el-input v-model="props.employee.phone" clearable placeholder="请输入联系电话，例如：13745968672"
                  @blur="validPhone(props.employee.phone)"/>
      </el-form-item>
      <div v-if="errWord.phone" class="error-word left-space">{{ errWord.phone }}</div>
      <el-form-item label="职位：">
        <el-input v-model="props.employee.position" clearable placeholder="请输入职位，例如：仓库管理员"
                  @blur="validPosition(props.employee.position)"/>
      </el-form-item>
      <div v-if="errWord.position" class="error-word left-space">{{ errWord.position }}</div>
      <el-form-item label="薪资：">
        <el-input v-model.number="props.employee.salary" clearable placeholder="请输入职位，例如：仓库管理员"
                  @blur="validSalary(props.employee.salary)"/>
      </el-form-item>
      <div v-if="errWord.salary" class="error-word left-space">{{ errWord.salary }}</div>

    </el-form>
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="cancel">取消</el-button>
        <el-button type="primary" @click="update">
          更新
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>

import {reactive} from "vue";
import {useValid} from "@/hooks/common/useValid.js";
import {ElMessage, ElMessageBox} from "element-plus";
import request from "@/axios/request.js";

let props = defineProps({
  detailDialogV: Boolean,
  employee: {
    type: Object,
    default: {
      id: 0,
      username: '',
      name: '',
      phone: '',
      position: '',
      salary: 0
    }
  }
})

const emit = defineEmits(['closeDialog', 'updateEmployee'])

const closeDialog = () => {
  emit('closeDialog')
}

const errWord = reactive({
  id: '',
  account: '',
  name: '',
  phone: '',
  position: '',
  salary: ''
})

const {validName, validPosition, validPhone, validSalary} = useValid(errWord)

const cancel = () => {
  closeDialog()
}

const update = () => {
  console.log(props.employee)
  ElMessageBox.confirm('您确定要更改员工信息吗？', '更新员工', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
      request.put('/employee/update', {
        ...props.employee
      }).then(res => {
        if (res.code === 200) {
          ElMessage({message: '更新成功', type: 'success'})
          emit('updateEmployee')
        } else {
          console.log(res.msg)
          ElMessage({message: '参数错误', type: 'error'})
        }
      }).catch(err => {
        console.error(err)
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