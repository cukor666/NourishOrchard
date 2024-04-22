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


    <search-warehouse :search-dialog-v="searchDialogV" @closeDialog="closeSearchDialog" @search="search"/>

    <el-card style="width: 99%">
      <el-table :data="warehouseList" stripe style="width: 100%">
        <el-table-column prop="id" label="id" width="50"/>
        <el-table-column prop="address" label="仓库地址" width="200"/>
        <el-table-column prop="capacity" label="仓库容量" width="150"/>
        <el-table-column prop="remaining" label="剩余容量" width="150"/>

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

    <!--仓库详情弹窗-->
    <detail :detail-dialog-v="detailDialogV" :warehouse="warehouse" @closeDialog="closeDetailDialog"
            @updateWarehouse="updateWarehouse"></detail>

    <!--添加仓库弹窗-->
    <add :add-warehouse-dialog-v="addWarehouseDialogV" @cancelAddWarehouse="cancelAddWarehouse"
     @close="closeAddDialog"></add>
  </div>
</template>

<script setup>
import SearchWarehouse from "@/components/warehouse/SearchWarehouse.vue";
import Detail from "@/components/warehouse/dialog/Detail.vue";
import Add from "@/components/warehouse/dialog/Add.vue";

import {Search} from "@element-plus/icons-vue";
import {useSearch} from "@/hooks/list/warehouse/useSearch.js"
import {usePage} from "@/hooks/list/usePage.js";
import {useTable} from "@/hooks/list/warehouse/useTable.js";
import {onMounted, ref} from "vue";
import {useDetail} from "@/hooks/list/warehouse/useDetail.js";
import request from "@/axios/request.js";
import {ElMessage} from "element-plus";
import {WarehouseList} from "@/api/warehouse/warehouse-api.js";

const {searchDialogV, searchWarehouse, changeSearchDialog, closeSearchDialog, findWarehouse} = useSearch();
const warehouseList = ref([{...searchWarehouse.value}])
const {detailDialogV, warehouse, closeDetailDialog, updateWarehouse} = useDetail()
const {currentPage, pageSize, pageSizes, total} = usePage()
const {updateList, showDetail, deleteWarehouse} = useTable(warehouseList, pageSize, currentPage, total)

const search = (query) => {
  findWarehouse(query, pageSize, currentPage, total, warehouseList);
}

// 详情按钮点击事件
const detailInfo = (row) => {
  showDetail(row, detailDialogV, warehouse);
}

// 删除按钮点击事件
const deleteLink = (row) => {
  deleteWarehouse(row);
}

// 页码改变事件
const pageSizeChange = () => {
  updateList(searchWarehouse)
}

// 当前页码改变事件
const currentPageChange = () => {
  updateList(searchWarehouse)
}

const addWarehouseDialogV = ref(false)

const addWarehouse = () => {
  addWarehouseDialogV.value = true
}

const cancelAddWarehouse = () => {
    addWarehouseDialogV.value = false
}
const closeAddDialog = () => {
  addWarehouseDialogV.value = false
  updateList(searchWarehouse)
}

onMounted(async () => {
  // 从服务器端获取
  try {
    let res = await request.get(WarehouseList, {
      params: {
        pageSize: pageSize.value,
        pageNum: currentPage.value
      }
    })
    if (res.code === 200) {
      let v = res.data
      total.value = v.total
      warehouseList.value = v.warehouses
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