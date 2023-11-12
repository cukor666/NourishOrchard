<template>
    <div class="add-admin">
        <div>
            <el-card class="form-card">
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
                        <el-button type="warning" @click="onOverride">重新填写</el-button>
                    </el-form-item>
                </el-form>
            </el-card>
            <el-card>
                <h3>选中的用户：</h3>
                <ul style="margin-top: 5px; margin-bottom: 5px;">
                    <li v-for="(user, index) in selectUsers" :key="index">
                        用户名：{{ user.name }} 联系电话：{{ user.phone }}
                    </li>
                </ul>
                <el-button @click="addToAdminList">添加至管理员列表</el-button>
            </el-card>
        </div>
        <el-card class="table-card">
            <el-table :data="tableData">
                <el-table-column label="选择">
                    <template #default="{ row }">
                        <el-checkbox v-model="row.checked" @change="handleCheckboxChange(row)"></el-checkbox>
                    </template>
                </el-table-column>
                <el-table-column prop="ID" label="ID" width="50" />
                <el-table-column prop="name" label="用户名" width="100" />
                <el-table-column prop="gender" label="性别" width="80" />
                <el-table-column prop="phone" label="联系电话" width="120" />
                <el-table-column prop="address" label="家庭地址" width="190" />
            </el-table>

            <!-- 分页 -->
            <div style="margin-top: 10px; margin-left: 5px;">
                <el-pagination v-model:current-page="currentPage" v-model:page-size="pageSize" :page-sizes="[4, 8]" small
                    background layout="total, sizes, prev, pager, next, jumper" :total="total"
                    @size-change="handleSizeChange" @current-change="handleCurrentChange" />
            </div>
        </el-card>

    </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import request from '../../request';
import { ElMessage } from 'element-plus';

// 表单数据
const form = reactive({
    ID: -1,
    name: '',
    gender: '全部',
    phone: '',
    address: ''
})

// 分页
// 当前页
const currentPage = ref(1)

// 页面大小 
const pageSize = ref(8)

// 总条数
const total = ref(0)

// 表格数据
const tableData = ref([])

// 当第一次访问时就把表格初始化好
onMounted(() => {
    try {
        request.get('/user/list', {
            params: {
                currentPage: 1,
                pageSize: 8
            }
        }).then(response => {
            tableData.value = response.data.users
            total.value = response.data.rows
        })
    } catch (error) {
        console.log(error);
    }
})

// 被选择的用户，即将晋升管理员
const selectUsers = ref([])

// 复选框改变时触发
const handleCheckboxChange = (row) => {
    console.log('行选中状态变化：', row.checked);
    if (row.checked) {
        // console.log('选中的行数据：', row);
        // 如果复选框被选中，将对应的行数据添加到selectUsers数组中
        selectUsers.value.push(row);
        // console.log(selectUsers);
    } else {
        // 如果复选框被取消选中，将对应的行数据从selectUsers数组中移除
        const index = selectUsers.value.findIndex(user => user === row);
        if (index !== -1) {
            selectUsers.value.splice(index, 1);
        }
    }
};

// 添加至管理员列表
function addToAdminList() {
    // alert('添加至管理员列表')
    console.log(selectUsers.value);
    request.put('/admin/adds', selectUsers.value).then(response => {
        console.log(response);
        ElMessage({
            message: '添加成功',
            type: 'success',
        })
        selectUsers.value = []
    }).catch(err => {
        console.log(err);
    })
}

// 重新填写表单
function onOverride() {
    form.ID = -1
    form.name = ''
    form.phone = ''
    form.address = ''
    form.gender = '全部'
    ElMessage({
        message: '清空输入框成功',
        type: 'success',
    })
}

function onSubmit() {
    if (form.ID < 0) {
        try {
            request.get('/user/list', {
                params: {
                    currentPage: currentPage.value,
                    pageSize: pageSize.value
                }
            }).then(response => {
                tableData.value = response.data.users
                total.value = response.data.rows
            })
        } catch (error) {
            console.log(error);
        }
    } else if (form.ID == 0) {
        request.get('/user/struct', {
            params: {
                name: form.name,
                phone: form.phone,
                address: form.address,
                gender: form.gender == '全部' ? '' : form.gender,
                birthday: form.birthday,     // 这里传的是数组
                currentPage: currentPage.value,
                pageSize: pageSize.value
            }
        }).then(response => {
            console.log(response.data);
            tableData.value = response.data.users
            total.value = response.data.rows
        }).catch(err => {
            console.log(err);
        })
    } else {
        request.get('/user/' + form.ID).then(response => {
            tableData.value = []
            console.log(response.data);
            tableData.value.push(response.data)
            total.value = 1
        }).catch(err => {
            console.log(err);
        })
    }
}

// 改变每页大小 pageSize
function handleSizeChange() {
    // alert('handleSizeChange')
    ElMessage({
        message: '切换成功',
        type: 'success',
    })
}

// 改变当前页
function handleCurrentChange() {
    onSubmit()
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

.table-card {
    width: 650px;
}

.form-card {
    height: 320px;
}
</style>