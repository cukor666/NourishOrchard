import { reactive, ref } from "vue";
import { defineStore } from "pinia";

export const useUserStore = defineStore("user", () => {
  const loginUser = reactive({
    name: "xxx",
    promise: 1,
  });

  const UserInfoVisible = ref(false);
  const tempUser = reactive({
    ID: 0,
    name: "",
    nickname: "xxx",
    gender: "",
    birthday: "",
    phone: "",
    address: "",
    CreatedAt: "",
    UpdatedAt: "",
    promise: 0,
  });
  return {
    loginUser,
    UserInfoVisible,
    tempUser,
  };
});
