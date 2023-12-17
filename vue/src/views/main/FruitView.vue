<template>
    <div>
        <!-- 主体部分 -->
        <div class="main">
            <!-- 搜索选项 -->
            <div style="flex: 12; margin-left: 10px;">
                ID：<el-input type="number" style="width: 80px;" v-model="fruitForm.ID"></el-input>
                水果名：<el-input placeholder="苹果" style="width: 120px;" v-model="fruitForm.name"></el-input>
                含水量：<el-select v-model="fruitForm.water" style="width: 90px;" placeholder="不填">
                    <el-option v-for="item in water" :key="item.value" :label="item.label" :value="item.value" />
                </el-select>
                含糖量：<el-select v-model="fruitForm.sugar" style="width: 90px;" placeholder="不填">
                    <el-option v-for="item in sugar" :key="item.value" :label="item.label" :value="item.value" />
                </el-select>
                库存：<el-select v-model="fruitForm.inventory" style="width: 90px;" placeholder="不填">
                    <el-option v-for="item in inventory" :key="item.value" :label="item.label" :value="item.value" />
                </el-select>
            </div>

            <!-- 按钮组 -->
            <div style="flex: 1;">
                <el-button type="primary" circle title="搜索" @click="searchFruit">
                    <el-icon>
                        <Search />
                    </el-icon>
                </el-button>
                <el-button type="warning" circle title="清空" @click="clearFruitForm">
                    <el-icon>
                        <Refresh />
                    </el-icon>
                </el-button>
            </div>
        </div>

        <div style="margin-left: 10px;">
            <el-row style="display: flex;">
                <div style="flex: 5; margin-right: 10px;">
                    <el-table :data="tableData" style="border-radius: 15px;">
                        <el-table-column prop="ID" label="ID" width="50" />
                        <el-table-column prop="fruitname" label="水果名" width="100" />
                        <el-table-column prop="water" label="含水量" width="100" />
                        <el-table-column prop="sugar" label="含糖量" width="100" />
                        <el-table-column prop="inventory" label="库存" width="100" />
                        <el-table-column prop="price" label="单价" width="100" />
                        <el-table-column prop="suppherid" label="供应商" width="100" />


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
                <!-- <user-card style="flex: 3;" /> -->
            </el-row>
        </div>
    </div>
</template>

<script setup>
import { reactive, ref } from 'vue'
import { ElMessage } from 'element-plus'

// 水果搜索信息表单
const fruitForm = reactive({
    ID: -1,
    name: '',
    water: '',
    sugar: '',
    inventory: '',
    price: '',
    suppherid: ''
})

// 含水量选项
const water = ref([
    { value: 10, label: "≤10%" },
    { value: 20, label: "≤20%" },
    { value: 30, label: "≤30%" },
    { value: 40, label: "≤40%" },
    { value: 50, label: "≤50%" },
    { value: 60, label: "≤60%" },
    { value: 70, label: "≤70%" },
    { value: 80, label: "≤80%" },
    { value: 90, label: "≤90%" },
    { value: 100, label: "≤100%" },
])

// 含糖量选项
const sugar = ref([
    { value: 10, label: "≤10%" },
    { value: 20, label: "≤20%" },
    { value: 30, label: "≤30%" },
    { value: 40, label: "≤40%" },
    { value: 50, label: "≤50%" },
    { value: 60, label: "≤60%" },
    { value: 70, label: "≤70%" },
    { value: 80, label: "≤80%" },
    { value: 90, label: "≤90%" },
    { value: 100, label: "≤100%" },
])

// 库存选项
const inventory = ref([
    { value: 100, label: "≤100" },
    { value: 200, label: "≤200" },
    { value: 300, label: "≤300" },
    { value: 400, label: "≤400" },
    { value: 500, label: "≤500" },
    { value: 600, label: "≤600" },
    { value: 700, label: "≤700" },
    { value: 800, label: "≤800" },
    { value: 900, label: "≤900" },
    { value: 1000, label: "≤1000" },
])

// 表格数据
const tableData = ref([])

function searchFruit() {
    alert('搜索水果')
}

function clearFruitForm() {
    fruitForm.ID = -1
    fruitForm.name = ''
    fruitForm.water = ''
    fruitForm.sugar = ''
    fruitForm.inventory = ''
    fruitForm.price = ''
    fruitForm.suppherid = ''
    ElMessage({ message: '清空输入框成功', type: 'success' })
}


// 编辑水果信息
function handleEdit(index, row) {
    alert('编辑水果信息')
    // userStore.tempUser = row
    // router.push('/user/edit-info')
}

// 删除水果信息
function handleDelete(index, row) {
    alert('删除水果信息')
}

// 分页
const currentPage = ref(1)  // 当前页
const pageSize = ref(8)     // 页面大小 
const total = ref(0)        // 总条数

// 改变每页大小 pageSize
function handleSizeChange() { ElMessage({ message: '切换成功', type: 'success' }) }

// 改变当前页
function handleCurrentChange() { alert('切换当前页') }

</script>

<style scoped>
.main {
    display: flex;
}
</style>