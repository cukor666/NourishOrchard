<template>
    <div class="container">
      <div style="display: flex">
        <el-button class="search-button" @click="changeSearchDialog">
          <el-icon>
            <Search/>
          </el-icon>
          <span style="margin-left: 5px">搜索</span>
        </el-button>
        <el-button @click="addFruit">添加</el-button>
      </div>


      <search-fruit :searchDialogV="searchDialogV" @closeDialog="closeSearchDialog" @search="search"></search-fruit>

      <el-card style="width: 99%">
        <el-table :data="fruitList" stripe style="width: 100%">
          <el-table-column prop="id" label="id" width="50"/>
          <el-table-column prop="name" label="水果名称" width="180"/>
          <el-table-column prop="water" label="含水量" width="80"/>
          <el-table-column prop="sugar" label="含糖量" width="80"/>
          <el-table-column prop="shelfLife" label="保质期" width="80"/>
          <el-table-column prop="origin" label="原产地" width="180"/>
          <el-table-column prop="supplierId" label="供应商Id" width="80"/>

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

      <!--水果详情对话框-->
      <detail :detail-dialog-v="detailDialogV" :fruit="fruit" @closeDialog="closeDetailDialog"
              @updateFruit="updateFruitHandle"></detail>
      <!--添加水果-->
      <add :add-fruit-dialog-v="addFruitDialogV" @cancelAddFruit="cancelAddFruit" @closeAddDialog="closeAddDialog"></add>
    </div>
</template>

<script setup>
import SearchFruit from "@/components/fruit/SearchFruit.vue";
import Detail from "@/components/fruit/dialog/Detail.vue";
import Add from "@/components/fruit/dialog/Add.vue";

import {Search} from "@element-plus/icons-vue";
import {useSearch} from "@/hooks/list/fruit/useSearch.js";
import {useDetail} from "@/hooks/list/fruit/useDetail.js";
import {onMounted, ref} from "vue";
import {usePage} from "@/hooks/list/usePage.js";
import {useTable} from "@/hooks/list/fruit/useTable.js";
import request from "@/axios/request.js";
import {ElMessage} from "element-plus";
import {FruitList} from "@/api/fruit/fruit-api.js";
const {searchDialogV, searchFruit, changeSearchDialog, closeSearchDialog, findFruit} = useSearch()
const {detailDialogV, fruit, closeDetailDialog, updateFruit} = useDetail()
const {currentPage, pageSize, pageSizes, total} = usePage()

const fruitList = ref([{...searchFruit.value}])
const {updateList, showDetail, deleteFruit} = useTable(fruitList, pageSize, currentPage, total)


const search = (u) => {
  findFruit(u, pageSize, currentPage, total, fruitList)
}

const detailInfo = (item) => {
  showDetail(item, detailDialogV, fruit)
}

const deleteLink = (item) => {
  deleteFruit(item)
}

const pageSizeChange = () => {
  updateList(searchFruit)
}

const currentPageChange = () => {
  updateList(searchFruit)
}

const updateFruitHandle = () => {
  updateFruit()
  updateList(searchFruit)
}

onMounted(async () => {
  // 从服务器端获取
  try {
    let res = await request.get(FruitList, {
      params: {
        pageSize: pageSize.value,
        pageNum: currentPage.value
      }
    })
    if (res.code === 200) {
      let v = res.data
      total.value = v.total
      fruitList.value = v.fruits
    } else {
      console.log(res.msg)
      ElMessage({message: '参数错误', type: 'error'})
    }
  } catch (err) {
    console.error(err)
    ElMessage({message: '服务端错误', type: 'error'})
  }

})

const addFruitDialogV = ref(false)

const addFruit = () => {
  console.log('添加水果')
  addFruitDialogV.value = true
}

const closeAddDialog = () => {
  addFruitDialogV.value= false
  updateList(searchFruit)
}

const cancelAddFruit = () => {
  addFruitDialogV.value= false
}

</script>

<style lang="scss" scoped>
@import "@/scss/home/main/search-button.scss";
</style>