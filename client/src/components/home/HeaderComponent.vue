2
<template>
  <div class="container">
    <div class="path__content">
      <span style="font-size: 14px; font-weight: bold;">后台管理系统</span>
    </div>
    <div class="setting-bar">
      <el-dropdown :hide-on-click="false">
        <div style="display: flex; align-items: center; justify-content: space-around">
          <div style="color: #f5f1f1">Hi~ {{ promise }} !</div>
          <div style="margin-left: 10px; margin-right: 10px" class="avatar"></div>
        </div>

        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item @click="selfInfo">个人信息</el-dropdown-item>
            <el-dropdown-item @click="changePwdDialogV = true">修改密码</el-dropdown-item>
            <el-dropdown-item @click="exitDialogVisible = true">退出登录</el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>
  <div class="tags">
    <el-tag v-for="(tag, index) in tags" :key="index" closable type="primary" @close="removeTag(index)"
            @click="selectTag(tag)">
      {{ tag.name }}
    </el-tag>
  </div>

  <el-dialog v-model="selfInfoDialogVisible" title="个人信息" width="500" align-center @close="infoDialogCancel">
    <user-info v-if="promise === 'user'"></user-info>
    <admin-info v-else-if="promise === 'admin'"></admin-info>
    <employee-info v-else-if="promise === 'employee'"></employee-info>
    <div v-else>待开发</div>

    <template #footer>
      <div class="dialog-footer">
        <el-button @click="infoDialogCancel">取消</el-button>
        <el-button type="primary" @click="updateUser">
          更新
        </el-button>
      </div>
    </template>
  </el-dialog>

  <el-dialog v-model="changePwdDialogV" title="修改密码" width="500" align-center>
    <el-form label-width="100px">
      <el-form-item label="原密码：">
        <el-input v-model="oldPassword" type="password" clearable placeholder="请输入原密码"/>
      </el-form-item>
      <el-form-item label="新密码：">
        <el-input v-model="password" type="password" clearable placeholder="请输入新密码"
                  @blur="validPassword(password)"/>
      </el-form-item>
      <div class="error-word change-pwd" v-if="errWord2.password">{{ errWord2.password }}</div>
      <el-form-item label="再次确认：">
        <el-input v-model="password2" type="password" clearable placeholder="再次输入新密码"
                  @blur="validPassword2(password,password2)"/>
      </el-form-item>
      <div class="error-word change-pwd" v-if="errWord2.password2">{{ errWord2.password2 }}</div>
    </el-form>
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="changePwdCancel">取消</el-button>
        <el-button type="primary" @click="changePwd">修改</el-button>
      </div>
    </template>
  </el-dialog>

  <el-dialog v-model="exitDialogVisible" title="退出" width="300" align-center>
    <span>你确定要退出吗？</span>

    <template #footer>
      <div class="dialog-footer">
        <el-button @click="exitDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="exit">确定</el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>
import {reactive, ref} from "vue";
import {useRouter} from 'vue-router'
import {useUserInfo} from '@/hooks/header/useUserInfo'
import {useAdminInfo} from '@/hooks/header/useAdminInfo'
import {useEmployeeInfo} from "@/hooks/header/useEmployeeInfo";
import {useValid} from '@/hooks/common/useValid'
import request from '@/axios/request'
import {ElMessage, ElMessageBox} from "element-plus";
import UserInfo from "./dialog/UserInfo.vue";
import AdminInfo from "./dialog/AdminInfo.vue";
import EmployeeInfo from "./dialog/EmployeeInfo.vue";
import {useUserInfoStore} from "@/stores/userInfo"
import {useAdminInfoStore} from '@/stores/adminInfo'
import {useEmployeeInfoStore} from '@/stores/employeeInfo'
import {useTagsStore} from "@/stores/tags.js";

import {storeToRefs} from "pinia";
import {useLocalKey} from "@/hooks/common/useLocalKey.js";
import {useSessionKey} from "@/hooks/common/useSessionKey.js";
import {ActGet, ChPwd, Exit} from "@/api/account/account-api.js";

const {NourishAccount, NourishToken, NourishPromise} = useLocalKey()
const sessionKey = useSessionKey()

const promise = ref(sessionStorage.getItem(sessionKey.NourishPromise) || localStorage.getItem(NourishPromise))

// 用户store
const userInfoStore = useUserInfoStore()
const {user, gender} = storeToRefs(userInfoStore)
const {setUser} = userInfoStore

// 管理员store
const adminInfoStore = useAdminInfoStore()
const {admin} = storeToRefs(adminInfoStore)
const {setAdmin} = adminInfoStore

// 员工store
const employeeInfoStore = useEmployeeInfoStore()
const {employee} = storeToRefs(employeeInfoStore)
const {setEmployee} = employeeInfoStore

// 路由
const router = useRouter()

// 路由标签
const tagsStore = useTagsStore()
const {tags} = storeToRefs(tagsStore)
const {removeTag, selectTag} = tagsStore

