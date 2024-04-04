<template>
  <el-dialog :model-value="props.searchDialogV" title="搜索管理员信息" width="500" @close="closeDialog">
    <el-form label-width="100px">
      <el-form-item label="ID：">
        <el-input v-model.number="admin.id" type="number" clearable placeholder="请输入id"/>
      </el-form-item>

      <el-form-item label="账号：">
        <el-input v-model="admin.username" clearable placeholder="请输入账号，例如：CZKJ10..."/>
      </el-form-item>

      <el-form-item label="姓名：">
        <el-input v-model="admin.name" clearable placeholder="请输入姓名，例如：吴宣仪..."/>
      </el-form-item>

      <el-form-item label="邮箱：">
        <el-input v-model="admin.email" clearable placeholder="请输入邮箱，例如：saca@gmail.com"/>
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
import {useAdmin} from "@/hooks/search/useAdmin.js";

let props = defineProps({
  searchDialogV: Boolean
})

const emit = defineEmits(['closeDialog', 'search'])

const closeDialog = () => {
  emit('closeDialog')
}

const {admin} = useAdmin()

const cancel = () => {
  closeDialog()
}
const search = () => {
  emit('search', admin.value)
}

</script>

<style scoped lang="scss">

</style>