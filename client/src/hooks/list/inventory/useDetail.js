import {ref} from "vue";

export function useDetail() {
    const detailDialogV = ref(false)
    const inventory = ref({
        id: 0,
        commodityId: 0,
        quantity: 0,
        employeeId: 0,
        warehouseId: 0,
        createdAt: '',
        updatedAt: ''
    })

    const closeDetailDialog = () => {
        detailDialogV.value = false
    }
    const updateInventory = () => {
        console.log('执行了更新库存信息的函数')
    }

    return {
        detailDialogV, inventory, closeDetailDialog, updateInventory
    }
}