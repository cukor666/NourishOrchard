<template>
  <div class="container">
    <el-button class="search-button" @click="changeSearchDialog">
      <el-icon>
        <Search/>
      </el-icon>
      <span style="margin-left: 5px">搜索</span>
    </el-button>
    <search-user :searchDialogV="searchDialogV" @closeDialog="handleCloseDialog" @search="search"></search-user>
    <el-card style="width: 99%">
      <el-table :data="userList" stripe style="width: 100%">
        <el-table-column prop="id" label="id" width="50"/>
        <el-table-column prop="username" label="账号" width="180"/>
        <el-table-column prop="name" label="姓名" width="180"/>
        <el-table-column prop="gender" label="性别" width="60"/>
        <el-table-column prop="phone" label="联系电话" width="180"/>
        <el-table-column prop="address" label="联系地址" width="200"/>
        <el-table-column prop="birthday" label="出生日期" width="180"/>
        <el-table-column fixed="right" label="操作" width="120">
          <template v-slot="scope">
            <el-button link type="primary" size="small" @click="recoverUser(scope.row)">恢复</el-button>
            <el-button link type="primary" size="small" @click="deleteUser(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
    <el-pagination style="margin-top: 5px"
                   v-model:current-page="currentPage"
                   v-model:page-size="pageSize"
                   :page-sizes="pageSizes"
                   background
                   layout="total, sizes, prev, pager, next, jumper"
                   :total="total"
                   @size-change="handleSizeChange"
                   @current-change="handleCurrentChange"
    />
  </div>
</template>

<script setup>
import SearchUser from "@/components/user/SearchUser.vue";

import {Search} from "@element-plus/icons-vue";
import {computed, onMounted, ref} from "vue";
import request from "@/axios/request.js";
import {ElMessage} from "element-plus";

const searchDialogV = ref(false)
const changeSearchDialog = () => {
  searchDialogV.value = true
}

const handleCloseDialog = () => {
  searchDialogV.value = false
}

const searchUser = ref({
  id: 0,
  username: '',
  name: '',
  gender: '',
  phone: '',
  address: '',
  birthday: ''
})

const search = async (u) => {
  searchUser.value = u
  console.log('searchUser = ', searchUser.value)
  searchDialogV.value = false
  // 向服务器端做查询，然后将结果放到userList中
  try {
    let res = await request.get('/user/logout-list', {
      params: {
        pageSize: pageSize.value,
        pageNum: currentPage.value,
        ...searchUser.value
      }
    })
    if (res.code === 200) {
      total.value = res.data.total
      userList.value = res.data.users
      ElMessage({message: '查询成功', type: 'success'})
    } else {
      ElMessage({message: '查询失败', type: 'error'})
    }
  } catch (err) {
    console.error(err)
    ElMessage({message: '服务器错误', type: 'error'})
  }
}

// 展示数据的用户列表
const userList = ref([{
  id: 1,
  username: 'CZKJ4634',
  name: 'CukorZhong',
  gender: '男',
  phone: '18577659826',
  address: '广东省深圳市福田区下梅林',
  birthday: '2000-08-07'
}])

const userInfoUpdated = computed(() => {
  return sessionStorage.getItem('nourish-logout-user-info-updated') || "false"
})

const updateUser = () => {
  console.log('执行了更新用户的函数')
  sessionStorage.setItem('nourish-logout-user-info-updated', "true")
}

// 分页部分
const currentPage = ref(1)
const pageSize = ref(3)
const pageSizes = ref([3, 6, 10, 15, 20])
const total = ref(Number(sessionStorage.getItem('nourish-logout-user-total')) || 0)

// 改变pageSize
const handleSizeChange = async () => {
  try {
    let res = await request.get('/user/logout-list', {
      params: {
        pageSize: pageSize.value,
        pageNum: currentPage.value,
        ...searchUser.value
      }
    })
    if (res.code === 200) {
      total.value = res.data.total
      userList.value = res.data.users
    } else {
      console.log(res.msg)
      ElMessage({message: '参数错误', type: 'error'})
    }
  } catch (err) {
    console.error(err)
    ElMessage({message: '服务器错误', type: 'error'})
  }
}

// 改变currentPage
const handleCurrentChange = async () => {
  try {
    let res = await request.get('/user/logout-list', {
      params: {
        pageSize: pageSize.value,
        pageNum: currentPage.value,
        ...searchUser.value
      }
    })
    if (res.code === 200) {
      total.value = res.data.total
      userList.value = res.data.users
    } else {
      console.log(res.msg)
      ElMessage({message: '参数错误', type: 'error'})
    }
  } catch (err) {
    console.error(err)
    ElMessage({message: '服务器错误', type: 'error'})
  }
}

onMounted(async () => {
  // 从sessionStorage中获取，如果没有则访问服务器获取
  let users = sessionStorage.getItem('nourish-logout-user-list');
  total.value = Number(sessionStorage.getItem('nourish-logout-user-total'))
  if (users === null || userInfoUpdated.value === "true") {
    // 从服务器端获取
    try {
      let res = await request.get('/user/logout-list', {
        params: {
          pageSize: pageSize.value,
          pageNum: currentPage.value
        }
      })
      if (res.code === 200) {
        let v = res.data
        total.value = v.total
        userList.value = v.users
        users = JSON.stringify(userList.value)
        sessionStorage.setItem('nourish-logout-user-list', users)
        sessionStorage.setItem('nourish-logout-user-total', total.value.toString())
        sessionStorage.removeItem('nourish-logout-user-info-updated')
      } else {
        ElMessage({message: '参数错误', type: 'error'})
      }
    } catch (err) {
      console.error(err)
      ElMessage({message: '服务器错误', type: 'error'})
    }
  } else {
    userList.value = JSON.parse(users);
  }
})

const user = ref({
  id: 6,
  username: "CukorZhong",
  name: "",
  gender: "",
  phone: "",
  address: "",
  birthday: ""
})

// 恢复用户
const recoverUser = async (item) => {
  try {
    let res = await request.post('/user/recover', {
      username: item.username
    })
    if (res.code === 200) {
      console.log(res.data)
      ElMessage({message: '恢复用户成功', type: 'success'})
    } else {
      ElMessage({message: '恢复用户失败', type: 'error'})
    }
    res = await request.get('/user/logout-list', {
      params: {
        pageSize: pageSize,
        pageNum: currentPage
      }
    })
    if (res.code === 200) {
      console.log('更新注销用户列表成功')
      let v = res.data
      total.value = v.total
      userList.value = v.users
      sessionStorage.setItem('nourish-logout-user-list', JSON.stringify(userList.value))
      sessionStorage.setItem('nourish-logout-user-total', total.value.toString())
      sessionStorage.removeItem('nourish-logout-user-info-updated')
    } else {
      ElMessage({message: '参数错误', type: 'error'})
    }
  } catch (err) {
    console.error(err)
    ElMessage({message: '系统错误', type: 'error'})
  }
}

const deleteUser = (item) => {
  console.log(item)
}

</script>

<style lang="scss" scoped>
@import "@/scss/home/main/search-button.scss";
</style>