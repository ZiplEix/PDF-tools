<script lang="ts">
	import { Button, buttonVariants } from "$lib/components/ui/button";
	import { Card, CardContent } from "$lib/components/ui/card";
	import { UploadCloud, Trash2, ArrowUp, ArrowDown, FileText, Send } from "@lucide/svelte";
	import { cn } from "$lib/utils";
	import PdfCard from "$lib/components/pdf-card.svelte";

	let fileInput: HTMLInputElement | null = null;
	let files: File[] = [];
	let isDragging = false;
	let dragCounter = 0;

	function handleFileList(list: FileList | null | undefined) {
		if (!list) return;
		const incoming = Array.from(list).filter((f) => f.type === "application/pdf" || f.name.toLowerCase().endsWith(".pdf"));
		// déduplique grossièrement par name+size+mtime
		const key = (f: File) => `${f.name}|${f.size}|${f.lastModified}`;
		const seen = new Set(files.map(key));
		for (const f of incoming) {
			const k = key(f);
			if (!seen.has(k)) {
				files = [...files, f];
				seen.add(k);
			}
		}
	}

	function onInputChange(e: Event) {
		const target = e.target as HTMLInputElement;
		handleFileList(target.files ?? undefined);
		if (target && target.value) target.value = ""; // reset
	}

	function onDragEnter(e: DragEvent) {
		e.preventDefault();
		dragCounter += 1;
		isDragging = true;
	}

	function onDragOver(e: DragEvent) {
		e.preventDefault();
	}

	function onDragLeave(e: DragEvent) {
		e.preventDefault();
		dragCounter -= 1;
		if (dragCounter <= 0) {
			isDragging = false;
			dragCounter = 0;
		}
	}

	function onDrop(e: DragEvent) {
		e.preventDefault();
		handleFileList(e.dataTransfer?.files ?? undefined);
		isDragging = false;
		dragCounter = 0;
	}

	function removeAt(idx: number) {
		files = files.filter((_, i) => i !== idx);
	}

	import { mergePdf, filenameFromContentDisposition } from "$lib/api";
	let merging = false;
	let errorMsg: string | null = null;

	async function mergeNow() {
		if (merging || files.length < 2) return;
		errorMsg = null;
		merging = true;
		try {
			const res = await mergePdf(files);
			const blob = res.data as Blob;
			const cd = res.headers["content-disposition"] as string | undefined;
			let filename = filenameFromContentDisposition(cd, "merged.pdf");
			const url = URL.createObjectURL(blob);
			const a = document.createElement("a");
			a.href = url;
			a.download = filename;
			document.body.appendChild(a);
			a.click();
			a.remove();
			URL.revokeObjectURL(url);
		} catch (err) {
			// axios error handling
			const anyErr = err as any;
			if (anyErr?.response) {
				const status = anyErr.response.status;
				const data = anyErr.response.data;
				const msg = typeof data === "string" ? data : data?.message || data?.error;
				errorMsg = msg || `Erreur ${status}`;
			} else if (anyErr?.message) {
				errorMsg = anyErr.message;
			} else {
				errorMsg = "Une erreur est survenue.";
			}
		} finally {
			merging = false;
		}
	}
</script>

<svelte:window
    on:dragenter|preventDefault={onDragEnter}
    on:dragover|preventDefault={onDragOver}
    on:dragleave|preventDefault={onDragLeave}
    on:drop|preventDefault={onDrop}
/>

<svelte:head>
	<title>Fusionner PDF — PDF Tools</title>
	<meta name="description" content="Sélectionnez ou déposez plusieurs PDF pour les fusionner en un seul fichier." />
</svelte:head>

<section class="px-6 sm:px-8 pt-12 pb-16 {files.length === 0 ? "max-w-5xl" : "max-w-7xl"} mx-auto">
	<div class="text-center">
		<h1 class="text-3xl sm:text-4xl font-semibold">Fusionner des PDF</h1>
		<p class="mt-2 text-muted-foreground">Ajoutez plusieurs fichiers PDF, organisez-les dans l’ordre souhaité, puis fusionnez-les en un seul document.</p>
	</div>

	<div class="mt-8">
		<input bind:this={fileInput} id="merge-file-input" type="file" accept="application/pdf,.pdf" multiple class="sr-only" on:change={onInputChange} />

		{#if files.length === 0}
			<Card class="border-dashed border-2 border-ring bg-background/50">
				<CardContent class="py-16">
					<div class="max-w-xl mx-auto text-center">
						<UploadCloud class="mx-auto size-10 text-muted-foreground" />
						<div class="mt-4 text-lg font-medium">Glissez–déposez vos PDF ici</div>
						<div class="text-sm text-muted-foreground">ou utilisez le bouton ci‑dessous</div>
                        <div class="mt-6">
                            <label
                                for="merge-file-input"
                                class={cn(buttonVariants({ size: 'lg' }), 'h-12 px-6 inline-flex items-center justify-center cursor-pointer')}
                            >
                                Sélectionner des fichiers PDF
                            </label>
                        </div>
					</div>
				</CardContent>
			</Card>
		{:else}
			<div class="grid gap-6 lg:grid-cols-[1fr_340px]">
				<div class="min-h-[55vh] rounded-xl border border-dashed border-border bg-background/50 p-4">
					<div class="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-4 gap-4">
						{#each files as f, i}
							<div class="cursor-pointer">
								<PdfCard file={f} index={i} onRemove={removeAt} />
							</div>
						{/each}
					</div>
				</div>
				<aside class="h-full flex flex-col">
					<div class="space-y-4">
						<div>
							<h2 class="text-xl font-semibold">Merge PDF</h2>
							<div class="h-px bg-border my-3"></div>
						</div>
						<Card>
							<CardContent class="p-4 text-sm text-muted-foreground">
								Pour changer l’ordre de vos PDF, glissez‑déposez les fichiers comme vous voulez.
							</CardContent>
						</Card>
						<Card>
							<CardContent class="p-4 text-sm text-muted-foreground">
								Pour ajouter d’autres fichiers, vous pouvez les glisser‑déposer.
							</CardContent>
						</Card>
						{#if errorMsg}
							<div class="text-sm text-red-600 border border-red-200 bg-red-50 rounded-md p-3">{errorMsg}</div>
						{/if}
					</div>
					<div class="mt-auto pt-6">
						<Button size="lg" class="w-full h-14 text-base gap-2 hover:cursor-pointer" disabled={merging || files.length < 2} onclick={mergeNow}>
							<Send class="h-5 w-5" />
							{merging ? "Fusion en cours…" : "Fusionner les PDF"}
						</Button>
					</div>
				</aside>
			</div>
		{/if}
	</div>
</section>

{#if isDragging}
    <div
        class="fixed inset-0 z-20 flex items-center justify-center bg-background/80 backdrop-blur-sm border-2 border-dashed border-ring pointer-events-none"
    >
		<div class="text-center">
			<UploadCloud class="mx-auto size-8 text-muted-foreground" />
			<div class="mt-3 text-lg font-medium">Déposez vos PDF ici</div>
			<div class="text-sm text-muted-foreground">ou cliquez sur le bouton pour sélectionner</div>
		</div>
	</div>
{/if}
