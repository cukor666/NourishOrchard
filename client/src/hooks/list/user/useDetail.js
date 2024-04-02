import {computed, ref} from "vue";
import {useSessionKey} from "@/hooks/common/useSessionKey.js";

const { NourishUserInfoUpdated } = useSessionKey()

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
        sessionStorage.setItem(NourishUserInfoUpdated, "true")
    }

    return {
        detailDialogV, user, closeDetailDialog, updateUser
    }
}