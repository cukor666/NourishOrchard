import {computed, ref} from "vue";

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
    const userInfoUpdated = computed(() => {
        return sessionStorage.getItem('nourish-user-info-updated') || "false"
    })
    const closeDetailDialog = () => {
        detailDialogV.value = false
    }
    const updateUser = () => {
        console.log('执行了更新用户的函数')
        sessionStorage.setItem('nourish-user-info-updated', "true")
    }

    return {
        detailDialogV, user, userInfoUpdated, closeDetailDialog, updateUser
    }
}