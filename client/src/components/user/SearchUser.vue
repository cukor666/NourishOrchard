<template>
  <el-dialog :model-value="props.searchDialogV" title="搜索用户信息" width="500" @close="closeDialog">
    <el-form label-width="100px">
      <el-form-item label="ID：">
        <el-input v-model.number="user.id" type="number" clearable placeholder="请输入id" @blur="validId(user.id)"/>
      </el-form-item>
      <div v-if="errWord.id" class="error-word left-space">{{ errWord.id }}</div>

      <el-form-item label="账号：">
        <el-input v-model="user.username" clearable placeholder="请输入账号，例如：CZKJ10..."
                  @blur="validAccount(user.username)"/>
      </el-form-item>
      <div v-if="errWord.account" class="error-word left-space">{{ errWord.account }}</div>

      <el-form-item label="姓名：">
        <el-input v-model="user.name" clearable placeholder="请输入姓名，例如：吴宣仪..." @blur="validName(user.name)"/>
      </el-form-item>
      <div v-if="errWord.name" class="error-word left-space">{{ errWord.name }}</div>

      <el-form-item label="性别：">
        <el-select v-model="user.gender" clearable placeholder="请选择性别，例如：男" @blur="validGender(user.gender)">
          <el-option v-for="item in genderOption" :key="item.label" :label="item.label" :value="item.value"/>
        </el-select>
      </el-form-item>
      <div v-if="errWord.gender" class="error-word left-space">{{ errWord.gender }}</div>

      <el-form-item label="联系电话：">
        <el-input v-model="user.phone" clearable placeholder="请输入联系电话，例如：13745968672" @blur="validPhone(user.phone)"/>
      </el-form-item>
      <div v-if="errWord.phone" class="error-word left-space">{{ errWord.phone }}</div>

      <el-form-item label="联系地址：">
        <el-input v-model="user.address" clearable placeholder="请输入联系地址，例如：广东省深圳市福田区..." @blur="validAddress(user.address)"/>
      </el-form-item>
      <div v-if="errWord.address" class="error-word left-space">{{ errWord.address }}</div>

    </el-form>
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="cancel">取消</el-button>
        <el-button type="primary" @click="search">
          搜索
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>
import {reactive, ref} from 'vue'
import {useUser} from "@/hooks/search/useUser.js";
import {useValid} from "@/hooks/common/useValid.js";
import {ElMessage} from "element-plus";
import request from "@/axios/request.js";

let props = defineProps({
  searchDialogV: Boolean
})

const emit = defineEmits(['closeDialog'])

const closeDialog = () => {
  emit('closeDialog')
}

const errWord = reactive({
  id: '',
  account: '',
  name: '',
  gender: '',
  phone: '',
  address: ''
})

const {user} = useUser()

const genderOption = ref([
  {label: '男', value: '男'},
  {label: '女', value: '女'},
  {label: '全部', value: ''}
])

const cancel = () => {
  closeDialog()
}


const {validId, validAccount, validName, validGender, validPhone, validAddress} = useValid(errWord)

const valid = () => {
  validId(user.value.name)
  validAccount(user.value.username)
  validName(user.value.name)
  validGender(user.value.gender)
  validPhone(user.value.phone)
  validAddress(user.value.address)
}

const search = () => {
  valid()
  let filterErrWord = Object.keys(errWord).filter(e => {
    return e !== ''
  })
  if (filterErrWord.length !== 0) {
    ElMessage({
      message: '表单校验失败',
      type: 'error'
    })
    return
  }
  request.get('/user/list', {
    params: {
      pageSize: 3,
      pageNum: 2,
      ...user.value
    }
  }).then(res => {
    if (res.code === 200) {
      console.log(res.data)
    } else {
      ElMessage({
        message: '参数错误',
        type: 'error'
      })
    }
  }).catch(err => {
    console.log(err)
    ElMessage({
      message: '服务端错误',
      type: 'error'
    })
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