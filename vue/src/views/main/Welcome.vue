<template>
    <div>
        欢迎光临
    </div>
</template>

<script setup>
import { onMounted } from 'vue'
import { ElMessage } from 'element-plus';

onMounted(() => {
    ElMessage({
        message: "欢迎光临",
        type: "success",
    });
})
</script>