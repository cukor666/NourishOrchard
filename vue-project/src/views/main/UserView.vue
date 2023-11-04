<template>
    <div>
        <div style="display: flex; justify-content: center; align-items: center; height: 100px;">
            <div style="flex: 6; margin-left: 10px;">
                ID：<el-input type="number" style="width: 80px;" v-model="userForm.ID"></el-input>
                用户名：<el-input placeholder="Cukor" style="width: 80px;" v-model="userForm.name"></el-input>
                电话：<el-input placeholder="12345678910" style="width: 120px;" v-model="userForm.phone"></el-input>
                地址：<el-input placeholder="地址" style="width: 200px;" v-model="userForm.address"></el-input>
                性别：<el-select v-model="userForm.gender" style="width: 80px;" placeholder="其他">
                    <el-option v-for="item in genders" :key="item.value" :label="item.label" :value="item.value" />
                </el-select>

                <div style="margin-top: 10px;">
                    <span class="demonstration">生日：</span>
                    <el-date-picker v-model="userForm.birthday" type="daterange" range-separator="To"
                        start-placeholder="Start date" end-placeholder="End date"  />
                </div>

            </div>

            <div style="flex: 1;">
                <el-button type="primary" @click="searchUser">
                    <el-icon>
                        <Search />
                    </el-icon>
                    搜索
                </el-button>
                <el-button type="warning">
                    <el-icon>
                        <Refresh />
                    </el-icon>
                    清空
                </el-button>
            </div>
        </div>
        <div>
            <el-scrollbar>
                <el-table :data="tableData">
                    <el-table-column prop="ID" label="ID" width="50" />
                    <el-table-column prop="name" label="用户名" width="100" />
                    <el-table-column prop="gender" label="性别" width="80" />
                    <el-table-column prop="phone" label="联系电话" width="120" />
                    <el-table-column prop="address" label="家庭地址" width="120" />
                    <el-table-column prop="birthday" label="生日" width="200" />
                    <el-table-column prop="CreatedAt" label="创建时间" width="200" />
                    <el-table-column prop="UpdatedAt" label="更新时间" width="200" />
                    <el-table-column label="操作" width="150">
                        <template #default="scope">
                            <!-- 下标index从0开始 -->
                            <el-button size="small" @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
                            <el-button size="small" type="danger"
                                @click="handleDelete(scope.$index, scope.row)">删除</el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </el-scrollbar>
            <div>
                <!-- :locale="locale" -->
                <el-config-provider>
                    <el-pagination :total="20" />
                </el-config-provider>
            </div>
        </div>
    </div>
</template>
<script setup>
import { ref } from 'vue'
import request from '../../request';
import router from '../../router';
import { useUserStore } from '../../stores/stores'
import { ElMessage, ElMessageBox } from 'element-plus'


const userStore = useUserStore()

const tableData = ref([])

const userForm = ref({
    ID: 1,
    name: '',
    phone: '',
    address: '',
    gender: '',
    birthday: ''
})

const genders = ref([
    {
        value: "男",
        label: "男"
    },
    {
        value: "女",
        label: "女"
    },
    {
        value: "其他",
        label: "其他"
    }
])

// 查询所有用户
function searchUser() {
    try {
        request.get('/user-list').then(response => {
            // console.log(response.data);
            tableData.value = response.data
        })
    } catch (error) {
        console.log(error);
    }
}

// 编辑用户信息
function handleEdit(index, row) {
    // console.log(index);
    // console.log(row);
    userStore.userRow = row
    // console.log("UserView: userStore.userRow ");
    // console.log(userStore.userRow);
    router.push('/edit-user-info')
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
            // /delete-user?ID=userStore.userRow.ID   65,1,2,7
            request.delete('/delete-user', {
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

</script>

<style scoped>
.el-button {
    /* margin: 0; */
    padding: 0 4px;
}
</style>