<script lang="ts">
  const items = [
    { href: '/pdf/merge', title: 'Fusionner PDF', desc: 'Combinez plusieurs PDF en un seul.' },
    { href: '/pdf/split', title: 'Scinder PDF', desc: 'Un fichier par page (ZIP).' },
    { href: '/pdf/extract', title: 'Extraire pages', desc: 'Plages type 1-3,5.' },
    { href: '/pdf/reorder', title: 'Réorganiser / Dupliquer / Supprimer', desc: 'Ex: 3,1,1,4-7.' },
    { href: '/pdf/rotate', title: 'Pivoter pages', desc: '90/180/270°, pages ciblées.' }
  ] as const;
</script>

<svelte:head>
  <title>Outils PDF — PDF Tools</title>
  <meta name="description" content="Actions PDF: fusion, scission, extraction, réorganisation et rotation des pages." />
</svelte:head>

<section class="relative px-6 sm:px-8 pt-16 pb-10 sm:pt-24 sm:pb-12 max-w-6xl mx-auto text-center">
  <h1 class="text-4xl sm:text-5xl font-semibold leading-tight tracking-tight">Gérez vos PDF facilement</h1>
  <p class="mt-4 sm:mt-5 text-zinc-500 max-w-2xl mx-auto">Fusionner, scinder, extraire, réorganiser ou pivoter des pages — sans inscription, simplement.</p>
  <div class="mt-7 flex flex-col sm:flex-row items-center justify-center gap-3">
    <a href="/pdf/merge" class="inline-flex items-center justify-center rounded-full px-5 py-3 text-white bg-rose-600 hover:bg-rose-700 transition text-sm font-medium">Fusionner PDF</a>
    <a href="/pdf/split" class="inline-flex items-center justify-center rounded-full px-5 py-3 text-zinc-700 bg-white border border-zinc-300 hover:bg-zinc-50 transition text-sm font-medium">Scinder PDF</a>
    <a href="/images" class="inline-flex items-center gap-2 text-sm text-zinc-600 hover:text-zinc-900">Besoin d’outils images →</a>
  </div>
  </section>

<section class="px-6 sm:px-8 pt-12 pb-16 max-w-6xl mx-auto text-zinc-800">
  <h1 class="text-3xl sm:text-4xl font-semibold">Outils PDF</h1>
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
