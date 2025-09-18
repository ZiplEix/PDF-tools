<script lang="ts">
    import { Button } from "$lib/components/ui/button";
    import ToolCard from "$lib/components/tool-card.svelte";
    import { otherTools } from "$lib/tools";
</script>

<svelte:head>
    <title>Autres outils — PDF Tools</title>
    <meta name="description" content="Autres utilitaires en développement: OCR, métadonnées et plus." />
</svelte:head>

<section class="relative px-6 sm:px-8 pt-16 pb-10 sm:pt-24 sm:pb-12 max-w-6xl mx-auto text-center">
    <h1 class="text-4xl sm:text-5xl font-semibold leading-tight tracking-tight">Boîte à outils pratique</h1>
    <p class="mt-4 sm:mt-5 text-zinc-500 max-w-2xl mx-auto">Des utilitaires complémentaires pour aller plus loin: OCR, métadonnées et plus à venir.</p>
    <div class="mt-7 flex flex-col sm:flex-row items-center justify-center gap-3">
        <Button href="/tools/ocr" class="bg-emerald-600 text-white hover:bg-emerald-700">Essayer l’OCR</Button>
        <Button href="/tools/meta" variant="outline">Métadonnées</Button>
        <Button href="/" variant="link" class="text-muted-foreground">Revenir à l’accueil →</Button>
    </div>
</section>

<section class="px-6 sm:px-8 pt-12 pb-16 max-w-6xl mx-auto">
    <h1 class="text-3xl sm:text-4xl font-semibold">Autres outils</h1>
    <p class="mt-3 text-muted-foreground">Fonctionnalités à venir.</p>

    <div class="mt-8 grid gap-4 sm:grid-cols-2 md:grid-cols-3">
        {#each otherTools as it}
            <ToolCard href={it.href} title={it.title} desc={it.desc} dotClass="bg-zinc-500" />
        {/each}
    </div>
</section>
