<script setup>
import { onMounted, reactive } from "vue";
import { useRouter } from "vue-router";

import About from "~/components/misc/About.vue";
import CodeViewer from "~/components/misc/CodeViewer.vue";
import Modal from "~/components/misc/Modal.vue";
import SecondaryMenu from "~/components/menu/SecondaryMenu.vue";

const router = useRouter();

const data = reactive({
	code: "",
	lang: "auto",
	openAbout: false,
	tabSize: 2
});

const props = defineProps({
	args: Array
});

function copy() {
	navigator.clipboard.writeText(data.code);
}

onMounted(async () => {
	const { args } = props;
	if (args.length != 1) {
		router.push("/");
		return;
	}

	const res = await fetch(import.meta.env.VITE_API_URL + "?q=" + args[0]);
	if (!res.ok) {
		router.push("/");
		return;
	}

	const { code, lang } = await res.json();
	data.code = code;
	data.lang = lang;
});
</script>

<template>
	<CodeViewer
		:code="data.code"
		:lang="data.lang"
		:style="'tab-size:' + data.tabSize"
	/>
	<SecondaryMenu
		@change-tab-size="data.tabSize = 6 - data.tabSize"
		@copy="copy"
		@open-about="data.toggleAbout = true"
		:tab-size="data.tabSize"
	/>
	<Modal @hide-modal="data.toggleAbout = false" v-if="data.toggleAbout">
		<About />
	</Modal>
</template>
