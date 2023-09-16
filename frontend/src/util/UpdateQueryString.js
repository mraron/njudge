import queryString from "query-string";

function UpdateQueryString(location, navigate, arg, value) {
    const qStringOld = location.search
    const qData = queryString.parse(qStringOld)
    qData[arg] = value

    const qStringNew = queryString.stringify(qData)
    navigate(`${location.pathname}?${qStringNew}`)
}

export default UpdateQueryString