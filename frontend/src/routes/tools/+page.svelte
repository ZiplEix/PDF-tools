<script lang="ts">
  const items = [
    { href: '/tools/ocr', title: 'OCR (à venir)', desc: 'Reconnaissance de texte.' },
    { href: '/tools/meta', title: 'Métadonnées', desc: 'Lire/éditer les métadonnées (à venir).' }
  ] as const;
</script>

<svelte:head>
  <title>Autres outils — PDF Tools</title>
  <meta name="description" content="Autres utilitaires en développement: OCR, métadonnées et plus." />
</svelte:head>

<section class="relative px-6 sm:px-8 pt-16 pb-10 sm:pt-24 sm:pb-12 max-w-6xl mx-auto text-center">
  <h1 class="text-4xl sm:text-5xl font-semibold leading-tight tracking-tight">Boîte à outils pratique</h1>
  <p class="mt-4 sm:mt-5 text-zinc-500 max-w-2xl mx-auto">Des utilitaires complémentaires pour aller plus loin: OCR, métadonnées et plus à venir.</p>
  <div class="mt-7 flex flex-col sm:flex-row items-center justify-center gap-3">
    <a href="/tools/ocr" class="inline-flex items-center justify-center rounded-full px-5 py-3 text-white bg-emerald-600 hover:bg-emerald-700 transition text-sm font-medium">Essayer l’OCR</a>
    <a href="/tools/meta" class="inline-flex items-center justify-center rounded-full px-5 py-3 text-zinc-700 bg-white border border-zinc-300 hover:bg-zinc-50 transition text-sm font-medium">Métadonnées</a>
    <a href="/" class="inline-flex items-center gap-2 text-sm text-zinc-600 hover:text-zinc-900">Revenir à l’accueil →</a>
  </div>
</section>

<section class="px-6 sm:px-8 pt-12 pb-16 max-w-6xl mx-auto text-zinc-800">
  <h1 class="text-3xl sm:text-4xl font-semibold">Autres outils</h1>
  <p class="mt-3 text-zinc-500">Fonctionnalités à venir.</p>

  <div class="mt-8 grid gap-4 sm:grid-cols-2 md:grid-cols-3">
    {#each items as it}
      <a href={it.href} class="block rounded-lg bg-white p-5 border border-zinc-400 hover:-translate-y-0.5 hover:shadow-lg transition-all">
        <div class="text-lg font-medium">{it.title}</div>
        <div class="text-sm text-zinc-500">{it.desc}</div>
      </a>
    {/each}
  </div>
</section>
