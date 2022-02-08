<script setup>
import { reactive } from "vue";

import AboutIcon from "~/components/icons/AboutIcon.vue";
import IndentIcon from "~/components/icons/IndentIcon.vue";
import LanguageIcon from "~/components/icons/LanguageIcon.vue";
import MenuItem from "~/components/menu/MenuItem.vue";
import MenuItemGroup from "~/components/menu/MenuItemGroup.vue";
import MiniMenuItem from "~/components/menu/MiniMenuItem.vue";
import SettingsIcon from "~/components/icons/SettingsIcon.vue";
import UploadIcon from "~/components/icons/UploadIcon.vue";

defineEmits([
	"changeTabSize",
	"changeLanguage",
	"openAbout",
	"upload"
]);

defineProps({
	lang: String,
	tabSize: Number
});

const data = reactive({
	toggleSettings: false
});
</script>

<template>
	<div class="fixed top-6 right-6 flex flex-col gap-3 justify-between">
		<MenuItem
			@click="$emit('changeLanguage')"
			:tooltip="'language: ' + lang"
		>
			<LanguageIcon />
		</MenuItem>
		<MenuItem
			@click="$emit('upload')"
			tooltip="upload"
		>
			<UploadIcon />
		</MenuItem>
		<MenuItem
			@click="data.toggleSettings = !data.toggleSettings"
			tooltip="toggle settings"
		>
			<SettingsIcon />
		</MenuItem>
		<MenuItemGroup v-if="data.toggleSettings">
			<MiniMenuItem
				@click="$emit('changeTabSize')"
				:tooltip="'tab size: ' + tabSize"
			>
				<IndentIcon />
			</MiniMenuItem>
		</MenuItemGroup>
		<MenuItem
			@click="$emit('openAbout')"
			tooltip="about"
		>
			<AboutIcon />
		</MenuItem>
	</div>
</template>
