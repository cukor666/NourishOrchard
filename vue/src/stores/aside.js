import { ref } from "vue";
import { defineStore } from "pinia";

export const useAsideStore = defineStore("aside", () => {
  const isCollapse = ref(true);
  // 侧边栏宽度
  const asideWidth = ref(64);
  return { isCollapse, asideWidth };
});
