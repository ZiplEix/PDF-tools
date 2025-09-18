import { PUBLIC_API_BASE } from "$env/static/public";

// Base URL de l'API. Peut être défini via PUBLIC_API_BASE, sinon fallback localhost.
export const API_BASE: string = PUBLIC_API_BASE || "http://localhost:8080";
