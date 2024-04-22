import {ref} from "vue";

export function useDetail() {
    const detailDialogV = ref(false)
    const order = ref({
        id: 0,
        title: "",
        status: 0,
        commodityId: 0,
        quantity: 0,
        buyerId: 0,
        adminId: 0,
        warehouseId: 0,
        receiverName: "",
        receiverPhone: "",
        address: "",
        remark: "",
        createdAt: "",
        updatedAt: ""
    })

    const closeDetailDialog = () => {
        detailDialogV.value = false
    }
    const updateOrder = () => {
        console.log('执行了更新订单的函数')
    }

    return {
        detailDialogV, order, closeDetailDialog, updateOrder
    }
}