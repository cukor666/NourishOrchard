<template>
  <el-dialog :model-value="props.detailDialogV" title="用户详情" width="500" @close="closeDialog">
    <el-form label-width="100px">
      <el-form-item label="ID：">
        <el-input v-model.number="props.user.id" disabled/>
      </el-form-item>
      <el-form-item label="账号：">
        <el-input v-model="props.user.username" disabled/>
      </el-form-item>
      <el-form-item label="姓名：">
        <el-input v-model="props.user.name" clearable placeholder="请输入姓名，例如：吴宣仪..."
                  @blur="validName(props.user.name)"/>
      </el-form-item>
      <div v-if="errWord.name" class="error-word left-space">{{ errWord.name }}</div>
      <el-form-item label="性别：">
        <el-select v-model="props.user.gender" clearable placeholder="请选择性别，例如：男"
                   @blur="validGender(props.user.gender)">
          <el-option v-for="item in genderOption" :key="item.label" :label="item.label" :value="item.value"/>
        </el-select>
      </el-form-item>
      <div v-if="errWord.gender" class="error-word left-space">{{ errWord.gender }}</div>
      <el-form-item label="联系电话：">
        <el-input v-model="props.user.phone" clearable placeholder="请输入联系电话，例如：13745968672"
                  @blur="validPhone(props.user.phone)"/>
      </el-form-item>
      <div v-if="errWord.phone" class="error-word left-space">{{ errWord.phone }}</div>
      <el-form-item label="联系地址：">
        <el-input v-model="props.user.address" clearable placeholder="请输入联系地址，例如：广东省深圳市福田区..."
                  @blur="validAddress(props.user.address)"/>
      </el-form-item>
      <div v-if="errWord.address" class="error-word left-space">{{ errWord.address }}</div>
      <el-form-item label="出生日期：">
        <el-input v-model="props.user.birthday" clearable/>
      </el-form-item>

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
import {reactive, ref} from "vue";
import {useValid} from "@/hooks/common/useValid.js";
import {ElMessageBox, ElMessage} from "element-plus";
import request from "@/axios/request.js";

let props = defineProps({
  detailDialogV: Boolean,
  user: {
    type: Object,
    default: {
      id: 0,
      username: "",
      name: "",
      gender: "",
      phone: "",
      address: "",
      birthday: ""
    }
  }
})

const emit = defineEmits(['closeDialog', 'updateUser'])

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

const genderOption = ref([
  {label: '男', value: '男'},
  {label: '女', value: '女'},
  {label: '全部', value: ''}
])

const cancel = () => {
  closeDialog()
}

const update = () => {
  console.log(props.user)
  // 弹出一个确认是否更新用户数据的对话框
  // 如果点击取消，结束函数
  // 如果点击确定，向服务器发送更新请求，完成更新操作
  ElMessageBox.confirm('您确定要更改该用户信息吗？', '更新用户', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    request.put('/user/update', {
      ...props.user
    }).then(res => {
      if (res.code === 200) {
        ElMessage({message: '更新成功', type: 'success'})
        emit('updateUser')
      } else {
        ElMessage({message: '更新失败', type: 'error'})
      }
    }).catch(err => {
      console.log(err)
      ElMessage({message: '服务器错误', type: 'error'})
    })
  }).catch(() => {
    ElMessage({message: '取消', type: 'info'})
  })
}

const {validName, validGender, validPhone, validAddress} = useValid(errWord)


</script>

<style scoped lang="scss">
@import "@/scss/common/ErrorWord.scss";

.left-space {
  margin-left: 100px;
  margin-top: -15px;
}
</style>