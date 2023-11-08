import { ref } from "vue";
import { defineStore } from "pinia";

export const useAsideStore = defineStore("aside", () => {
  const isCollapse = ref(false);
  // 侧边栏宽度
  const asideWidth = ref(200);
  return { isCollapse, asideWidth };
});
