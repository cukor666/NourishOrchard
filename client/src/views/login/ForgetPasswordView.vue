<template>
<!--  todo-->
  <div class="container">
    <div class="card">
      <el-form label-width="100px">
        <el-form-item label="账号：">
          <el-input v-model="user.account" clearable placeholder="请输入账号"/>
        </el-form-item>
        <el-form-item label="权限：">
          <el-radio-group v-model="promiseRadio" @change="promiseChange">
            <el-radio-button label="用户" value="user"/>
            <el-radio-button label="管理员" value="admin"/>
            <el-radio-button label="员工" value="employee"/>
          </el-radio-group>
        </el-form-item>
        <el-form-item v-if="user.promise === 'user' || user.promise === 'employee'" label="电话：">
          <el-input v-model="user.phone" clearable placeholder="请输入联系电话"/>
        </el-form-item>
        <el-form-item v-else-if="user.promise === 'admin'" label="邮箱：">
          <el-input v-model="user.email" clearable placeholder="请输入邮箱"/>
        </el-form-item>
        <el-form-item label="验证码：">
          <el-input v-model="user.code" clearable placeholder="请输入验证码"/>
        </el-form-item>

        <el-form-item label="新密码：">
          <el-input v-model="user.password" type="password" clearable placeholder="请输入新密码"/>
        </el-form-item>
        <el-form-item label="再次确认：">
          <el-input v-model="user.password2" type="password" clearable placeholder="再次确认新密码"/>
        </el-form-item>
      </el-form>
      <div class="button-group">
        <el-button @click="routerBack">返回</el-button>
        <el-button @click="update">更改</el-button>
      </div>
    </div>
  </div>
</template>

<script setup>
import {reactive, ref, watch} from "vue";
import router from "@/router/index.js";

const user = reactive({
  account: '',
  promise: '',
  password: '',
  password2: '',
  phone: '',
  email: '',
  code: ''
})

const promiseRadio = ref('用户')

watch(promiseRadio, (v) => {
  if (promiseRadio.value === '用户') {
    user.promise = 'user'
  } else if (promiseRadio.value === '管理员') {
    user.promise = 'admin'
  } else if (promiseRadio.value === '员工') {
    user.promise = 'employee'
  } else {
    user.promise = 'xxx'
  }
})

const promiseChange = () => {
  console.log('user.promise = ', user.promise)
}

const routerBack = () => {
  router.back()
}

const update = () => {
  console.log('更改')
}

</script>

<style lang="scss" scoped>
.container {
  min-width: 100vw;
  min-height: 100vh;
  animation-name: changeBackground;
  animation-duration: 8s;
  animation-iteration-count: infinite;
  display: flex;
  justify-content: center;
  align-items: center;

  .card {
    width: 400px;
    border: 1px solid red;
    padding: 20px;
    border-radius: 30px;

    .el-input {
      width: 90%;
    }

    .button-group {
      display: flex;
      justify-content: space-evenly;
      align-items: center;

      .el-button {
        width: 100px;
        border-radius: 30px;
      }
    }
  }
}

@keyframes changeBackground {
  0% {
    background: #a8f6f6;
  }
  25% {
    background: #71eeee;
  }
  50% {
    background: #48eeee;
  }
  75% {
    background: #71eeee;
  }
  100% {
    background: #a8f6f6;
  }
}
</style>
