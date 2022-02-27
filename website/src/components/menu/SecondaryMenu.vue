<script setup>
import { reactive } from "vue";
import { useRouter } from "vue-router";

import AboutIcon from "~/components/icons/AboutIcon.vue";
import AddIcon from "~/components/icons/AddIcon.vue";
import CopyIcon from "~/components/icons/CopyIcon.vue";
import IndentIcon from "~/components/icons/IndentIcon.vue";
import MenuItem from "~/components/menu/MenuItem.vue";
import MenuItemGroup from "~/components/menu/MenuItemGroup.vue";
import MiniMenuItem from "~/components/menu/MiniMenuItem.vue";
import SettingsIcon from "~/components/icons/SettingsIcon.vue";

const emit = defineEmits([
	"changeTabSize",
	"copy",
	"openAbout"
]);

defineProps({
	lang: String,
	tabSize: Number
});

const router = useRouter();
const data = reactive({
	isCopied: false,
	toggleSettings: false
});

function copy() {
	emit("copy");
	data.isCopied = true;
}
</script>

<template>
	<div class="fixed top-6 right-6 flex flex-col gap-3 justify-between self-end">
		<MenuItem
			@click="router.push('/')"
			tooltip="new"
		>
			<AddIcon />
		</MenuItem>
		<MenuItem
			@click="copy"
			@mouseleave="data.isCopied = false"
			:tooltip="data.isCopied ? 'copied!' : 'copy to clipboard'"
		>
			<CopyIcon />
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
