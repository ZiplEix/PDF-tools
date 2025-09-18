<script lang="ts">
  const items = [
    { href: '/images/convert', title: 'Convertir Images', desc: 'PNG, JPG, WEBP, SVG…' },
    { href: '/images/resize', title: 'Redimensionner', desc: 'Changer largeur/hauteur rapidement.' },
    { href: '/images/img2pdf', title: 'Images → PDF', desc: 'Assembler des images en PDF.' }
  ] as const;
</script>

<svelte:head>
  <title>Outils Images — PDF Tools</title>
  <meta name="description" content="Toutes les actions images: conversion, redimensionnement, génération de PDF et plus." />
</svelte:head>

<section class="relative px-6 sm:px-8 pt-16 pb-10 sm:pt-24 sm:pb-12 max-w-6xl mx-auto text-center">
  <h1 class="text-4xl sm:text-5xl font-semibold leading-tight tracking-tight">Travaillez vos images en un clin d’œil</h1>
  <p class="mt-4 sm:mt-5 text-zinc-500 max-w-2xl mx-auto">Convertir, redimensionner, et créer des PDF à partir d’images — rapidement et gratuitement.</p>
  <div class="mt-7 flex flex-col sm:flex-row items-center justify-center gap-3">
    <a href="/images/convert" class="inline-flex items-center justify-center rounded-full px-5 py-3 text-white bg-sky-600 hover:bg-sky-700 transition text-sm font-medium">Convertir des images</a>
    <a href="/images/resize" class="inline-flex items-center justify-center rounded-full px-5 py-3 text-zinc-700 bg-white border border-zinc-300 hover:bg-zinc-50 transition text-sm font-medium">Redimensionner</a>
    <a href="/pdf" class="inline-flex items-center gap-2 text-sm text-zinc-600 hover:text-zinc-900">Aller aux outils PDF →</a>
  </div>
</section>

<section class="px-6 sm:px-8 pt-12 pb-16 max-w-6xl mx-auto text-zinc-800">
  <h1 class="text-3xl sm:text-4xl font-semibold">Outils Images</h1>
  <p class="mt-3 text-zinc-500">Choisissez une action pour commencer.</p>

  <div class="mt-8 grid gap-4 sm:grid-cols-2 md:grid-cols-3">
    {#each items as it}
      <a href={it.href} class="block rounded-lg bg-white p-5 border border-zinc-400 hover:-translate-y-0.5 hover:shadow-lg transition-all">
        <div class="text-lg font-medium">{it.title}</div>
        <div class="text-sm text-zinc-500">{it.desc}</div>
      </a>
    {/each}
  </div>
</section>
