import { authenticate } from "./auth";
import { apiRoute } from "../config/RouteConfig";

export async function updatePageData(
    location,
    abortController,
    setData,
    isMounted,
) {
    try {
        const fullPath = location.pathname + location.search;
        const response = await fetch(apiRoute(fullPath), {
            signal: abortController.signal,
        });
        const newData = await response.json();

        if (!response.ok) {
            console.error("Network response was not ok.");
            return;
        }
        if (isMounted()) {
            setData(newData);
        }
    } catch (error) {
        console.error(error);
    }
}

export async function updateData(
    location,
    abortController,
    setData,
    setUserData,
    setLoggedIn,
    isMounted,
) {
    const response = await authenticate();
    console.log(JSON.stringify(response))
    console.log(JSON.stringify(response !== null))
    if (response !== undefined) {
        setUserData(response);
        setLoggedIn(response !== null);
    }
    await updatePageData(location, abortController, setData, isMounted);
}
