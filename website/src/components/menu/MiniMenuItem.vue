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
		<div class="w-16 bg-one-dark-bg">
			<button
				class="flex justify-center items-center mx-auto w-12 h-12 text-3xl rounded-md border border-one-dark-fg"
				@mouseenter="data.hovered = true"
				@mouseleave="data.hovered = false"
			>
				<slot></slot>
			</button>
		</div>
	</div>
</template>
