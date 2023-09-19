import {findRouteIndex} from "./findRouteIndex";
import {authenticate} from "./auth";

export async function updatePageData(location,
                                     routesToFetch,
                                     abortController,
                                     setData,
                                     isMounted) {

    const fetchedRouteIndex = findRouteIndex(routesToFetch, location.pathname)
    if (fetchedRouteIndex !== -1) {
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
        }
    }
}

export async function updateData(location,
                           routesToFetch,
                           abortController,
                           setData,
                           setUserData,
                           setLoggedIn,
                           isMounted) {

    const response = await authenticate()
    if (response !== undefined) {
        setUserData(response)
        setLoggedIn(response !== null)
    }
    await updatePageData(location, routesToFetch, abortController, setData, isMounted)
}