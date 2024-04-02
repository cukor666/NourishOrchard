<template>
  <div class="container">
    <el-button class="search-button" @click="changeSearchDialog">
      <el-icon>
        <Search/>
      </el-icon>
      <span style="margin-left: 5px">搜索</span>
    </el-button>
    <search-user :searchDialogV="searchDialogV" @closeDialog="closeSearchDialog" @search="search"></search-user>
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
import {onMounted, ref} from "vue";
import request from "@/axios/request.js";
import {ElMessage} from "element-plus";
import {useSearch} from "@/hooks/list/logout-user/useSearch.js"
import {useTable} from "@/hooks/list/logout-user/useTable.js";
import {usePage} from "@/hooks/list/usePage.js";
import {LogoutUserList} from "@/api/logout-user/logout-user-api.js";

const {searchDialogV, searchUser, changeSearchDialog, closeSearchDialog, findUser} = useSearch()
const userList = ref([{...searchUser.value}])
const {currentPage, pageSize, pageSizes, total} = usePage()
const {updateList, recoverUser, deleteUser} = useTable(userList, pageSize, currentPage, total)

const search = async (u) => {
  findUser(u, pageSize, currentPage, total, userList)
}

const updateUser = () => {
  console.log('执行了更新用户的函数')
}

// 改变pageSize
const handleSizeChange = () => {
  updateList(searchUser)
}

// 改变currentPage
const handleCurrentChange = () => {
  updateList(searchUser)
}

onMounted(async () => {
  // 从服务器端获取
  try {
    let res = await request.get(LogoutUserList, {
      params: {
        pageSize: pageSize.value,
        pageNum: currentPage.value
      }
    })
    if (res.code === 200) {
      let v = res.data
      total.value = v.total
      userList.value = v.users
    } else {
      ElMessage({message: '参数错误', type: 'error'})
    }
  } catch (err) {
    console.error(err)
    ElMessage({message: '服务器错误', type: 'error'})
  }
})


</script>

<style lang="scss" scoped>
@import "@/scss/home/main/search-button.scss";
</style>