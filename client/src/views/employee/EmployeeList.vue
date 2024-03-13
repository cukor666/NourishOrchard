<template>
  <div class="container">
    <el-button class="search-button" @click="changeSearchDialog">
      <el-icon>
        <Search/>
      </el-icon>
      <span style="margin-left: 5px">搜索</span>
    </el-button>
  </div>
</template>

<script setup>

import {Search} from "@element-plus/icons-vue";

const changeSearchDialog = () => {

}

</script>

<style scoped lang="scss">
.search-button {
  display: flex;
  align-items: center;
  width: 60px;
  transition: all 1s;
}

.search-button:hover {
  background: #2394f0;
  color: #eef4f8;
  box-shadow: 0 0 10px #2394f0;
  border: none;
  width: 200px;
  font-size: 20px;
}
</style>