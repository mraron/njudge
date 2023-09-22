import queryString from "query-string";

function UpdateQueryString(location, navigate, args, values, validArgs, invalidArgs) {
    const qString = location.search
    const qData = queryString.parse(qString)

    args.forEach((arg, index) => {
        qData[arg] = values[index]
    })
    let urlNew = `${location.pathname}?${queryString.stringify(qData)}`
    if (validArgs) {
        urlNew = queryString.pick(urlNew, validArgs)
    }
    if (invalidArgs) {
        urlNew = queryString.exclude(urlNew, invalidArgs)
    }
    return navigate(urlNew)
}

export default UpdateQueryString