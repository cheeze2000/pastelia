<script setup>
import pastelia from "~/assets/img/pastelia.png";

const REPO_URL = "https://github.com/cheeze2000/pastelia";
</script>

<template>
	<div class="z-10 mx-6 w-full max-w-screen-sm h-fit text-one-dark-fg bg-one-dark-bg">
		<div class="flex justify-center p-3 rounded-t bg-one-dark-fg">
			<img
				alt="pastelia logo"
				class="w-16 h-16"
				:src="pastelia"
			/>
		</div>
		<div class="p-3 text-2xl rounded-b border border-one-dark-fg">
			<p class="text-center">
				Pastelia, a no-nonsense pastebin
			</p>
			<div class="mx-auto w-fit">
				<ul>
					<li>• minimalist design</li>
					<li>• syntax highlighting</li>
					<li>
						•
						<a
							class="underline cursor-pointer"
							:href="REPO_URL"
							rel="noreferrer"
							target="_blank"
						>
							open-source
						</a>
					</li>
				</ul>
			</div>
		</div>
	</div>
</template>
