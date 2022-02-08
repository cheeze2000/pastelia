<script setup>
import { computed, onMounted, reactive, ref } from "vue";

import ReturnIcon from "~/components/icons/ReturnIcon.vue";

import languages from "~/assets/languages.json";

const emit = defineEmits(["language"]);
const search = ref(null);

const data = reactive({
	lang: computed(() => bestMatch(data.query)),
	query: ""
});

function bestMatch(query) {
	query = query.trim().toLowerCase();
	if (query == "") return "auto";
	if (query in languages) return query;

	const matches = Object.keys(languages).filter(lang => {
		let i = 0;
		for (const c of lang) {
			if (c == query[i]) i++;
			if (i == query.length) return true;
		}

		return false;
	});

	if (matches.length == 0) return "plaintext";
	return matches.sort((a, b) => a.length - b.length)[0];
}

function handleKeydown(e) {
	if (e.keyCode != 13) return;
	emit("language", data.lang);
}

onMounted(() => search.value.focus());
</script>

<template>
	<div class="z-10 mx-6 w-full max-w-screen-sm h-fit text-one-dark-fg bg-one-dark-bg">
		<input
			class="p-3 w-full text-3xl bg-one-dark-bg rounded-md border border-one-dark-fg outline-none placeholder:brightness-50"
			@keydown="handleKeydown"
			placeholder="e.g. html, css, js"
			ref="search"
			spellcheck="false"
			v-model="data.query"
		/>
		<div class="flex justify-between items-center mt-3 p-3 w-full text-3xl bg-one-dark-bg rounded-md border border-one-dark-fg outline-none placeholder:brightness-50">
			<span>
				language: {{ data.lang }}
			</span>
			<div
				class="px-3 py-1 text-one-dark-bg bg-one-dark-fg rounded-md cursor-pointer"
				@click="emit('language', data.lang)"
			>
				<ReturnIcon />
			</div>
		</div>
	</div>
</template>
