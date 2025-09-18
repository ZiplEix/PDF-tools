<script lang="ts">
	const pdfTools = [
        { href: '/pdf/merge', title: 'Fusionner PDF', desc: 'Combinez plusieurs PDF en un seul.' },
        { href: '/pdf/split', title: 'Scinder PDF', desc: 'Séparez un PDF en plusieurs fichiers.' },
        { href: '/pdf/extract', title: 'Extraire pages', desc: 'Créez un PDF à partir de plages données.' },
        { href: '/pdf/reorder', title: 'Réorganiser / Dupliquer / Supprimer', desc: 'Spécifiez un nouvel ordre de pages.' },
        { href: '/pdf/rotate', title: 'Pivoter pages', desc: 'Rotation 90/180/270° de pages choisies.' }
	] as const;

	const imageTools = [
		{ href: '/images/convert', title: 'Convertir images', desc: 'PNG, JPG, WEBP, SVG…' },
		{ href: '/images/resize', title: 'Redimensionner', desc: 'Ajustez les dimensions en pixels.' },
		{ href: '/images/img2pdf', title: 'Images → PDF', desc: 'Assemblez des images en PDF.' },
		{ href: '/images/compress', title: 'Compresser images', desc: 'Allégez PNG/JPG/WEBP.' }
	] as const;

	const otherTools = [
		{ href: '/tools/ocr', title: 'OCR (à venir)', desc: 'Reconnaissance de texte.' },
		{ href: '/tools/meta', title: 'Métadonnées', desc: 'Lire/éditer des métadonnées.' }
	] as const;
</script>

<svelte:head>
    <title>PDF Tools — Outils PDF & Images en ligne</title>
    <meta name="description" content="Une suite d’outils PDF et Images, simple et rapide. Fusion, compression, conversion et plus, sans inscription." />
</svelte:head>

<div class="bg-gradient-to-b from-zinc-100 via-zinc-100 to-zinc-100 text-zinc-800 antialiased">
	<section class="relative px-6 sm:px-8 pt-16 pb-12 sm:pt-24 sm:pb-16 max-w-6xl mx-auto text-center">
		<h1 class="text-4xl sm:text-6xl font-semibold leading-tight tracking-tight">
			Des outils PDF & Images<br class="hidden sm:block"/> simples et rapides
		</h1>
        <p class="mt-4 sm:mt-6 text-zinc-500 max-w-2xl mx-auto">
            Fusionnez, scindez, compressez, convertissez… Accédez aux pages dédiées, sans inscription.
        </p>

        <div class="mt-8 flex flex-col sm:flex-row items-center justify-center gap-3">
            <a href="/pdf" class="inline-flex items-center justify-center rounded-full px-5 py-3 text-white bg-rose-600 hover:bg-rose-700 transition text-sm font-medium">
                Outils PDF
            </a>
            <a href="/images" class="inline-flex items-center justify-center rounded-full px-5 py-3 text-white bg-sky-600 hover:bg-sky-700 transition text-sm font-medium">
                Outils Images
            </a>
            <a href="/tools" class="inline-flex items-center justify-center rounded-full px-5 py-3 text-zinc-700 bg-white border border-zinc-200 hover:bg-zinc-50 transition text-sm font-medium">
                Tous les outils
            </a>
        </div>
	</section>

    <section class="px-6 sm:px-8 max-w-6xl mx-auto">
        <div class="grid gap-8">
            <div>
                <div class="flex items-center gap-2 mb-3"><span class="size-2 rounded-full bg-rose-600"></span><h2 class="text-xl font-semibold">PDF</h2></div>
                <div class="grid gap-4 sm:grid-cols-2 md:grid-cols-3">
                    {#each pdfTools as t}
                        <a href={t.href} class="group block rounded-xl bg-white p-5 border border-zinc-400 hover:-translate-y-0.5 hover:shadow-lg transition-all">
                            <div class="flex items-start gap-3">
                                <span class="mt-1 size-2 rounded-full bg-rose-500"></span>
                                <div>
                                    <div class="font-medium">{t.title}</div>
                                    <div class="text-sm text-zinc-500">{t.desc}</div>
                                </div>
                            </div>
                        </a>
                    {/each}
                </div>
            </div>

            <div>
                <div class="flex items-center gap-2 mb-3"><span class="size-2 rounded-full bg-sky-600"></span><h2 class="text-xl font-semibold">Images</h2></div>
                <div class="grid gap-4 sm:grid-cols-2 md:grid-cols-3">
                    {#each imageTools as t}
                        <a href={t.href} class="group block rounded-xl bg-white p-5 border border-zinc-400 hover:-translate-y-0.5 hover:shadow-lg transition-all">
                            <div class="flex items-start gap-3">
                                <span class="mt-1 size-2 rounded-full bg-sky-500"></span>
                                <div>
                                    <div class="font-medium">{t.title}</div>
                                    <div class="text-sm text-zinc-500">{t.desc}</div>
                                </div>
                            </div>
                        </a>
                    {/each}
                </div>
            </div>

            <div>
                <div class="flex items-center gap-2 mb-3"><span class="size-2 rounded-full bg-zinc-600"></span><h2 class="text-xl font-semibold">Autres</h2></div>
                <div class="grid gap-4 sm:grid-cols-2 md:grid-cols-3">
                    {#each otherTools as t}
                        <a href={t.href} class="group block rounded-xl bg-white p-5 border border-zinc-400 hover:-translate-y-0.5 hover:shadow-lg transition-all">
                            <div class="flex items-start gap-3">
                                <span class="mt-1 size-2 rounded-full bg-zinc-500"></span>
                                <div>
                                    <div class="font-medium">{t.title}</div>
                                    <div class="text-sm text-zinc-500">{t.desc}</div>
                                </div>
                            </div>
                        </a>
                    {/each}
                </div>
            </div>
        </div>
    </section>
</div>
