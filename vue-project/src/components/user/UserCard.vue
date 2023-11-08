<template>
    <div>
        <el-space wrap>
            <el-card>
                <template #header>
                    <div class="card-header">
                        <span class="address">地区</span>
                        <p>写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字</p>
                    </div>
                </template>
            </el-card>

            <!-- 最新成员 -->
            <el-card>
                <template #header>
                    <div class="card-header newMenber">
                        <span>最新成员</span>
                        <p>写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字</p>
                    </div>
                </template>
            </el-card>

            <el-card>
                <template #header>
                    <div class="card-header newDeleted">
                        <span>最新删除</span>
                        <p>写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字写了一大段文字</p>

                    </div>
                </template>
            </el-card>

        </el-space>
    </div>
</template>

<style>
.card-header {
    text-align: center;
    width: auto;
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