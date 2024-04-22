import {ref} from "vue";

export function useOrder() {
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

    return {
        order
    }
}