<template>
    <el-space wrap>
        <el-card>
            <!-- 地区 -->
            <template #header>
                <div class="card-header">
                    <span class="address">地区</span>
                </div>
            </template>
            <div>
                <p>写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字</p>
            </div>
        </el-card>

        <!-- 最新成员 -->
        <el-card>
            <template #header>
                <div class="card-header">
                    <span>最新成员</span>
                </div>
            </template>
            <div class="newMenber">
                <p>写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字</p>
            </div>
        </el-card>

        <!-- 最近删除 -->
        <el-card>
            <template #header>
                <div class="card-header">
                    <span>最新删除</span>
                </div>
            </template>
            <div class="newDeleted">
                <p>写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字</p>
            </div>
        </el-card>
    </el-space>
</template>

<style>
.card-header {
    text-align: center;
    width: auto;
    /* 自适应 */
}

.el-card {
    border-radius: 15px;
}

.newMenber {
    height: 250px;
}

.newDeleted {
    height: 200px;
}

.address {
    letter-spacing: 1em;
}
</style>