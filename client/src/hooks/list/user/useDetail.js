import {ref} from "vue";

export function useDetail() {
    const detailDialogV = ref(false)
    const user = ref({
        id: 6,
        username: "CukorZhong",
        name: "",
        gender: "",
        phone: "",
        address: "",
        birthday: ""
    })

    const closeDetailDialog = () => {
        detailDialogV.value = false
    }
    const updateUser = () => {
        console.log('执行了更新用户的函数')
    }

    return {
        detailDialogV, user, closeDetailDialog, updateUser
    }
}