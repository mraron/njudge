import {findRouteIndex} from "./FindRouteIndex";
import {authenticate} from "./User";

export async function updatePageData(location,
                                     routesToFetch,
                                     abortController,
                                     setData,
                                     setLoadingCount,
                                     isMounted) {

    const fetchedRouteIndex = findRouteIndex(routesToFetch, location.pathname)
    if (fetchedRouteIndex !== -1) {
        setLoadingCount(arg => arg + 1)
        try {
            const fullPath = location.pathname + location.search
            const response = await fetch(`/api/v2${fullPath}`, {signal: abortController.signal})
            const newData = await response.json()
            newData.route = routesToFetch[fetchedRouteIndex]

            if (!response.ok) {
                console.error("Network response was not ok.")
                return
            }
            if (isMounted()) {
                setData(newData)
            }
        } catch (error) {
            console.error(error)
        } finally {
            setLoadingCount(arg => arg - 1)
        }
    }
}

export function updateData(location,
                           routesToFetch,
                           abortController,
                           setData,
                           setLoadingCount,
                           setUserData,
                           setLoggedIn,
                           isMounted) {

    authenticate().then(resp => {
        setUserData(resp)
        setLoggedIn(resp != null)
    }).then(() => {
        updatePageData(location, routesToFetch, abortController, setData, setLoadingCount, isMounted)
    })
}