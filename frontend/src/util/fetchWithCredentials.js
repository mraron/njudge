async function fetchWithCredentials(route, options = {}) {
    return await fetch(route, {...options})
}

export default fetchWithCredentials