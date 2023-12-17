<template>
    <!-- 搜索选项 -->
    <div class="search-from">
        <div :style="{ 'height': searchHeight + 'px' }" class="search-info" v-if="searchVisiable">
            <el-form :model="userForm" label-width="120px">
                <el-form-item style="display: flex;">
                    <el-form-item label="ID：">
                        <el-input type="number" v-model.number="userForm.ID" placeholder="请输入ID" />
                    </el-form-item>
                    <el-form-item label="用户名：">
                        <el-input v-model="userForm.name" placeholder="请输入用户名" />
                    </el-form-item>
                    <el-form-item label="昵称：">
                        <el-input v-model="userForm.nickname" placeholder="请输入昵称" />
                    </el-form-item>
                </el-form-item>
                <el-form-item style="display: flex;">
                    <el-form-item label="电话：">
                        <el-input v-model="userForm.phone" placeholder="请输入电话" />
                    </el-form-item>
                    <el-form-item label="地址：">
                        <el-input v-model="userForm.address" placeholder="请输入地址" />
                    </el-form-item>
                    <el-form-item label="性别：">
                        <el-select v-model="userForm.gender" placeholder="不填">
                            <el-option v-for="item in genders" :key="item.value" :label="item.label" :value="item.value" />
                        </el-select>
                    </el-form-item>
                </el-form-item>
                <el-form-item style="display: flex;">
                    <el-form-item label="生日：">
                        <el-date-picker v-model="userForm.birthday" type="daterange" range-separator="To"
                            start-placeholder="Start date" end-placeholder="End date" />
                    </el-form-item>

                    <el-button style="margin-left: 10px;" type="primary" title="搜索" @click="searchUser">
                        <el-icon>
                            <Search />
                        </el-icon>
                        搜索
                    </el-button>
                    <el-button type="warning" title="清空" @click="clearUserForm">
                        <el-icon>
                            <Refresh />
                        </el-icon>
                        清空
                    </el-button>
                </el-form-item>
            </el-form>
        </div>

        <div class="search-button">
            <el-button :title="searchVisiable ? '收起搜索' : '展开搜索'" @click="upAndDown">
                <el-icon v-if="searchVisiable">
                    <CaretTop />
                </el-icon>
                <el-icon v-else>
                    <CaretBottom />
                </el-icon>
            </el-button>
            <user-helper />
        </div>

    </div>

    <!-- 表格部分 -->
    <el-row>
        <div>
            <el-table :data="tableData">
                <el-table-column prop="ID" label="ID" width="50" />
                <el-table-column prop="name" label="用户名" width="100" />
                <el-table-column prop="nickname" label="昵称" width="180" />
                <el-table-column prop="gender" label="性别" width="80" />
                <el-table-column prop="phone" label="联系电话" width="120" />
                <el-table-column prop="address" label="家庭地址" width="190" />

                <el-table-column label="操作" width="150">
                    <template #default="{ row, index }">
                        <el-button size="small" title="详情与编辑" @click="handleEdit(index, row)"><el-icon>
                                <Edit />
                            </el-icon></el-button>
                        <el-button size="small" type="danger" title="删除" @click="handleDelete(index, row)"><el-icon>
                                <Delete />
                            </el-icon></el-button>
                    </template>
                </el-table-column>
            </el-table>
            <!-- 分页 -->
            <el-pagination v-model:current-page="currentPage" v-model:page-size="pageSize" :page-sizes="[8, 16, 32]" small
                background layout="total, sizes, prev, pager, next, jumper" :total="total" @size-change="handleSizeChange"
                @current-change="handleCurrentChange" />
        </div>
        <!-- 卡片部分 -->
        <user-card style="flex: 3;" />
    </el-row>
</template>
<script setup>
import { reactive, ref, onMounted } from 'vue';
import request from '@/request';
import router from '@/router';
import { useUserStore } from '@/stores/stores';
import { ElMessage, ElMessageBox } from 'element-plus';
import UserHelper from '@/components/user/UserHelper.vue';
import UserCard from '@/components/user/UserCard.vue';

// 用户状态
const userStore = useUserStore()

// 性别选项
const genders = ref([
    { value: "男", label: "男" },
    { value: "女", label: "女" }
])

// 用户表单
const userForm = reactive({
    ID: -1,
    name: '',
    nickname: '',
    phone: '',
    address: '',
    gender: '',
    birthday: ''    // 它是一个数组， 0 表示开始的时间，1 表示结束的时间
})

// 分页
const currentPage = ref(1)  // 当前页
const pageSize = ref(8)     // 页面大小 
const total = ref(0)        // 总条数
const tableData = ref([])   // 表格数据

