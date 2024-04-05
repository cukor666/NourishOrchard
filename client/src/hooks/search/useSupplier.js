import {ref} from "vue";

export function useSupplier() {
    const supplier = ref({
        id: 0,
        name: "",
        address: "",
        contactPerson: "",
        phone: "",
        reputation: 0
    })

    return {
        supplier
    }
}