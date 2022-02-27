<script setup>
defineEmits([
	"update:modelValue"
]);

function handleKeydown(e) {
	if (e.keyCode != 9) return;
	e.preventDefault();

	const textarea = e.target;
	const code = textarea.value;
	const xs = code.slice(0, textarea.selectionStart);
	const ys = code.slice(textarea.selectionEnd);

	textarea.value = xs + "\t" + ys;
	textarea.selectionStart = xs.length + 1;
	textarea.selectionEnd = xs.length + 1;
}
</script>

<template>
	<textarea
		class="p-6 w-full h-full text-3xl text-one-dark-fg whitespace-pre bg-one-dark-bg border-none outline-none resize-none"
		@keydown="handleKeydown"
		@input="$emit('update:modelValue', $event.target.value)"
		spellcheck="false"
	/>
</template>
