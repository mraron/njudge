export function trimRoute(arg) {
    return arg.replace(/\/+$/, '')
}

export function matchRoute(arg1, arg2) {
    return trimRoute(arg1) === trimRoute(arg2)
}