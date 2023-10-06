import { authenticate } from "./auth"
import { apiRoute } from "../config/RouteConfig"
import fetchWithCredentials from "./fetchWithCredentials"

export async function updatePageData(location, abortController, setData, isMounted) {
    try {
        const fullPath = location.pathname + location.search
        const response = await fetchWithCredentials(apiRoute(fullPath), {
            signal: abortController.signal,
        })
        const newData = await response.json()

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

export async function updateData(location, abortController, setData, setUserData, setLoggedIn, isMounted) {
    const response = await authenticate()
    if (response !== undefined) {
        setUserData(response)
        setLoggedIn(response !== null)
    }
    await updatePageData(location, abortController, setData, isMounted)
}
