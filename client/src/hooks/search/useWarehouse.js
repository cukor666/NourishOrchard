import {ref} from "vue";

export function useWarehouse() {
    const warehouse = ref({
        id: 0,
        address: "",
        capacity: 0,
        remaining: 0,
    })

    return {
        warehouse
    }
}