// 错误提示
const errWord = reactive({
  name: '',
  gender: '',
  phone: '',
  address: '',
  birthday: '',
  email: '',
  position: '',
  salary: ''
})

// 表单参数校验
const {
  validName, validGender, validPhone, validAddress, validBirthday, validEmail,
  validPosition, validSalary
} = useValid(errWord)

// 整体校验一遍
const valid = () => {
  if (promise.value === 'user') {
    // 校验用户
    validName(user.value.name)
    validGender(user.value.gender)
    validPhone(user.value.phone)
    validAddress(user.value.address)
    validBirthday(user.value.birthday)
  } else if (promise.value === 'admin') {
    // 校验管理员
    validName(admin.value.name)
    validEmail(admin.value.email)
  } else if (promise.value === 'employee') {
    validName(employee.value.name)
    validPhone(employee.value.phone)
    validPosition(employee.value.position)
    validSalary(employee.value.salary)
  }
}

// 个人信息对话框
const selfInfoDialogVisible = ref(false);

// 用户Hooks
const {updateUserInfo} = useUserInfo(user, selfInfoDialogVisible)

// 管理员Hooks
const {updateAdminInfo} = useAdminInfo(admin, selfInfoDialogVisible)

// 员工Hooks
const {updateEmployeeInfo} = useEmployeeInfo(employee, selfInfoDialogVisible)

const infoDialogCancel = () => {
  for (const key in errWord) {
    errWord[key] = ''
  }
  selfInfoDialogVisible.value = false
}

// 更新信息
const updateUser = () => {
  valid()
  for (const key in errWord) {
    if (errWord[key] !== '') {
      console.log('type: ', typeof errWord[key])
      ElMessage({message: '用户信息不正确', type: 'error'})
      return
    }
  }
  if (promise.value === 'user') {
    updateUserInfo()
  } else if (promise.value === 'admin') {
    updateAdminInfo()
  } else if (promise.value === 'employee') {
    updateEmployeeInfo()
  } else {
    console.log('待开发');
  }
}

// 获取数据并打开信息对话框
const selfInfo = () => {
  // 向后端请求数据
  request.get(ActGet).then(res => {
    if (res.code === 200) {
      let v = res.data
      if (v.promise === "user") {
        setUser(v.data)
        gender.value = user.value.gender === '男'
      } else if (v.promise === "admin") {
        setAdmin(v.data)
      } else if (v.promise === 'employee') {
        setEmployee(v.data)
      } else {
        ElMessage({message: '尚未开发', type: 'error'})
      }
    } else {
      ElMessage({message: '请求失败，请退出重新登录', type: 'error'})
      localStorage.removeItem(NourishToken)
      localStorage.removeItem(NourishAccount)
      sessionStorage.removeItem(sessionKey.NourishPromise)
      router.replace({name: 'Login'})
    }
  }).catch(err => {
    console.error('请求失败：', err)
  })

  selfInfoDialogVisible.value = true
}

const changePwdDialogV = ref(false)

const oldPassword = ref('')
const password = ref('')
const password2 = ref('')

const errWord2 = reactive({
  password: '',
  password2: ''
})

const {validPassword, validPassword2} = useValid(errWord2)

const changePwdCancel = () => {
  oldPassword.value = ''
  password.value = ''
  password2.value = ''
  changePwdDialogV.value = false
}

const changePwd = () => {
  validPassword(password.value)
  validPassword2(password.value, password2.value)
  for (let key in errWord2) {
    if (errWord2[key] !== '') {
      ElMessage({message: '校验不通过', type: 'error'})
      return
    }
  }

  ElMessageBox.confirm('您确定要修改密码吗', 'Warning', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  }).then(() => {
    request.put(ChPwd, {
      oldPassword: oldPassword.value,
      newPassword: password.value
    }).then(res => {
      if (res.code === 200) {
        exit()
        ElMessage({message: '修改密码成功，请重新登录', type: 'success'})
      } else {
        ElMessage({message: '参数错误', type: 'error'})
      }
    }).catch(err => {
      console.error(err)
      ElMessage({message: '系统错误', type: 'error'})
    })
  }).catch(() => {
    ElMessage({message: '取消更改', type: 'info'})
  })
}

const exitDialogVisible = ref(false)

const exit = async () => {
  // 先向后端发送删除token的请求，然后再删除前端的storage
  try {
    let res = await request.get(Exit)
    if (res.code === 200) {
      ElMessage({message: res.data + '退出成功', type: 'success'})
    } else {
      ElMessage({message: '参数错误', type: 'error'})
    }
  } catch (e) {
    console.error(e)
    ElMessage({message: '服务器端错误', type: 'error'})
  }
  localStorage.removeItem(NourishToken)
  localStorage.removeItem(NourishAccount)
  exitDialogVisible.value = false
  await router.push({name: "Login"})
}

</script>

<style lang="scss" scoped>
@import "@/scss/home/header/header.scss";
@import "@/scss/common/ErrorWord.scss";

.left-space {
  margin-left: 100px;
  margin-top: -18px;
}

.change-pwd {
  margin-left: 100px;
  margin-top: -18px;
}

</style>