<script lang="ts">
    import { Button } from "$lib/components/ui/button";
    import ToolCard from "$lib/components/tool-card.svelte";
    import { imageTools } from "$lib/tools";
</script>

<svelte:head>
    <title>Outils Images — PDF Tools</title>
    <meta name="description" content="Toutes les actions images: conversion, redimensionnement, génération de PDF et plus." />
</svelte:head>

<section class="relative px-6 sm:px-8 pt-16 pb-10 sm:pt-24 sm:pb-12 max-w-6xl mx-auto text-center">
    <h1 class="text-4xl sm:text-5xl font-semibold leading-tight tracking-tight">Travaillez vos images en un clin d’œil</h1>
    <p class="mt-4 sm:mt-5 text-zinc-500 max-w-2xl mx-auto">Convertir, redimensionner, et créer des PDF à partir d’images — rapidement et gratuitement.</p>
    <div class="mt-7 flex flex-col sm:flex-row items-center justify-center gap-3">
        <Button href="/images/convert" class="bg-sky-600 text-white hover:bg-sky-700">Convertir des images</Button>
        <Button href="/images/resize" variant="outline">Redimensionner</Button>
        <Button href="/pdf" variant="link" class="text-muted-foreground">Aller aux outils PDF →</Button>
    </div>
</section>

<section class="px-6 sm:px-8 pt-12 pb-16 max-w-6xl mx-auto">
    <h1 class="text-3xl sm:text-4xl font-semibold">Outils Images</h1>
    <p class="mt-3 text-muted-foreground">Choisissez une action pour commencer.</p>

    <div class="mt-8 grid gap-4 sm:grid-cols-2 md:grid-cols-3">
        {#each imageTools as it}
            <ToolCard href={it.href} title={it.title} desc={it.desc} dotClass="bg-sky-500" />
        {/each}
    </div>
</section>
