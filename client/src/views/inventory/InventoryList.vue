<template>
  <div class="container">
    <div style="display: flex">
      <el-button class="search-button" @click="changeSearchDialog">
        <el-icon>
          <Search/>
        </el-icon>
        <span style="margin-left: 5px">搜索</span>
      </el-button>
      <el-button @click="addWarehouse">添加</el-button>
    </div>

    <search-inventory :search-dialog-v="searchDialogV" @closeDialog="closeSearchDialog" @search="search"/>

    <el-card style="width: 99%">
      <el-table :data="inventoryList" stripe style="width: 100%">
        <el-table-column prop="id" label="id" width="50"/>
        <el-table-column prop="commodityId" label="商品id" width="100"/>
        <el-table-column prop="quantity" label="数量" width="100"/>
        <el-table-column prop="employeeId" label="员工id" width="100"/>
        <el-table-column prop="warehouseId" label="仓库id" width="100"/>
        <el-table-column prop="createdAt" label="创建时间" width="250"/>
        <el-table-column prop="updatedAt" label="更新时间" width="250"/>

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

    <!--库存信息详情弹窗-->
    <detail :detail-dialog-v="detailDialogV" :inventory="inventory" @closeDialog="closeDetailDialog"
            @updateInventory="updateInventory"></detail>

    <!--添加库存信息弹窗-->
    <add :add-inventory-dialog-v="addInventoryDialogV" @cancelAddInventory="cancelAddInventory"
         @closeDialog="closeAddDialog"></add>
  </div>
</template>

<script setup>
import SearchInventory from "@/components/inventory/SearchInventory.vue";
import Detail from "@/components/inventory/dialog/Detail.vue";
import Add from "@/components/inventory/dialog/Add.vue";

import {Search} from "@element-plus/icons-vue";
import {useSearch} from "@/hooks/list/inventory/useSearch.js"
import {usePage} from "@/hooks/list/usePage.js";
import {useTable} from "@/hooks/list/inventory/useTable.js";
import {onMounted, ref} from "vue";
import {useDetail} from "@/hooks/list/inventory/useDetail.js";
import request from "@/axios/request.js";
import {ElMessage} from "element-plus";
import {InventoryList} from "@/api/inventory/inventory-api.js";

const {searchDialogV, searchInventory, changeSearchDialog, closeSearchDialog, findInventory} = useSearch();
const inventoryList = ref([{...searchInventory.value}])
const {detailDialogV, inventory, closeDetailDialog, updateInventory} = useDetail()
const {currentPage, pageSize, pageSizes, total} = usePage()
const {updateList, showDetail, deleteInventory} = useTable(inventoryList, pageSize, currentPage, total)

const search = (query) => {
  findInventory(query, pageSize, currentPage, total, inventoryList);
}

// 详情按钮点击事件
const detailInfo = (row) => {
  showDetail(row, detailDialogV, inventory);
}

// 删除按钮点击事件
const deleteLink = (row) => {
  deleteInventory(row);
}

// 页码改变事件
const pageSizeChange = () => {
  updateList(searchInventory)
}

// 当前页码改变事件
const currentPageChange = () => {
  updateList(searchInventory)
}

const addInventoryDialogV = ref(false)

const addWarehouse = () => {
  addInventoryDialogV.value = true
}

const cancelAddInventory = () => {
  addInventoryDialogV.value = false
}
const closeAddDialog = () => {
  addInventoryDialogV.value = false
  updateList(searchInventory)
}

onMounted(async () => {
  // 从服务器端获取
  try {
    let res = await request.get(InventoryList, {
      params: {
        pageSize: pageSize.value,
        pageNum: currentPage.value
      }
    })
    if (res.code === 200) {
      let v = res.data
      total.value = v.total
      inventoryList.value = v.inventors
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