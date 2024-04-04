import {ref} from "vue";

export function useAdmin() {
    const admin = ref({
        id: 0,
        username: "",
        name: "",
        email: ""
    })

    return {
        admin
    }
}