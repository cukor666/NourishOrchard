import {ref} from "vue";

export function useInventory() {
    const inventory = ref({
        id: 0,
        commodityId: 0,
        quantity: 0,
        employeeId: 0,
        warehouseId: 0,
        createdAt: "",
        updatedAt: ""
    })

    return {
        inventory
    }
}