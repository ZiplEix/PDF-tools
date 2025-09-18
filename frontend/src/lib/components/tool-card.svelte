<script lang="ts">
    import { Card, CardHeader, CardTitle, CardDescription, CardContent } from "$lib/components/ui/card";
    import { cn } from "$lib/utils";

    type Props = {
        href: string;
        title: string;
        desc?: string;
        dotClass?: string;
        class?: string;
    };

    let {
        href,
        title,
        desc = "",
        dotClass = "bg-zinc-500",
        class: className,
    }: Props = $props();
</script>

<a
    href={href}
    aria-label={title}
    class={cn(
        "block rounded-xl outline-none focus-visible:ring-[3px] focus-visible:ring-ring/50 focus-visible:border-ring",
        className
    )}
>
    <Card class="hover:shadow-md hover:-translate-y-0.5 transition-all cursor-pointer">
        <CardHeader class="">
            <CardTitle class="text-base flex items-center gap-2">
                <span class={cn("size-2 rounded-full", dotClass)}></span>
                {title}
            </CardTitle>
        </CardHeader>
        {#if desc}
            <CardContent class="pt-0 text-sm text-muted-foreground">{desc}</CardContent>
        {/if}
    </Card>
</a>
