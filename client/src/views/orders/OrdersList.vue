<template>
  <div class="container">
    <el-button class="search-button" @click="changeSearchDialog">
      <el-icon>
        <Search/>
      </el-icon>
      <span style="margin-left: 5px">搜索</span>
    </el-button>

    <search-order :searchDialogV="searchDialogV" @closeDialog="closeSearchDialog" @search="search"></search-order>

    <el-card style="width: 99%">
      <el-table :data="orderList" stripe style="width: 100%">
        <el-table-column prop="id" label="id" width="50"/>
        <el-table-column prop="title" label="订单标题" width="150"/>
        <el-table-column prop="status" label="订单状态" width="100"/>
        <el-table-column prop="commodityId" label="商品ID" width="100"/>
        <el-table-column prop="buyerId" label="用户ID" width="100"/>
        <el-table-column prop="adminId" label="管理员ID" width="100"/>
        <el-table-column prop="warehouseId" label="仓库ID" width="100"/>
        <el-table-column prop="receiverName" label="收货人名" width="100"/>
        <el-table-column prop="receiverPhone" label="收货人电话" width="100"/>
        <el-table-column prop="address" label="收货地址" width="100"/>
        <el-table-column prop="remark" label="备注" width="100"/>
        <el-table-column prop="createdAt" label="创建时间" width="150"/>
        <el-table-column prop="updatedAt" label="更新时间" width="150"/>

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

    <!--订单详情对话框-->
    <detail :detail-dialog-v="detailDialogV" :order="order"  @closeDialog="closeDetailDialog"
            @updateOrder="updateOrder"></detail>
  </div>
</template>

<script setup>
import SearchOrder from "@/components/order/SearchOrder.vue";
import Detail from "@/components/order/dialog/Detail.vue";
import {Search} from "@element-plus/icons-vue";
import {onMounted, ref} from "vue";
import request from "@/axios/request.js";
import {ElMessage} from "element-plus";
import {useSearch} from "@/hooks/list/order/useSearch.js"
import {useDetail} from "@/hooks/list/order/useDetail.js";
import {usePage} from "@/hooks/list/usePage.js";
import {useTable} from "@/hooks/list/order/useTable.js";
import {OrderList} from "@/api/order/order-api.js";

const {searchDialogV, searchOrder, changeSearchDialog, closeSearchDialog, findOrder} = useSearch()
const orderList = ref([{...searchOrder.value}])
const {detailDialogV, order, closeDetailDialog, updateOrder} = useDetail()
const {currentPage, pageSize, pageSizes, total} = usePage()
const {updateList, showDetail, deleteOrder} = useTable(orderList, pageSize, currentPage, total)

const pageSizeChange = () => {
  updateList(searchOrder)
}

const currentPageChange = () => {
  updateList(searchOrder)
}

const search = (u) => {
  findOrder(u, pageSize, currentPage, total, orderList)
}

const detailInfo = (item) => {
  showDetail(item, detailDialogV, order)
}

const deleteLink = (item) => {
  deleteOrder(item)
}

onMounted(async () => {
  // 从服务器端获取
  try {
    let res = await request.get(OrderList, {
      params: {
        pageSize: pageSize.value,
        pageNum: currentPage.value
      }
    })
    if (res.code === 200) {
      let v = res.data
      total.value = v.total
      orderList.value = v.orders
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