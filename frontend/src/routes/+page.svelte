<script lang="ts">
    import ToolCard from "$lib/components/tool-card.svelte";
    import { imageTools, otherTools, pdfTools } from "$lib/tools";
    import PageHeader from "./PageHeader.svelte";
</script>

<svelte:head>
    <title>PDF Tools — Outils PDF & Images en ligne</title>
    <meta name="description" content="Une suite d’outils PDF et Images, simple et rapide. Fusion, compression, conversion et plus, sans inscription." />
 </svelte:head>

<div class="bg-background text-foreground antialiased">
	<PageHeader />

    <section class="px-6 sm:px-8 max-w-6xl mx-auto">
        <div class="grid gap-8">
            <div>
                <div class="flex items-center gap-2 mb-3"><span class="size-2 rounded-full bg-rose-600"></span><h2 class="text-xl font-semibold">PDF</h2></div>
                <div class="grid gap-4 sm:grid-cols-2 md:grid-cols-3">
                    {#each pdfTools as t}
                        <ToolCard href={t.href} title={t.title} desc={t.desc} dotClass="bg-rose-500" />
                    {/each}
                </div>
            </div>

            <div>
                <div class="flex items-center gap-2 mb-3"><span class="size-2 rounded-full bg-sky-600"></span><h2 class="text-xl font-semibold">Images</h2></div>
                <div class="grid gap-4 sm:grid-cols-2 md:grid-cols-3">
                    {#each imageTools as t}
                        <ToolCard href={t.href} title={t.title} desc={t.desc} dotClass="bg-sky-500" />
                    {/each}
                </div>
            </div>

            <div>
                <div class="flex items-center gap-2 mb-3"><span class="size-2 rounded-full bg-zinc-600"></span><h2 class="text-xl font-semibold">Autres</h2></div>
                <div class="grid gap-4 sm:grid-cols-2 md:grid-cols-3">
                    {#each otherTools as t}
                        <ToolCard href={t.href} title={t.title} desc={t.desc} dotClass="bg-zinc-500" />
                    {/each}
                </div>
            </div>
        </div>
    </section>
</div>
