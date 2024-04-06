import {ref} from "vue";

export function useDetail() {
    const detailDialogV = ref(false)
    const fruit = ref({
        fId: 0,
        fName: '',
        fWater: 0,
        fSugar: 0,
        fShelfLife: 0,
        fOrigin: '',
        sId: 0,
        sName: '',
        sAddress: '',
        sContactPerson: '',
        sPhone: '',
        sReputation: 0
    })

    const closeDetailDialog = () => {
        detailDialogV.value = false
    }
    const updateFruit = () => {
        console.log('执行了更新水果的函数')
        detailDialogV.value = false
    }

    return {
        detailDialogV, fruit, closeDetailDialog, updateFruit
    }
}