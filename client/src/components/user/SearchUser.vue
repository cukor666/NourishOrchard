<template>
  <el-dialog :model-value="props.searchDialogV" title="搜索用户信息" width="500" @close="closeDialog">
    <el-form label-width="100px">
      <el-form-item label="ID：">
        <el-input v-model.number="user.id" type="number" clearable placeholder="请输入id"/>
      </el-form-item>

      <el-form-item label="账号：">
        <el-input v-model="user.username" clearable placeholder="请输入账号，例如：CZKJ10..."/>
      </el-form-item>

      <el-form-item label="姓名：">
        <el-input v-model="user.name" clearable placeholder="请输入姓名，例如：吴宣仪..."/>
      </el-form-item>

      <el-form-item label="性别：">
        <el-select v-model="user.gender" clearable placeholder="请选择性别，例如：男">
          <el-option v-for="item in genderOption" :key="item.label" :label="item.label" :value="item.value"/>
        </el-select>
      </el-form-item>

      <el-form-item label="联系电话：">
        <el-input v-model="user.phone" clearable placeholder="请输入联系电话，例如：13745968672"/>
      </el-form-item>

      <el-form-item label="联系地址：">
        <el-input v-model="user.address" clearable placeholder="请输入联系地址，例如：广东省深圳市福田区..."/>
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
import {ref} from 'vue'
import {useUser} from "@/hooks/search/useUser.js";

let props = defineProps({
  searchDialogV: Boolean
})

const emit = defineEmits(['closeDialog', 'search'])

const closeDialog = () => {
  emit('closeDialog')
}

const {user} = useUser()

const genderOption = ref([
  {label: '男', value: '男'},
  {label: '女', value: '女'},
  {label: '全部', value: ''}
])

const cancel = () => {
  closeDialog()
}
const search = () => {
  emit('search', user.value)
}

</script>

<style scoped lang="scss"></style>