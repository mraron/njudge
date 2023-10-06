function extractParams(pathname, route) {
    const routeSegments = route.split("/")
    const pathnameSegments = pathname.split("/")
    const result = {}

    routeSegments
        .map((item, index) => (item.startsWith(":") ? index : -1))
        .filter((index) => index !== -1)
        .forEach((index) => {
            result[routeSegments[index].slice(1)] = decodeURIComponent(pathnameSegments[index])
        })
    return result
}

export default extractParams
