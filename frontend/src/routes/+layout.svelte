<script lang="ts">
	import '../app.css';
	import favicon from '$lib/assets/favicon.svg';
	const year = new Date().getFullYear();
	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<div class="min-h-dvh flex flex-col bg-background text-foreground">
	<header class="sticky top-0 z-10 bg-background/90 backdrop-blur supports-[backdrop-filter]:bg-background/70 border-b border-border">
		<div class="max-w-6xl mx-auto px-6 sm:px-8 h-14 flex items-center justify-between">
			<a href="/" class="font-semibold">PDF Tools</a>
			<nav aria-label="Navigation principale" class="text-sm text-muted-foreground">
				<ul class="flex items-center gap-4 sm:gap-6">
					<li><a href="/pdf" class="hover:text-foreground">PDF</a></li>
					<li><a href="/images" class="hover:text-foreground">Images</a></li>
					<li><a href="/tools" class="hover:text-foreground">Autres</a></li>
				</ul>
			</nav>
		</div>
	</header>

	<main class="flex-1">
		{@render children?.()}
	</main>

	<footer class="px-6 sm:px-8 py-10 border-t border-border mt-8">
		<div class="max-w-6xl mx-auto flex flex-col sm:flex-row items-center justify-between gap-4 text-sm text-muted-foreground">
			<div>© {year} PDF Tools</div>
			<div class="flex items-center gap-3">
				<a href="/legal" class="hover:text-foreground transition">Mentions légales</a>
				<span class="opacity-30">•</span>
				<a href="mailto:contact@pdftools.local" class="hover:text-foreground transition">Contact</a>
			</div>
		</div>
	</footer>
</div>
