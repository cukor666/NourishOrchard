import { defineStore } from "pinia";
import { ref } from "vue";

export const useUserInfoStore = defineStore("userInfoStore", () => {
  const promise = ref("");
  function setPromise(newPromise) {
    promise.value = newPromise;
  }
  return {
    promise,
    setPromise,
  };
});
