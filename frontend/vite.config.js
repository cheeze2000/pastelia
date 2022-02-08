import { defineConfig } from "vite";
import { resolve } from "path";

import Pages from "vite-plugin-pages";
import Vue from "@vitejs/plugin-vue";

export default defineConfig({
	plugins: [
		Vue(),
		Pages()
	],
	resolve: {
		alias: {
			"~": resolve("src")
		}
	}
});
