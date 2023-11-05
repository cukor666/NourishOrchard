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
                        start-placeholder="Start date" end-placeholder="End date" />
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
                <el-button>
                    <el-icon>
                        <Refresh />
                    </el-icon>
                    使用说明
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
import { reactive, ref } from 'vue'
import request from '../../request';
import router from '../../router';
import { useUserStore } from '../../stores/stores'
import { ElMessage, ElMessageBox } from 'element-plus'

// 用户状态
const userStore = useUserStore()

// 表格数据
const tableData = ref([])

// 用户表单
const userForm = reactive({
    ID: -1,
    name: '',
    phone: '',
    address: '',
    gender: '',
    birthday: ''    // 它是一个数组， 0 表示开始的时间，1 表示结束的时间
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
// 优化搜索，todo
function searchUser() {
    // 先获取一下前端填写的表单
    console.log(userForm);
    // 获取到表单信息之后，将表单封装成json发送给后端，让后端进行查询操作
    if (userForm.ID < 0) {  // 如果userForm.ID < 0 则查询所有用户
        try {
            request.get('/user-list').then(response => {
                // console.log(response.data);
                tableData.value = response.data
            })
        } catch (error) {
            console.log(error);
        }
    } else if (userForm.ID == 0) {   // 如果userForm.ID == 0 则表示不使用ID这个字段
        console.log(userForm);
        request.get('/user-struct', {
            params: {
                name: userForm.name,
                phone: userForm.phone,
                address: userForm.address,
                gender: userForm.gender,
                birthday: userForm.birthday     // 这里传的是数组
            }
        }).then(response => {
            console.log(response.data);
            tableData.value = response.data
        }).catch(err => {
            console.log(err);
        })
    } else {    // 如果userForm.ID > 0 则表示使用ID这个字段
        //因为ID是唯一的，所以。当使用了ID这个字段，那就直接忽略掉其他字段，否则没有任何意义
        request.get('/user/' + userForm.ID).then(response => {
            tableData.value = []
            console.log(response.data);
            tableData.value.push(response.data)
        }).catch(err => {
            console.log(err);
        })
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