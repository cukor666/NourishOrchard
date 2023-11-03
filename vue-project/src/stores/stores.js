import { ref } from "vue";
import { defineStore } from "pinia";

export const useUserStore = defineStore("user", () => {
  const loginUserName = ref('xxx');
  const UserInfoVisible = ref(false);
  const token = ref('')
  const userGridData = ref([
    {
      ID: 0,
      name: "xxx",
      gender: "xxx",
      birthday: "xxx",
      phone: "xxx",
      address: "xxx",
    },
  ]);
  const userRow = ref({
    ID: 0,
    name: "",
    gender: "",
    birthday: "",
    phone: "",
    address: "",
    CreatedAt: "",
    UpdatedAt: "",
  });
  return { loginUserName, UserInfoVisible, userGridData, userRow, token };
});
