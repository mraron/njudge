import {findRouteIndex} from "./FindRouteIndex";

export async function updatePageData(location, routesToFetch, abortController, setData, setLoadingCount, isMounted) {
    const fullPath = location.pathname + location.search
    if (findRouteIndex(routesToFetch, location.pathname) !== -1) {
        setLoadingCount(arg => arg + 1)
        try {
            const response = await fetch(`/api/v2${fullPath}`, {signal: abortController.signal})
            if (!response.ok) {
                console.error("Network response was not ok.")
                setLoadingCount(arg => arg - 1)
                return
            }
            const newData = await response.json()
            if (isMounted()) {
                setData(newData)
                setLoadingCount(arg => arg - 1)
            }
        } catch {
            setLoadingCount(arg => arg - 1)
        }
    }
}