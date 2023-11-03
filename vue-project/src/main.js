import { createApp } from "vue";
import { createPinia } from "pinia";
import ElementPlus from "element-plus";
import "element-plus/dist/index.css";
import App from "./App.vue";
import router from "./router";
import * as ElementPlusIconsVue from "@element-plus/icons-vue";
import axios from "axios";
import VueAxios from 'vue-axios'

const app = createApp(App);

app.use(createPinia());
app.use(ElementPlus);
app.use(router);
app.use(VueAxios, axios);

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component);
}

app.mount("#app");
