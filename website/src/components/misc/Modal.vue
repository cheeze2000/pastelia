<template>
	<div class="flex fixed top-0 justify-center items-center w-full h-full">
		<div
			class="fixed top-0 z-0 w-full h-full bg-one-dark-bg opacity-75"
			@click="$emit('hideModal')"
		/>
		<slot></slot>
	</div>
</template>
