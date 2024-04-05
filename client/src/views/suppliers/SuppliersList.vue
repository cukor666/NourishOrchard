<template>
  <div class="container">
    <div style="display: flex">
      <el-button class="search-button" @click="changeSearchDialog">
        <el-icon>
          <Search/>
        </el-icon>
        <span style="margin-left: 5px">搜索</span>
      </el-button>
      <el-button @click="addSupplier">添加</el-button>
    </div>


    <search-supplier :search-dialog-v="searchDialogV" @closeDialog="closeSearchDialog" @search="search"></search-supplier>

    <el-card style="width: 99%">
      <el-table :data="supplierList" stripe style="width: 100%">
        <el-table-column prop="id" label="id" width="50"/>
        <el-table-column prop="name" label="姓名" width="180"/>
        <el-table-column prop="address" label="供应商地址" width="200"/>
        <el-table-column prop="contactPerson" label="联系人名" width="200"/>
        <el-table-column prop="phone" label="联系电话" width="180"/>
        <el-table-column prop="reputation" label="信誉" width="180"/>
        <el-table-column fixed="right" label="操作" width="120">
          <template v-slot="scope">
            <el-button link type="primary" size="small" @click="detailInfo(scope.row)">详情</el-button>
            <el-button link type="primary" size="small" @click="deleteLink(scope.row)">删除</el-button>
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
                   @size-change="pageSizeChange"
                   @current-change="currentPageChange"
    />

    <!--供应商详情对话框-->
    <detail :detail-dialog-v="detailDialogV" :supplier="supplier" @closeDialog="closeDetailDialog"
            @updateSupplier="updateSupplier"></detail>
    <!--添加供应商-->
    <add :add-supplier-dialog-v="addSupplierDialogV" @closeAddDialog="closeAddDialog" @cancelAddSupplier="cancelAddSupplier"></add>
  </div>
</template>

<script setup>
import SearchSupplier from "@/components/supplier/SearchSupplier.vue";
import {Search} from "@element-plus/icons-vue";
import {useSearch} from "@/hooks/list/supplier/useSearch.js";
import {useTable} from "@/hooks/list/supplier/useTable.js";
import {onMounted, ref} from "vue";
import {usePage} from "@/hooks/list/usePage.js";
import {useDetail} from "@/hooks/list/supplier/useDetail.js";
import Detail from "@/components/supplier/dialog/Detail.vue";
import request from "@/axios/request.js";
import {ElMessage} from "element-plus";
import {SupplierList} from "@/api/supplier/sup-api.js";
import Add from "@/components/supplier/dialog/Add.vue";

const {searchDialogV, searchSupplier, changeSearchDialog, closeSearchDialog, findSupplier} = useSearch()
const supplierList = ref([{...searchSupplier.value}])
const {detailDialogV, supplier, closeDetailDialog, updateSupplier} = useDetail()
const {currentPage, pageSize, pageSizes, total} = usePage()
const {updateList, showDetail, deleteSupplier} = useTable(supplierList, pageSize, currentPage, total)

const search = (u) => {
  findSupplier(u, pageSize, currentPage, total, supplierList)
}

const detailInfo = (item) => {
  showDetail(item, detailDialogV, supplier)
}

const deleteLink = (item) => {
  deleteSupplier(item)
}

const pageSizeChange = () => {
  updateList(searchSupplier)
}

const currentPageChange = () => {
  updateList(searchSupplier)
}

const addSupplierDialogV = ref(false)

const addSupplier = () => {
  console.log('添加供应商')
  addSupplierDialogV.value = true
}

const closeAddDialog = () => {
  addSupplierDialogV.value = false
  updateList(searchSupplier)
}

const cancelAddSupplier = () => {
  addSupplierDialogV.value = false
}

onMounted(async () => {
  // 从服务器端获取
  try {
    let res = await request.get(SupplierList, {
      params: {
        pageSize: pageSize.value,
        pageNum: currentPage.value
      }
    })
    if (res.code === 200) {
      let v = res.data
      total.value = v.total
      supplierList.value = v.suppliers
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


<style scoped lang="scss">
@import "@/scss/home/main/search-button.scss";
</style>