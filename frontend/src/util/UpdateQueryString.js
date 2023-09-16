import queryString from "query-string";

function UpdateQueryString(location, navigate, arg, value, validArgs) {
    const qString = location.search
    const qData = queryString.parse(qString)
    qData[arg] = value

    let urlNew = `${location.pathname}?${queryString.stringify(qData)}`
    if (validArgs) {
        urlNew = queryString.pick(urlNew, validArgs)
    }
    return navigate(urlNew)
}

export default UpdateQueryString