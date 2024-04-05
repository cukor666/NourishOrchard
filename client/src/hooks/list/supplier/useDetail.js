import {ref} from "vue";

export function useDetail() {
    const detailDialogV = ref(false)
    const supplier = ref({
        id: 0,
        name: "",
        address: "",
        contactPerson: "",
        phone: "",
        reputation: 0
    })

    const closeDetailDialog = () => {
        detailDialogV.value = false
    }
    const updateSupplier = () => {
        console.log('执行了更新供应商的函数')
    }

    return {
        detailDialogV, supplier, closeDetailDialog, updateSupplier
    }
}