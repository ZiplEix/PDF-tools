<script lang="ts">
    import { Button } from "$lib/components/ui/button";
    import ToolCard from "$lib/components/tool-card.svelte";
    import { pdfTools } from "$lib/tools";
</script>

<svelte:head>
    <title>Outils PDF — PDF Tools</title>
    <meta name="description" content="Actions PDF: fusion, scission, extraction, réorganisation et rotation des pages." />
</svelte:head>

<section class="relative px-6 sm:px-8 pt-16 pb-10 sm:pt-24 sm:pb-12 max-w-6xl mx-auto text-center">
    <h1 class="text-4xl sm:text-5xl font-semibold leading-tight tracking-tight">Gérez vos PDF facilement</h1>
    <p class="mt-4 sm:mt-5 text-zinc-500 max-w-2xl mx-auto">Fusionner, scinder, extraire, réorganiser ou pivoter des pages — sans inscription, simplement.</p>
    <div class="mt-7 flex flex-col sm:flex-row items-center justify-center gap-3">
        <Button href="/pdf/merge">Fusionner PDF</Button>
        <Button href="/pdf/split" variant="outline">Scinder PDF</Button>
        <Button href="/images" variant="link" class="text-muted-foreground">Besoin d’outils images →</Button>
    </div>
    </section>

    <section class="px-6 sm:px-8 pt-12 pb-16 max-w-6xl mx-auto">
    <h1 class="text-3xl sm:text-4xl font-semibold">Outils PDF</h1>
    <p class="mt-3 text-muted-foreground">Choisissez une action pour commencer.</p>

    <div class="mt-8 grid gap-4 sm:grid-cols-2 md:grid-cols-3">
        {#each pdfTools as it}
            <ToolCard href={it.href} title={it.title} desc={it.desc} dotClass="bg-rose-500" />
        {/each}
    </div>
</section>
