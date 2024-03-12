<template>
  <div class="container">
    <el-button class="search-button" @click="changeSearchDialog">
      <el-icon>
        <Search/>
      </el-icon>
      <span style="margin-left: 5px">搜索</span>
    </el-button>
    <search-user :searchDialogV="searchDialogV" @closeDialog="handleCloseDialog"></search-user>
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
            <el-button link type="primary" size="small" @click="detailInfo(scope.row)">详情</el-button>
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

    <!--用户详情对话框-->
    <detail :detail-dialog-v="detailDialogV" :user="user" @closeDialog="handleCloseDetailDialog"
            @updateUser="updateUser"></detail>
  </div>
</template>

<script setup>
import SearchUser from "@/components/user/SearchUser.vue";
import Detail from "@/components/user/dialog/Detail.vue";

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

// 展示数据的用户列表
const userList = ref([
  {
    id: 1,
    username: 'CZKJ4634',
    name: 'CukorZhong',
    gender: '男',
    phone: '18577659826',
    address: '广东省深圳市福田区下梅林',
    birthday: '2000-08-07'
  }
])

const userInfoUpdated = computed(() => {
  return sessionStorage.getItem('nourish-user-info-updated') || "false"
})

const updateUser = () => {
  console.log('执行了更新用户的函数')
  sessionStorage.setItem('nourish-user-info-updated', "true")
}

// 分页部分
const currentPage = ref(1)
const pageSize = ref(3)
const pageSizes = ref([3, 6, 10, 15, 20])
const total = ref(Number(sessionStorage.getItem('nourish-user-total')) || 0)

// 改变pageSize
const handleSizeChange = () => {
  request.get('/user/list', {
    params: {
      pageSize: pageSize.value,
      pageNum: currentPage.value
    }
  }).then(res => {
    let v = res.data
    total.value = v.total
    userList.value = v.users
    sessionStorage.setItem('nourish-user-list', JSON.stringify(userList.value))
  }).catch(err => {
    ElMessage({
      message: '服务器出错',
      type: 'error'
    })
    console.log(err)
  })
}

// 改变currentPage
const handleCurrentChange = () => {
  request.get('/user/list', {
    params: {
      pageSize: pageSize.value,
      pageNum: currentPage.value
    }
  }).then(res => {
    let v = res.data
    total.value = v.total
    userList.value = v.users
    sessionStorage.setItem('nourish-user-list', JSON.stringify(userList.value))
  }).catch(err => {
    ElMessage({
      message: '服务器出错',
      type: 'error'
    })
    console.log(err)
  })
}

onMounted(() => {
  // 从sessionStorage中获取，如果没有则访问服务器获取
  let users = sessionStorage.getItem('nourish-user-list');
  total.value = Number(sessionStorage.getItem('nourish-user-total'))
  if (users === null || userInfoUpdated.value === "true") {
    // 从服务器端获取
    request.get('/user/list', {
      params: {
        pageSize: pageSize.value,
        pageNum: currentPage.value
      }
    }).then(res => {
      let v = res.data
      total.value = v.total
      userList.value = v.users
      users = JSON.stringify(userList.value)
      sessionStorage.setItem('nourish-user-list', users)
      sessionStorage.setItem('nourish-user-total', total.value.toString())
      sessionStorage.removeItem('nourish-user-info-updated')
    }).catch(err => {
      console.log(err)
    })
  } else {
    userList.value = JSON.parse(users);
  }
})

const detailDialogV = ref(false)

const user = ref({
  id: 6,
  username: "CukorZhong",
  name: "",
  gender: "",
  phone: "",
  address: "",
  birthday: ""
})

const handleCloseDetailDialog = () => {
  detailDialogV.value = false
}

const detailInfo = (item) => {
  console.log(item)
  detailDialogV.value = true
  user.value = item
}

// 删除用户，将删除的用户存入到注销表中
const deleteUser = async (item) => {
  try {
    let res = await request.delete('/user/delete', {
      params: {
        username: item.username
      }
    })
    if (res.code === 200) {
      ElMessage({message: '删除成功', type: 'success'})
    } else {
      ElMessage({message: '删除失败', type: 'error'})
    }
    res = await request.get('/user/logout-list', {
      params: {
        pageSize: pageSize,
        pageNum: currentPage
      }
    })
    if (res.code !== 200) {
      ElMessage({message: '未能及时更新列表，请刷新', type: 'warning'})
    }
  } catch {
    ElMessage({message: '系统错误', type: 'error'})
  }
}

</script>

<style lang="scss" scoped>
.search-button {
  display: flex;
  align-items: center;
  width: 60px;
  transition: all 1s;
}

.search-button:hover {
  background: #2394f0;
  color: #eef4f8;
  box-shadow: 0 0 10px #2394f0;
  border: none;
  width: 200px;
  font-size: 20px;
}
</style>