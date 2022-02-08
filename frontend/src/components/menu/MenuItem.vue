<script setup>
import { reactive } from "vue";

defineProps({
	tooltip: String
});

const data = reactive({
	hovered: false
});
</script>

<template>
	<div class="flex justify-end items-center text-one-dark-fg">
		<span class="p-2 mr-3 text-xl bg-one-dark-bg rounded-md border border-one-dark-fg" v-if="data.hovered && tooltip">
			{{ tooltip }}
		</span>
		<button
			class="flex justify-center items-center w-16 h-16 text-4xl bg-one-dark-bg rounded-md border border-one-dark-fg"
			@mouseenter="data.hovered = true"
			@mouseleave="data.hovered = false"
		>
			<slot></slot>
		</button>
	</div>
</template>
