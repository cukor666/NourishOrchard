import {ref} from "vue";

export function useDetail() {
    const detailDialogV = ref(false)
    const warehouse = ref({
        id: 0,
        address: '',
        capacity: 0,
        remaining: 0
    })

    const closeDetailDialog = () => {
        detailDialogV.value = false
    }
    const updateWarehouse = () => {
        console.log('执行了更新仓库的函数')
    }

    return {
        detailDialogV, warehouse, closeDetailDialog, updateWarehouse
    }
}