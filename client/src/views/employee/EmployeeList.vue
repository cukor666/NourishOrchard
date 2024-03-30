<template>
  <div class="container">
    <el-button class="search-button" @click="changeSearchDialog">
      <el-icon>
        <Search/>
      </el-icon>
      <span style="margin-left: 5px">搜索</span>
    </el-button>
    <search-employee :search-dialog-v="searchDialogV" @closeDialog="handleCloseDialog"
                     @search="search"></search-employee>
  </div>
  <el-card style="width: 99%;">
    <el-table :data="employeeList" stripe style="width: 100%">
      <el-table-column prop="id" label="id" width="50"/>
      <el-table-column prop="username" label="账号" width="180"/>
      <el-table-column prop="name" label="姓名" width="180"/>
      <el-table-column prop="phone" label="联系电话" width="180"/>
      <el-table-column prop="position" label="职位" width="180"/>
      <el-table-column prop="salary" label="薪资" width="180"/>
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

  <!--员工详情对话框-->
  <detail :detail-dialog-v="detailDialogV" :employee="employee" @closeDialog="handleCloseDetailDialog"
          @update-employee="updateEmployee"></detail>
</template>

<script setup>
import {Search} from "@element-plus/icons-vue";
import SearchEmployee from "@/components/employee/SearchEmployee.vue";
import {onMounted, ref} from "vue";
import request from "@/axios/request.js";
import {ElMessage} from "element-plus";
import Detail from "@/components/employee/dialog/Detail.vue";
import {EmpList} from "@/api/emp/emp-api.js";

const searchDialogV = ref(false)

const searchEmployee = ref({
  id: 0,
  username: '',
  name: '',
  phone: '',
  position: '',
  salary: 0
})

const employeeList = ref([{...searchEmployee.value}])

const total = ref(0)
const pageSize = ref(3)
const currentPage = ref(1)
const pageSizes = ref([3, 6, 10])

onMounted(async () => {
  try {
    let res = await request.get(EmpList, {
      params: {
        pageSize: pageSize.value,
        pageNum: currentPage.value,
        ...searchEmployee.value
      }
    })
    if (res.code === 200) {
      total.value = res.data.total
      employeeList.value = res.data.employees
    } else {
      ElMessage({message: '参数错误', type: 'error'})
    }
  } catch (err) {
    console.error(err)
    ElMessage({message: '服务器错误', type: 'error'})
  }
})

const changeSearchDialog = () => {
  searchDialogV.value = true
}

const handleCloseDialog = () => {
  searchDialogV.value = false
}

const search = (emp) => {

}

const detailDialogV = ref(false)
const employee = ref({
  id: 6,
  username: 'CukorZhong',
  name: '',
  phone: '',
  position: '',
  salary: 0
})

const detailInfo = (item) => {
  detailDialogV.value = true
  employee.value = item
}

const handleCloseDetailDialog = () => {
  detailDialogV.value = false
}

const updateEmployee = () => {
  console.log('hahaha')
}

const deleteUser = (item) => {

}

const handleSizeChange = () => {

}

const handleCurrentChange = () => {

}

</script>

<style scoped lang="scss">
@import "@/scss/home/main/search-button.scss";
</style>