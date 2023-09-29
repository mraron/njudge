async function fetchWithCredentials(route, options = {}) {
    return await fetch(route, {...options, credentials: "include"})
}

export default fetchWithCredentials