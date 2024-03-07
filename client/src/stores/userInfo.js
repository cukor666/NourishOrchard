import {defineStore} from "pinia";
import {ref, watch} from "vue";

export const useUserInfoStore = defineStore("userInfoStore", () => {
    const user = ref({
        username: "",
        name: "",
        gender: "",
        phone: "",
        address: "",
        birthday: "",
    });
    const setUser = (newUser) => {
        user.value = newUser;
    };
    const gender = ref(true);
    watch(gender, (isMan) => {
        user.value.gender = isMan ? "男" : "女";
    });

    return {
        user,
        gender,
        setUser,
    };
});
