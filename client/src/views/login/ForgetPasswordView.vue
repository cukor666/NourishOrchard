<template>
    <div class="container">
        <div style="width: 200px;">
            <el-input v-model="user.username" @blur="validName(user.username)" />
            <el-input v-model="user.age" />
        </div>
        <el-button @click="changeUser">更改</el-button>
    </div>
</template>

<script setup>
import { reactive } from "vue";


const user = reactive({
    username: 'Cukor',
    age: 18
})

const validName = (name) => {
    console.log(name);
}

const changeUser = () => {
    console.log(user);
}

</script>

<style lang="scss" scoped></style>
