<script setup>
import { computed, nextTick, reactive, ref, watch } from "vue";
import hljs from "highlight.js";

import languages from "~/assets/languages.json";

const props = defineProps({
	code: String,
	lang: String
});

const codeView = ref(null);
const data = reactive({
	langClass: computed(() => {
		const lang = languages[props.lang];
		if (lang == "auto") return "";
		return "language-" + lang;
	}),
	numbers: computed(() => {
		const n = props.code.split("\n").length;
		return [...Array(n)].map((_, i) => i + 1).join("\n");
	})
});

watch(props, async () => {
	await nextTick();
	hljs.highlightElement(codeView.value);
});
</script>

<template>
	<div class="flex w-full min-h-full text-3xl text-one-dark-fg whitespace-pre bg-one-dark-bg">
		<div class="pl-10 pr-6 py-6 w-fit text-right brightness-50 select-none">
			{{ data.numbers }}
		</div>
		<div
			class="grow py-6 overflow-x-auto"
			:class="data.langClass"
			ref="codeView"
		>
			{{ code }}
		</div>
	</div>
</template>
