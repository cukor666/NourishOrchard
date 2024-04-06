<template>
  <!--  todo-->
  <div class="container">
    <div class="card">
      <el-form label-width="100px">
        <el-form-item label="账号：">
          <el-input v-model="user.account" clearable placeholder="请输入账号" @blur="validAccount(user.account)"/>
        </el-form-item>
        <el-form-item label="权限：">
          <el-radio-group v-model="promiseRadio" @change="promiseChange">
            <el-radio-button label="用户" value="user"/>
            <el-radio-button label="管理员" value="admin"/>
            <el-radio-button label="员工" value="employee"/>
          </el-radio-group>
        </el-form-item>
        <el-form-item v-if="user.promise === 'user' || user.promise === 'employee'" label="电话：">
          <el-input v-model="user.phone" clearable placeholder="请输入联系电话" @blur="validPhone(user.phone)"/>
        </el-form-item>
        <el-form-item v-else-if="user.promise === 'admin'" label="邮箱：">
          <el-input v-model="user.email" clearable placeholder="请输入邮箱" @blur="validEmail(user.email)"/>
        </el-form-item>
        <el-form-item label="验证码：">
          <el-input v-model="user.code" clearable placeholder="请输入验证码"/>
        </el-form-item>

        <el-form-item label="新密码：">
          <el-input v-model="user.password" type="password" clearable placeholder="请输入新密码"
                    @blur="validPassword(user.password)"/>
        </el-form-item>
        <el-form-item label="再次确认：">
          <el-input v-model="user.password2" type="password" clearable placeholder="再次确认新密码"
                    @blur="validPassword2(user.password, user.password2)"/>
        </el-form-item>
      </el-form>
      <div class="button-group">
        <el-button class="back-button" @click="routerBack">返回</el-button>
        <el-button type="primary" class="update-button" @click="update">更改</el-button>
      </div>
    </div>
  </div>
</template>

<script setup>
import {reactive, ref, watch} from "vue";
import router from "@/router/index.js";
import {useValid} from "@/hooks/common/useValid.js";
import {ElMessage} from "element-plus";
import request from "@/axios/request.js";
import {ForgetPwd} from "@/api/login/login-api.js";

const user = reactive({
  account: '',
  promise: 'user',
  password: '',
  password2: '',
  phone: '',
  email: '',
  code: ''
})

const promiseRadio = ref('用户')

watch(promiseRadio, (v) => {
  if (v === '用户') {
    user.promise = 'user'
  } else if (v === '管理员') {
    user.promise = 'admin'
  } else if (v === '员工') {
    user.promise = 'employee'
  } else {
    user.promise = 'xxx'
  }
})

const promiseChange = () => {
  console.log('user.promise = ', user.promise)
}

const errWord = reactive({
  account: '',
  password: '',
  password2: '',
  email: ''
})

const {validPassword, validPassword2, validEmail, validPhone, validAccount} = useValid(errWord)

const valid = () => {
  validAccount(user.account)
  validPassword(user.password)
  validPassword2(user.password, user.password2)
  switch (user.promise) {
    case "user":
    case "employee":
      validPhone(user.phone)
      break
    case "admin":
      validEmail(user.email)
      break
    default:
      console.log('权限错误');
      ElMessage({message: '权限错误', type: 'error'})
  }
}

const routerBack = () => {
  router.back()
}

const update = async () => {
  valid()
  for (let key in errWord) {
    if (errWord[key] !== '') {
      ElMessage({message:'数据校验不通过', type: 'error'})
      return
    }
  }
  try {
    let res = await request.put(ForgetPwd, {
      ...user,
      username: user.account
    })
    if (res.code === 200) {
      console.log(res.data)
      ElMessage({message: '更新成功', type: 'success'})
    } else {
      console.log(res.msg)
      ElMessage({message: '参数错误', type: 'error'})
    }
  }catch (e) {
    console.error(e)
    ElMessage({message: '服务器端错误', type: 'error'})
  }
}

</script>

<style lang="scss" scoped>
@import "@/scss/login/ForgetPasswordView.scss";

</style>
