<template>
    <el-button @click="drawer = true">
        <p style="color: rgb(167, 35, 199);">使用说明</p>
    </el-button>
    <el-drawer v-model="drawer" title="查询用户帮助教程" :with-header="false">
        <h3>使用说明</h3>
        <br>
        <h4>ID</h4>
        <p>ID &lt; 0：表示查询所有用户</p>
        <p>ID = 0：表示不使用ID字段</p>
        <p>ID &gt; 0：表示使用ID字段，此时无论其他字段是否填写都只使用ID字段</p>
        <br>
        <h4>其他字段</h4>
        <p>仅在ID为0时才有效，且字段不填写则不使用该字段查询，填写了就表示使用该字段查询</p>
        <br>
        <h4>搜索与清空</h4>
        <p>选择好使用的字段之后，点击右边的搜索按钮即可开始查询，点击清空按钮即可将输入框还原。</p>
        <br>
        <router-link to="/see-use-info">查看教程</router-link>
    </el-drawer>
</template>

<script setup>
import { ref } from 'vue'

// 使用教程
const drawer = ref(false)
</script>

<style scoped>
.el-button {
    margin-left: 0;
    margin-right: 20px;
}
</style>
