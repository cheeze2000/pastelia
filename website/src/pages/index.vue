<script setup>
import { reactive } from "vue";
import { useRouter } from "vue-router";

import About from "~/components/misc/About.vue";
import CodeEditor from "~/components/misc/CodeEditor.vue";
import LanguageSelector from "~/components/misc/LanguageSelector.vue";
import Modal from "~/components/misc/Modal.vue";
import PrimaryMenu from "~/components/menu/PrimaryMenu.vue";

const router = useRouter();

const data = reactive({
	code: "",
	lang: "auto",
	toggleAbout: false,
	toggleLanguageSelector: false,
	tabSize: 2
});

function setLanguage(lang) {
	data.lang = lang;
	data.toggleLanguageSelector = false;
}

async function upload() {
	const code = data.code.trimEnd();
	if (code == "") return;

	const lang = data.lang;
	const res = await fetch(import.meta.env.VITE_API_URL, {
		method: "POST",
		body: JSON.stringify({ code, lang })
	});

	if (res.ok) {
		const slug = await res.text();
		router.push(slug);
	}
}
</script>

<template>
	<CodeEditor
		:style="'tab-size:' + data.tabSize"
		v-model="data.code"
	/>
	<PrimaryMenu
		@change-language="data.toggleLanguageSelector = true"
		@change-tab-size="data.tabSize = 6 - data.tabSize"
		@open-about="data.toggleAbout = true"
		@upload="upload"
		:lang="data.lang"
		:tab-size="data.tabSize"
	/>
	<Modal @hide-modal="data.toggleLanguageSelector = false" v-if="data.toggleLanguageSelector">
		<LanguageSelector @language="setLanguage" />
	</Modal>
	<Modal @hide-modal="data.toggleAbout = false" v-if="data.toggleAbout">
		<About />
	</Modal>
</template>