// 当第一次访问时就把表格初始化好
onMounted(() => {
    request.get('/user/list', {
        params: {
            currentPage: currentPage.value,
            pageSize: pageSize.value
        }
    }).then(response => {
        tableData.value = response.data.users
        total.value = response.data.rows
    }).catch(error => console.log(error))
})

// 查询所有用户
function searchUser() {
    // 先获取一下前端填写的表单
    // console.log(userForm);
    // 获取到表单信息之后，将表单封装成json发送给后端，让后端进行查询操作
    if (userForm.ID < 0) {  // 如果userForm.ID < 0 则查询所有用户
        request.get('/user/list', {
            params: {
                currentPage: currentPage.value,
                pageSize: pageSize.value
            }
        }).then(response => {
            tableData.value = response.data.users
            total.value = response.data.rows
        }).catch(error => console.log(error))
    } else if (userForm.ID == 0) {   // 如果userForm.ID == 0 则表示不使用ID这个字段
        console.log(userForm);
        request.get('/user/struct', {
            params: {
                name: userForm.name,
                nickname: userForm.nickname,
                phone: userForm.phone,
                address: userForm.address,
                gender: userForm.gender,
                birthday: userForm.birthday,     // 这里传的是数组
                currentPage: currentPage.value,
                pageSize: pageSize.value
            }
        }).then(response => {
            console.log(response.data);
            tableData.value = response.data.users
            total.value = response.data.rows
        }).catch(err => { console.log(err); })
    } else {    // 如果userForm.ID > 0 则表示使用ID这个字段
        //因为ID是唯一的，所以。当使用了ID这个字段，那就直接忽略掉其他字段，否则没有任何意义
        request.get('/user/' + userForm.ID).then(response => {
            tableData.value = []
            console.log(response.data);
            tableData.value.push(response.data)
            total.value = 1
        }).catch(err => { console.log(err); })
    }
}

// 清空搜索输入
function clearUserForm() {
    userForm.ID = -1
    userForm.name = ''
    userForm.nickname = ''
    userForm.phone = ''
    userForm.address = ''
    userForm.gender = ''
    userForm.birthday = ''
    ElMessage({ message: '清空输入框成功', type: 'success' })
}

// 编辑用户信息
function handleEdit(index, row) {
    userStore.tempUser = row
    router.push('/user/edit-info')
}

// 删除用户，这是一个比较危险的操作，所以要弹出一个消息框提示用户
function handleDelete(index, row) {
    ElMessageBox.confirm('您确定要删除这个数据吗?', 'Warning', {
        confirmButtonText: '是的',
        cancelButtonText: '取消',
        type: 'warning',
    }).then(() => {
        userStore.tempUser = row
        request.delete('/user/delete', {
            params: {
                ID: userStore.tempUser.ID,
                promise: userStore.loginUser.promise    // 删除动作很危险，需要传入权限
            }
        }).then(response => {
            console.log(response);
            if (response.code === 200) {
                ElMessage({ message: '删除成功', type: 'success' })
                searchUser()    // 删除成功后要刷新一下用户列表，而不是让用户在自己点击一次查询按钮
            } else { ElMessage({ message: "删除失败，权限不够", type: 'error' }) }
        }).catch(error => {
            console.log(error);
            ElMessage({ type: 'error', message: '删除失败' })
        })
    }).catch(() => { ElMessage({ type: 'info', message: '取消删除' }) })
}

// 改变每页大小 pageSize
function handleSizeChange() { ElMessage({ message: '切换成功', type: 'success' }) }

// 改变当前页
function handleCurrentChange() { searchUser() }

// 搜索下拉菜单高度
const searchHeight = ref(0)
const searchVisiable = ref(false)

// 搜索下拉菜单
function upAndDown() {
    searchVisiable.value = !searchVisiable.value
    searchHeight.value = searchVisiable.value ? 150 : 0
}

</script>

<style lang="scss" scoped>
.el-button {
    padding: 0 4px;
}

.el-table {
    border-radius: 15px;
}

.el-pagination {
    margin-top: 10px;
    margin-left: 5px;
}

.search-from {
    display: flex;
    flex-direction: column;
    background-color: #F8E2DC;

    .search-info {
        margin-top: 10px;
    }
}

.search-button {
    display: flex;
    margin-top: auto;
    width: 100%;
    background-color: #EBE6E6;

    .el-button {
        width: 50px;
        margin-bottom: 5px;
        margin-left: auto;
        margin-right: 10px;
    }
}

.el-row {
    display: flex;
    margin-left: 10px;

    div {
        flex: 5;
        margin-right: 10px;
    }
}
</style>