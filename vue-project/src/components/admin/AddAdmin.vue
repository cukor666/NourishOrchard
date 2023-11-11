<template>
    <div class="add-admin">
        <el-card>
            <el-form :model="form" label-width="80px">
                <el-form-item label="ID">
                    <el-input v-model="form.ID" />
                </el-form-item>
                <el-form-item label="用户名">
                    <el-input v-model="form.name" />
                </el-form-item>
                <el-form-item label="性别">
                    <el-radio-group v-model="form.gender">
                        <el-radio label="男">男</el-radio>
                        <el-radio label="女">女</el-radio>
                        <el-radio label="全部">全部</el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item label="联系电话">
                    <el-input v-model="form.phone" />
                </el-form-item>
                <el-form-item label="家庭地址">
                    <el-input v-model="form.address" />
                </el-form-item>

                <el-form-item>
                    <el-button type="primary" @click="onSubmit">搜索</el-button>
                    <el-button type="warning">重新填写</el-button>
                </el-form-item>
            </el-form>
        </el-card>

        <el-card>
            <p>afgag</p>
        </el-card>
    </div>
</template>

<script setup>
import { reactive } from 'vue'


const form = reactive({
    ID: 0,
    name: '',
    gender: '全部',
    phone: '',
    address: ''
})


function onSubmit() {
    alert('搜索')
}

</script>

<style scoped>
.add-admin {
    display: flex;
    justify-content: center;
}

.el-card {
    width: 450px;
    margin: 10px 5px;
    border-radius: 15px;
}
</style>