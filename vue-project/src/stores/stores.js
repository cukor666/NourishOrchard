import { ref } from "vue";
import { defineStore } from "pinia";

export const useUserStore = defineStore("user", () => {
  const loginUserName = ref("xxx");
  const loginUserPromise = ref(1);
  const UserInfoVisible = ref(false);
  const token = ref("");
  const userGridData = ref([
    {
      ID: 0,
      name: "xxx",
      nickname: "xxx",
      gender: "xxx",
      birthday: "xxx",
      phone: "xxx",
      address: "xxx",
      promise: "xxx",
    },
  ]);
  const userRow = ref({
    ID: 0,
    name: "",
    nickname: "xxx",
    gender: "",
    birthday: "",
    phone: "",
    address: "",
    CreatedAt: "",
    UpdatedAt: "",
    promise: "xxx",
  });
  return {
    loginUserName,
    loginUserPromise,
    UserInfoVisible,
    userGridData,
    userRow,
    token,
  };
});
