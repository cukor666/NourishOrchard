<template>
  <div class="container">
    <el-button class="search-button" @click="changeSearchDialog">
      <el-icon>
        <Search/>
      </el-icon>
      <span style="margin-left: 5px">搜索</span>
    </el-button>

    <search-admin :searchDialogV="searchDialogV" @closeDialog="closeSearchDialog" @search="search"></search-admin>

    <el-card style="width: 99%">
      <el-table :data="adminList" stripe style="width: 100%">
        <el-table-column prop="id" label="id" width="50"/>
        <el-table-column prop="username" label="账号" width="180"/>
        <el-table-column prop="name" label="姓名" width="150"/>
        <el-table-column prop="email" label="邮箱" width="200"/>
        <!--          <el-table-column fixed="right" label="操作" width="120">-->
        <!--            <template v-slot="scope">-->
        <!--              <el-button link type="primary" size="small" @click="deleteLink(scope.row)">删除</el-button>-->
        <!--            </template>-->
        <!--          </el-table-column>-->
      </el-table>
    </el-card>

    <el-pagination style="margin-top: 5px"
                   v-model:current-page="currentPage"
                   v-model:page-size="pageSize"
                   :page-sizes="pageSizes"
                   background
                   layout="total, sizes, prev, pager, next, jumper"
                   :total="total"
                   @size-change="pageSizeChange"
                   @current-change="currentPageChange"
    />
  </div>
</template>

<script setup>
import SearchAdmin from "@/components/admin/SearchAdmin.vue";

import {Search} from "@element-plus/icons-vue";
import {useSearch} from "@/hooks/list/admin/useSearch.js";
import {onMounted, ref} from "vue";
import {usePage} from "@/hooks/list/usePage.js";
import {useTable} from "@/hooks/list/admin/useTable.js";
import request from "@/axios/request.js";
import {ElMessage} from "element-plus";
import {AdminList} from "@/api/admin/admin-api.js";

const {searchDialogV, searchAdmin, changeSearchDialog, closeSearchDialog, findAdmin} = useSearch()
const {currentPage, pageSize, pageSizes, total} = usePage()
const adminList = ref([{...searchAdmin.value}])
const {updateList} = useTable(adminList, pageSize, currentPage, total)

const search = (u) => {
  findAdmin(u, pageSize, currentPage, total, adminList)
}

const pageSizeChange = () => {
  updateList(searchAdmin)
}

const currentPageChange = () => {
  updateList(searchAdmin)
}

onMounted(async () => {
  // 从服务器端获取
  try {
    let res = await request.get(AdminList, {
      params: {
        pageSize: pageSize.value,
        pageNum: currentPage.value
      }
    })
    if (res.code === 200) {
      let v = res.data
      total.value = v.total
      adminList.value = v.admins
    } else {
      console.log(res.msg)
      ElMessage({message: '参数错误', type: 'error'})
    }
  } catch (err) {
    console.error(err)
    ElMessage({message: '服务端错误', type: 'error'})
  }

})

</script>

<style lang="scss" scoped>
@import "@/scss/home/main/search-button.scss";
</style>