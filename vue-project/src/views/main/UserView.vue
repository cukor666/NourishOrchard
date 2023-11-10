<template>
    <div>
        <!-- 使用说明 -->
        <user-helper style="float: right;" />

        <!-- 主体部分 -->
        <div class="main">
            <!-- 搜索选项 -->
            <div style="flex: 12; margin-left: 10px;">
                ID：<el-input type="number" style="width: 80px;" v-model="userForm.ID"></el-input>
                用户名：<el-input placeholder="Cukor" style="width: 80px;" v-model="userForm.name"></el-input>
                昵称：<el-input placeholder="Sugar" style="width: 80px;" v-model="userForm.nickname"></el-input>
                电话：<el-input placeholder="12345678910" style="width: 120px;" v-model="userForm.phone"></el-input>
                地址：<el-input placeholder="地址" style="width: 200px;" v-model="userForm.address"></el-input>
                性别：<el-select v-model="userForm.gender" style="width: 80px;" placeholder="不填">
                    <el-option v-for="item in genders" :key="item.value" :label="item.label" :value="item.value" />
                </el-select>

                <div style="margin-top: 10px;">
                    <span class="demonstration">生日：</span>
                    <el-date-picker v-model="userForm.birthday" type="daterange" range-separator="To"
                        start-placeholder="Start date" end-placeholder="End date" />
                </div>

            </div>

            <!-- 按钮组 -->
            <div style="flex: 1;">
                <el-button type="primary" circle title="搜索" @click="searchUser">
                    <el-icon>
                        <Search />
                    </el-icon>
                </el-button>
                <el-button type="warning" circle title="清空" @click="clearUserForm">
                    <el-icon>
                        <Refresh />
                    </el-icon>
                </el-button>
            </div>
        </div>

        <!-- 表格部分 -->
        <div style="margin-left: 10px;">
            <el-row style="display: flex;">
                <div style="flex: 5; margin-right: 10px;">
                    <el-table :data="tableData">
                        <el-table-column prop="ID" label="ID" width="50" />
                        <el-table-column prop="name" label="用户名" width="100" />
                        <el-table-column prop="nickname" label="昵称" width="180" />
                        <el-table-column prop="gender" label="性别" width="80" />
                        <el-table-column prop="phone" label="联系电话" width="120" />
                        <el-table-column prop="address" label="家庭地址" width="190" />

                        <el-table-column label="操作" width="150">
                            <template #default="scope">
                                <!-- 下标index从0开始 -->
                                <el-button size="small" title="详情与编辑" @click="handleEdit(scope.$index, scope.row)">
                                    <el-icon>
                                        <Edit />
                                    </el-icon>
                                </el-button>
                                <el-button size="small" type="danger" title="删除"
                                    @click="handleDelete(scope.$index, scope.row)">
                                    <el-icon>
                                        <Delete />
                                    </el-icon>
                                </el-button>
                            </template>
                        </el-table-column>
                    </el-table>
                    <!-- 分页 -->
                    <div style="margin-top: 10px; margin-left: 5px;">
                        <el-pagination v-model:current-page="currentPage" v-model:page-size="pageSize"
                            :page-sizes="[8, 16, 32]" small background layout="total, sizes, prev, pager, next, jumper"
                            :total="total" @size-change="handleSizeChange" @current-change="handleCurrentChange" />
                    </div>
                </div>
                <!-- 卡片部分 -->
                <user-card style="flex: 3;" />
            </el-row>

        </div>
    </div>
</template>
<script setup>
import { reactive, ref } from 'vue'
import request from '../../request';
import router from '../../router';
import { useUserStore } from '../../stores/stores'
import { ElMessage, ElMessageBox } from 'element-plus'
import UserHelper from '../../components/user/UserHelper.vue'
import UserCard from '../../components/user/UserCard.vue';

// 用户状态
const userStore = useUserStore()

// 表格数据
const tableData = ref([])

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

// 性别选项
const genders = ref([
    {
        value: "男",
        label: "男"
    },
    {
        value: "女",
        label: "女"
    }
])

// 查询所有用户
function searchUser() {
    // 先获取一下前端填写的表单
    // console.log(userForm);
    // 获取到表单信息之后，将表单封装成json发送给后端，让后端进行查询操作
    if (userForm.ID < 0) {  // 如果userForm.ID < 0 则查询所有用户
        // console.log(pageSize.value);
        // console.log(currentPage.value);
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
        }).catch(err => {
            console.log(err);
        })
    } else {    // 如果userForm.ID > 0 则表示使用ID这个字段
        //因为ID是唯一的，所以。当使用了ID这个字段，那就直接忽略掉其他字段，否则没有任何意义
        request.get('/user/' + userForm.ID).then(response => {
            tableData.value = []
            console.log(response.data);
            tableData.value.push(response.data)
            total.value = 1
        }).catch(err => {
            console.log(err);
        })
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
    ElMessage({
        message: '清空输入框成功',
        type: 'success',
    })
}

// 编辑用户信息
function handleEdit(index, row) {
    userStore.userRow = row
    // console.log(row);
    router.push('/user/edit-info')
}

// 删除用户，这是一个比较危险的操作，所以要弹出一个消息框提示用户
function handleDelete(index, row) {
    ElMessageBox.confirm(
        '您确定要删除这个数据吗?',
        'Warning',
        {
            confirmButtonText: '是的',
            cancelButtonText: '取消',
            type: 'warning',
        }
    ).then(() => {
        userStore.userRow = row
        // console.log("UserView: userStore.userRow ");
        // console.log(userStore.userRow);
        try {
            request.delete('/user/delete', {
                params: {
                    ID: userStore.userRow.ID
                }
            }).then(response => {
                ElMessage({
                    message: '删除成功',
                    type: 'success',
                })
                console.log(response);
                searchUser()    // 删除成功后要刷新一下用户列表，而不是让用户在自己点击一次查询按钮
            })
        } catch (error) {
            console.log(error);
            ElMessage({
                type: 'error',
                message: '删除失败',
            })
        }
    }).catch(() => {
        ElMessage({
            type: 'info',
            message: '取消删除',
        })
    })
}

// 分页
// 当前页
const currentPage = ref(1)

// 页面大小 
const pageSize = ref(8)

// 总条数
const total = ref(0)

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
    searchUser()
}
</script>

<style scoped>
.el-button {
    /* margin: 0; */
    padding: 0 4px;
}

.el-table {
    border-radius: 15px;
}

.main {
    display: flex;
    justify-content: center;
    align-items: center;
    height: auto;
    margin-top: 5px;
    margin-bottom: 5px;
}
</style>