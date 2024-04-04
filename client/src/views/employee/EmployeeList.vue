<template>
  <div class="container">
    <el-button class="search-button" @click="changeSearchDialog">
      <el-icon>
        <Search/>
      </el-icon>
      <span style="margin-left: 5px">搜索</span>
    </el-button>
    <search-employee :search-dialog-v="searchDialogV" @closeDialog="closeSearchDialog"
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
          <el-button link type="primary" size="small" @click="promotionHandle(scope.row)">晋升</el-button>
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

  <el-dialog v-model="promotionDV" title="晋升管理员">
    <el-form label-width="100px">
      <el-form-item label="密码：">
        <el-input v-model="password" type="password" placeholder="请输入当前登录的管理员密码" />
      </el-form-item>

      <el-form-item label="邮箱：">
        <el-input v-model="email" placeholder="请输入晋升管理员的邮箱" />
      </el-form-item>

      <el-form-item label="备注：">
        <el-input v-model="mark" placeholder="请输入晋升备注（可选）" />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button type="primary" @click="proPwdConfirm">确认</el-button>
      <el-button @click="cancel">取消</el-button>
    </template>
  </el-dialog>
</template>

<script setup>
import {Search} from "@element-plus/icons-vue";
import SearchEmployee from "@/components/employee/SearchEmployee.vue";
import {onMounted, ref} from "vue";
import request from "@/axios/request.js";
import {ElMessage} from "element-plus";
import Detail from "@/components/employee/dialog/Detail.vue";
import {EmpList} from "@/api/emp/emp-api.js";
import {useTable} from "@/hooks/list/emp/useTable.js";
import {useSearch} from "@/hooks/list/emp/useSearch.js";
import {usePage} from "@/hooks/list/usePage.js";

const {searchDialogV, searchEmp, changeSearchDialog, closeSearchDialog, findEmp} = useSearch()
const {currentPage, pageSize, pageSizes, total} = usePage()
const employeeList = ref([{...searchEmp.value}])

const { updateList, updateInfo, promotion} = useTable(employeeList, pageSize, currentPage, total)


onMounted(async () => {
  try {
    let res = await request.get(EmpList, {
      params: {
        pageSize: pageSize.value,
        pageNum: currentPage.value,
        ...searchEmp.value
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

const search = (emp) => {
  findEmp(emp, pageSize, currentPage, total, employeeList)
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
  detailDialogV.value = false
}

const handleSizeChange = () => {
  updateList(searchEmp)
}

const handleCurrentChange = () => {
  updateList(searchEmp)
}

const promotionDV = ref(false)
const password = ref('')
const email = ref('')
const mark = ref('')

const proPwdConfirm = () => {
  // 从sessionStorage中读取proEmp
  let proEmpStr = sessionStorage.getItem('proEmp');
  const proEmp = JSON.parse(proEmpStr)

  let e = {
    username: proEmp.username,
    password: password.value,
    name: proEmp.name,
    email: email.value,
    mark: mark.value
  }

  promotion(e)
  updateList(searchEmp)

  promotionDV.value = false
}

const cancel = () => {
  console.log('点击了取消')
  password.value = ''
  promotionDV.value = false
}

const promotionHandle =  (item) => {
  promotionDV.value = true
  // 把item的数据存到sessionStorage中
  const proEmpStr = JSON.stringify(item);
  sessionStorage.setItem('proEmp', proEmpStr)
}

</script>

<style scoped lang="scss">
@import "@/scss/home/main/search-button.scss";
</style>