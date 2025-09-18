<script lang="ts">
	import { onMount } from 'svelte';
	import { X } from '@lucide/svelte';

	let { file, index, onRemove }: { file: File; index?: number; onRemove?: (i: number) => void } = $props();

	let hovered = $state(false);
	let canvasEl: HTMLCanvasElement | null = $state(null);
	let loading = $state(true);
	let errorMsg: string | null = $state(null);
	let pageCount: number | null = $state(null);

	function formatSize(n: number) {
		if (n < 1024) return `${n} o`;
		if (n < 1024 * 1024) return `${(n / 1024).toFixed(1)} Ko`;
		if (n < 1024 * 1024 * 1024) return `${(n / 1024 / 1024).toFixed(1)} Mo`;
		return `${(n / 1024 / 1024 / 1024).toFixed(1)} Go`;
	}

	onMount(async () => {
		try {
			const pdfjs: any = await import('pdfjs-dist');
			const { getDocument, GlobalWorkerOptions } = pdfjs;
			try {
				const workerMod: any = await import('pdfjs-dist/build/pdf.worker.min.mjs?worker');
				GlobalWorkerOptions.workerPort = new workerMod.default();
			} catch {
				const workerUrlMod: any = await import('pdfjs-dist/build/pdf.worker.min.mjs?url');
				GlobalWorkerOptions.workerSrc = workerUrlMod.default || workerUrlMod;
			}

			const data = await file.arrayBuffer();
			const loadingTask = getDocument({ data });
			const pdf = await loadingTask.promise;
			pageCount = pdf.numPages ?? null;
			const page = await pdf.getPage(1);

			const viewport = page.getViewport({ scale: 1 });
			const baseWidth = 420;
			const dpr = (globalThis as any).devicePixelRatio || 1;
			const scale = (baseWidth / viewport.width) * dpr;
			const scaled = page.getViewport({ scale });

			if (!canvasEl) return;
			const canvas = canvasEl as HTMLCanvasElement;
			const ctx = canvas.getContext('2d');
			if (!ctx) throw new Error('Canvas context non disponible');

			canvas.width = Math.ceil(scaled.width);
			canvas.height = Math.ceil(scaled.height);

			const renderTask = page.render({ canvasContext: ctx, viewport: scaled, canvas });
			await renderTask.promise;

			canvas.style.width = '100%';
			canvas.style.height = 'auto';
			loading = false;
		} catch (e: any) {
			errorMsg = e?.message ?? 'Erreur de rendu PDF';
			loading = false;
		}
	});
</script>

<div class="inline-block w-full">
	<div
		class="relative min-h-[240px] rounded-md border border-border overflow-visible"
		role="group"
		onmouseenter={() => (hovered = true)}
		onmouseleave={() => (hovered = false)}
	>
		{#if hovered}
			<div class="absolute -top-2 left-2 translate-y-[-100%] z-20 rounded-md border border-border bg-popover text-popover-foreground px-2 py-1 text-xs shadow">
				{formatSize(file.size)}{#if pageCount !== null} • {pageCount} page{pageCount > 1 ? 's' : ''}{/if}
			</div>
		{/if}
		<canvas bind:this={canvasEl} class="block w-full bg-white"></canvas>

		{#if loading}
			<div class="absolute inset-0 flex items-center justify-center text-sm text-muted-foreground bg-muted/50">Chargement…</div>
		{:else if errorMsg}
			<div class="absolute inset-0 flex items-center justify-center text-sm text-destructive/80 bg-muted/50">{errorMsg}</div>
		{/if}

		{#if onRemove && index !== undefined}
			<button
				type="button"
				aria-label="Supprimer"
				class="absolute top-2 right-2 z-10 h-8 w-8 rounded-full border border-border bg-background/90 text-foreground shadow-sm transition-opacity inline-flex items-center justify-center hover:cursor-pointer hover:bg-destructive/5 hover:text-destructive"
				style:opacity={hovered ? 1 : 0}
				class:pointer-events-none={!hovered}
				class:pointer-events-auto={hovered}
				onclick={() => onRemove && index !== undefined && onRemove(index)}
			>
				<X class="h-4 w-4" />
			</button>
		{/if}
	</div>
	<div class="mt-2 text-xs text-muted-foreground truncate flex justify-center" title={file.name}>{file.name}</div>
</div>

<style>
	.bg-muted\/50 { background-color: color-mix(in oklab, var(--muted) 50%, transparent); }
	.bg-white { background: white; }
	.truncate { white-space: nowrap; overflow: hidden; text-overflow: ellipsis; }
</style>
