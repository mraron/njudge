export function updatePageData(route, setData) {
    setData(null)
    fetch(route).then(res => res.json()).then(data => setData(data))
}