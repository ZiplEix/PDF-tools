import axios, { type AxiosInstance } from "axios";
import { PUBLIC_API_BASE } from "$env/static/public";

export const API_BASE = PUBLIC_API_BASE || "http://localhost:8080";

export const api: AxiosInstance = axios.create({
    baseURL: API_BASE,
    withCredentials: false,
    // Pas de Content-Type global: axios gère form-data automatiquement
});

export function filenameFromContentDisposition(cd: string | null | undefined, fallback = "download.bin"): string {
    if (!cd) return fallback;
    const match = /filename\*=UTF-8''([^;]+)|filename="?([^";]+)"?/i.exec(cd);
    const value = match?.[1] || match?.[2];
    return value ? decodeURIComponent(value) : fallback;
}

// POST /pdf/merge - FormData files[] => PDF Blob
export async function mergePdf(files: File[]) {
    const fd = new FormData();
    for (const f of files) fd.append("files", f);
    const res = await api.post(`/pdf/merge`, fd, { responseType: "blob" });
    return res;
}

// POST /pdf/split - FormData file => ZIP Blob
export async function splitPdf(file: File) {
    const fd = new FormData();
    fd.append("file", file);
    const res = await api.post(`/pdf/split`, fd, { responseType: "blob" });
    return res;
}

// POST /pdf/extract?ranges=... - FormData file => PDF Blob
export async function extractPages(file: File, ranges: string) {
    const fd = new FormData();
    fd.append("file", file);
    const res = await api.post(`/pdf/extract`, fd, { params: { ranges }, responseType: "blob" });
    return res;
}

// POST /pdf/reorder?order=... - FormData file => PDF Blob
export async function reorderPages(file: File, order: string) {
    const fd = new FormData();
    fd.append("file", file);
    const res = await api.post(`/pdf/reorder`, fd, { params: { order }, responseType: "blob" });
    return res;
}

// POST /pdf/rotate?angle=...&pages=... - FormData file => PDF Blob
export async function rotatePages(file: File, angle: string, pages = "all") {
    const fd = new FormData();
    fd.append("file", file);
    const res = await api.post(`/pdf/rotate`, fd, { params: { angle, pages }, responseType: "blob" });
    return res;
}

