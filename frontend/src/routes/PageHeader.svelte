<script lang="ts">
    import { Button } from "$lib/components/ui/button";
</script>

<section class="relative px-6 sm:px-8 pt-16 pb-12 sm:pt-24 sm:pb-16 max-w-6xl mx-auto text-center">
    <h1 class="text-4xl sm:text-6xl font-semibold leading-tight tracking-tight">
        Des outils PDF & Images<br class="hidden sm:block"/> simples et rapides
    </h1>
    <p class="mt-4 sm:mt-6 text-zinc-500 max-w-2xl mx-auto">
        Fusionnez, scindez, compressez, convertissez… Accédez aux pages dédiées, sans inscription.
    </p>

    <div class="mt-8 flex flex-col sm:flex-row items-center justify-center gap-3">
        <Button href="/pdf">Outils PDF</Button>
        <Button href="/images" class="bg-sky-600 text-white hover:bg-sky-700">Outils Images</Button>
        <Button href="/tools" variant="outline">Tous les outils</Button>
    </div>
</section>