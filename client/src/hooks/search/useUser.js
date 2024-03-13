import {ref} from "vue";

export function useUser() {
    const user = ref({
        id: 0,
        username: "",
        name: "",
        gender: "",
        phone: "",
        address: "",
        birthday: ""
    })

    return {
        user
    }
}