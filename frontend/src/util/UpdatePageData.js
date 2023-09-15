import {findRouteIndex} from "./RouteUtil";

export async function updatePageData(location, routesToFetch, setData, setLoadingCount, isMounted) {
    const fullPath = location.pathname + location.search
    if (findRouteIndex(routesToFetch, location.pathname) !== -1) {
        setLoadingCount(arg => arg + 1)
        const response = await fetch(`/api/v2${fullPath}`)
        if (!response.ok) {
            console.error("Network response was not ok.")
            setLoadingCount(arg => arg - 1)
            return
        }
        const newData = await response.json()
        if (isMounted()) {
            setData(newData)
        }
        setLoadingCount(arg => arg - 1)
    }
}