<template>
  <div class="container">
    <div class="register-box">
      <h1>欢迎加入滋养果园</h1>
      <el-form>
        <el-form-item>
          <span>姓&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;名：</span>
          <el-input v-model="user.name" placeholder="请输入名字" clearable @blur="validName(user.name)"/>
        </el-form-item>
        <div v-if="errWord.name" class="errword">{{ errWord.name }}</div>
        <el-form-item>
          <span>密&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;码：</span>
          <el-input v-model="user.password" type="password" placeholder="请输入密码" clearable
                    @blur="validPassword(user.password)"/>
        </el-form-item>
        <div v-if="errWord.password" class="errword">{{ errWord.password }}</div>
        <el-form-item>
          <span>确认密码：</span>
          <el-input v-model="user.password2" type="password" placeholder="请再次输入密码" clearable
                    @blur="validPassword2(user.password, user.password2)"/>
        </el-form-item>
        <div v-if="errWord.password2" class="errword">{{ errWord.password2 }}</div>

        <el-form-item style="display: inline-flex;">
          <span>性&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;别：</span>
          <el-switch v-model="gender" class="ml-2" inline-prompt
                     style="--el-switch-on-color: #2394f0; --el-switch-off-color: #FF6699; margin-left: 5px;"
                     active-text="男" inactive-text="女"/>
        </el-form-item>
        <div v-if="errWord.gender" class="errword">{{ errWord.gender }}</div>

        <el-form-item>
          <span>联系电话：</span>
          <el-input v-model="user.phone" placeholder="请输入联系电话" clearable @blur="validPhone(user.phone)"/>
        </el-form-item>
        <div v-if="errWord.phone" class="errword">{{ errWord.phone }}</div>

        <el-form-item>
          <span>联系地址：</span>
          <el-input v-model="user.address" placeholder="请输入联系地址" clearable @blur="validAddress(user.address)"/>
        </el-form-item>
        <div v-if="errWord.address" class="errword">{{ errWord.address }}</div>

        <el-form-item>
          <span style="flex: 1;">出生日期：</span>
          <el-date-picker style="flex: 5;" v-model="birthday" type="date" placeholder="请选择出生日期"
                          @blur="validBirthday(user.birthday)"/>
        </el-form-item>
        <div v-if="errWord.birthday" class="errword">{{ errWord.birthday }}</div>

      </el-form>
      <div class="button-group">
        <el-button class="register-button" @click="register">注册</el-button>
        <el-button class="back-button" @click="gotoLogin">去登录</el-button>
      </div>
    </div>
  </div>
  <el-dialog v-model="dialogVisible" title="注册成功" width="400" :before-close="handleClose">
    <span>欢迎加入滋养果园，您的账号为：{{ account }}</span>
    <template #footer>
      <div class="dialog-footer">
        <el-button type="primary" @click="copyText">复制到剪切板</el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>
import {computed, reactive, ref} from 'vue';
import {useRouter} from 'vue-router'
import request from '@/axios/request'
import {ElMessage} from 'element-plus';
import {useValid} from '@/hooks/common/useValid';

const router = useRouter()

const gender = ref(true)
const birthday = ref()

const user = reactive({
  name: '',
  password: '',
  password2: '',
  gender: computed(() => {
    return gender.value ? '男' : '女';
  }),
  phone: '',
  address: '',
  birthday: computed(() => {
    return birthday.value ? birthday.value.toISOString() : ""
  })
})

const errWord = reactive({
  name: '',
  password: '',
  password2: '',
  gender: '',
  phone: '',
  address: '',
  birthday: ''
})

const {
  validName, validPassword, validPassword2, validGender,
  validPhone, validAddress, validBirthday
} = useValid(errWord)

const valid = () => {
  validName(user.name)
  validPassword(user.password)
  validPassword2(user.password, user.password2)
  validGender(user.gender)
  validPhone(user.phone)
  validAddress(user.address)
  validBirthday(user.birthday)
}

const account = ref('')

const gotoLogin = () => {
  router.push({name: 'Login'})
}

const dialogVisible = ref(false)
const handleClose = () => {
  console.log('对话框关闭');
  dialogVisible.value = false
}

const copyText = () => {
  console.log('复制到剪切板');
  navigator.clipboard.writeText(account.value)
  ElMessage({message: '成功复制到剪切板', type: 'success'})
  dialogVisible.value = false
}

const register = () => {
  console.log(user);
  valid()
  let ok = true
  for (let key in errWord) {
    if (errWord.hasOwnProperty(key) && errWord[key] !== '') {
      ok = false
      break
    }
  }
  if (!ok) {
    ElMessage({message: '表单校验不通过', type: 'error'})
    return
  }

  dialogVisible.value = true

  request.post('/register', {...user}).then(res => {
    console.log(res);
    account.value = res.data.username
    dialogVisible.value = true
  }).catch(error => {
    console.log(error);
  })
}


</script>

<style lang="scss" scoped>
@import "@/scss/register/RegisterView.scss";

.errword {
  color: #f79aac;
  margin-left: 85px;
  margin-top: -10px;
}
</style>