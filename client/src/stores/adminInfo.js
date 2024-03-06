import { defineStore } from "pinia";
import { ref } from "vue";

export const useAdminInfoStore = defineStore("adminInfoStore", () => {
  const admin = ref({
    username: "",
    name: "",
    email: "",
  });
  const setAdmin = (newAdmin) => {
    admin.value = newAdmin;
  };
  return {
    admin,
    setAdmin,
  };
});
