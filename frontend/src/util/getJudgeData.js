export async function getLanguages() {
    const response = await fetch("/api/v2/data/languages/");
    const data = await response.json();
    return { ...data, success: response.ok };
}

export async function getCategories() {
    const response = await fetch("/api/v2/data/categories/");
    const data = await response.json();
    return { ...data, success: response.ok };
}

export async function getTags() {
    const response = await fetch("/api/v2/data/tags/");
    const data = await response.json();
    return { ...data, success: response.ok };
}

export function getHighlightCodes() {
    return {
        cpp11: "cpp",
        cpp14: "cpp",
        cpp17: "cpp",
        cs: "cs",
        java: "java",
        python3: "python",
    };
}