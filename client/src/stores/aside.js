import { defineStore } from "pinia";
import { ref } from "vue";

export const useAsideStore = defineStore("asideStore", () => {
  const maxWidth = 200;
  const minWidth = 64;
  const asideWidth = ref(maxWidth);
  function changeAsideWidth() {
    asideWidth.value = asideWidth.value === maxWidth ? minWidth : maxWidth;
  }

  const defaultActive = ref('Home')

  return {
    asideWidth,
    changeAsideWidth,
    defaultActive
  };
});
