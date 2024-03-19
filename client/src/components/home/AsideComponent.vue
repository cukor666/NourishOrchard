<template>
  <div class="container">
    <div class="logo">
      <div class="img" @click="changeAside"></div>
      <div class="title" v-if="showItemContent">滋养果园</div>
    </div>
    <el-menu class="menu" :default-active="asideStore.defaultActive" :collapse="isCollapse" style="border: 0;">
      <el-menu-item index="Home" @click="goHome">
        <template #title>
          <el-icon>
            <House/>
          </el-icon>
          <span v-if="showItemContent">首页</span>
        </template>
        <el-icon v-if="!showItemContent">
          <House/>
        </el-icon>
      </el-menu-item>
      <el-menu-item index="About" @click="gotoAbout">
        <template #title>
          <el-icon>
            <StarFilled/>
          </el-icon>
          <span v-if="showItemContent">关于我</span>
        </template>
        <el-icon v-if="!showItemContent">
          <StarFilled/>
        </el-icon>
      </el-menu-item>
      <el-sub-menu index="User" v-if="promise === 'admin'">
        <template #title>
          <el-icon>
            <User/>
          </el-icon>
          <span v-if="showItemContent">人员管理</span>
        </template>
        <el-menu-item index="UserList" @click="gotoUserList">
          <template #title>
            <span>用户列表</span>
          </template>
        </el-menu-item>
        <el-menu-item index="AdminList" @click="gotoAdminList">
          <template #title>
            <span>管理员列表</span>
          </template>
        </el-menu-item>
        <el-menu-item index="EmployeeList" @click="gotoEmployeeList">
          <template #title>
            <span>员工列表</span>
          </template>
        </el-menu-item>
        <el-menu-item index="LogoutUserList" @click="gotoLogoutUserList">
          <template #title>
            <span>注销用户列表</span>
          </template>
        </el-menu-item>
      </el-sub-menu>
      <el-sub-menu index="Fruit">
        <template #title>
          <el-icon>
            <Grape/>
          </el-icon>
          <span v-if="showItemContent">水果管理</span>
        </template>
        <el-menu-item index="FruitList" @click="gotoFruitList">
          <template #title>
            <span>水果列表</span>
          </template>
        </el-menu-item>
      </el-sub-menu>
      <el-sub-menu index="Inventory">
        <template #title>
          <el-icon>
            <OfficeBuilding/>
          </el-icon>
          <span v-if="showItemContent">仓库管理</span>
        </template>
        <el-menu-item index="WarehouseList" @click="gotoWarehouseList">
          <template #title>
            <span>仓库列表</span>
          </template>
        </el-menu-item>
        <el-menu-item index="InventoryList" @click="gotoInventoryList">
          <template #title>
            <span>库存列表</span>
          </template>
        </el-menu-item>
      </el-sub-menu>
      <el-sub-menu index="Suppliers">
        <template #title>
          <el-icon>
            <Ship/>
          </el-icon>
          <span v-if="showItemContent">供应商管理</span>
        </template>
        <el-menu-item index="SuppliersList" @click="gotoSuppliersList">
          <template #title>
            <span>供应商列表</span>
          </template>
        </el-menu-item>
      </el-sub-menu>
      <el-sub-menu index="Purchase">
        <template #title>
          <el-icon>
            <Van/>
          </el-icon>
          <span v-if="showItemContent">采购管理</span>
        </template>
        <el-menu-item index="PurchaseList" @click="gotoPurchaseList">
          <template #title>
            <span>采购列表</span>
          </template>
        </el-menu-item>
      </el-sub-menu>
      <el-sub-menu index="Orders">
        <template #title>
          <el-icon>
            <ShoppingCartFull/>
          </el-icon>
          <span v-if="showItemContent">订单管理</span>
        </template>
        <el-menu-item index="PurchaseList" @click="gotoOrdersList">
          <template #title>
            <span>订单列表</span>
          </template>
        </el-menu-item>
      </el-sub-menu>
    </el-menu>
  </div>
</template>

<script setup>
import {computed, ref} from "vue";
import {useAsideStore} from '@/stores/aside'
import {useHome} from "@/hooks/aside/useHome.js";
import {useUser} from "@/hooks/aside/useUser.js";
import {useAdmin} from "@/hooks/aside/useAdmin.js";
import {useFruit} from "@/hooks/aside/useFruit.js";
import {useLogoutUser} from "@/hooks/aside/useLogoutUser.js";
import {useEmployee} from "@/hooks/aside/useEmployee.js";
import {useLocalKey} from "@/hooks/common/useLocalKey.js";
import {useSessionKey} from "@/hooks/common/useSessionKey.js";
import {useInventory} from "@/hooks/aside/useInventory.js";
import {useWarehouse} from "@/hooks/aside/useWarehouse.js";
import {useSuppliers} from "@/hooks/aside/useSuppliers.js";
import {usePurchase} from "@/hooks/aside/usePurchase.js";
import {useOrders} from "@/hooks/aside/useOrders.js";
import {useAbout} from "@/hooks/aside/useAbout.js";

const {goHome} = useHome()
const {gotoUserList} = useUser()
const {gotoAdminList} = useAdmin()
const {gotoLogoutUserList} = useLogoutUser()
const {gotoEmployeeList} = useEmployee()
const {gotoFruitList} = useFruit()
const {gotoWarehouseList} = useWarehouse()
const {gotoInventoryList} = useInventory()
const {gotoSuppliersList} = useSuppliers()
const {gotoPurchaseList} = usePurchase()
const {gotoOrdersList} = useOrders()
const {gotoAbout} = useAbout()

const asideStore = useAsideStore()

const {NourishPromise} = useLocalKey()
const sessionKey = useSessionKey()

const promise = computed(() => {
  return sessionStorage.getItem(sessionKey.NourishPromise) || localStorage.getItem(NourishPromise)
})

const showItemContent = computed(() => {
  return asideStore.asideWidth === 200
})

const isCollapse = ref(false)

const changeAside = () => {
  asideStore.changeAsideWidth()
  isCollapse.value = !isCollapse.value
}

</script>

<style lang="scss" scoped>
@import "@/scss/home/aside/aside.scss";
</style>