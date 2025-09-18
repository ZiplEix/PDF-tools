export const pdfTools = [
    { href: '/pdf/merge', title: 'Fusionner PDF', desc: 'Combinez plusieurs PDF en un seul.' },
    { href: '/pdf/split', title: 'Scinder PDF', desc: 'Séparez un PDF en plusieurs fichiers.' },
    { href: '/pdf/extract', title: 'Extraire pages', desc: 'Créez un PDF à partir de plages données.' },
    { href: '/pdf/reorder', title: 'Réorganiser / Dupliquer / Supprimer', desc: 'Spécifiez un nouvel ordre de pages.' },
    { href: '/pdf/rotate', title: 'Pivoter pages', desc: 'Rotation 90/180/270° de pages choisies.' }
] as const;

export const imageTools = [
    { href: '/images/convert', title: 'Convertir images', desc: 'PNG, JPG, WEBP, SVG…' },
    { href: '/images/resize', title: 'Redimensionner', desc: 'Ajustez les dimensions en pixels.' },
    { href: '/images/img2pdf', title: 'Images → PDF', desc: 'Assemblez des images en PDF.' },
    { href: '/images/compress', title: 'Compresser images', desc: 'Allégez PNG/JPG/WEBP.' }
] as const;

export const otherTools = [
    { href: '/tools/ocr', title: 'OCR (à venir)', desc: 'Reconnaissance de texte.' },
    { href: '/tools/meta', title: 'Métadonnées', desc: 'Lire/éditer des métadonnées.' }
] as const;
