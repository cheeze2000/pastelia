import { createApp } from "vue";

import "highlight.js/styles/atom-one-dark.css";
import "~/index.css";
import App from "~/App.vue";
import router from "~/router";

createApp(App).use(router).mount("#app");
