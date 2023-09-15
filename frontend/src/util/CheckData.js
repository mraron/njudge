function checkData(data) {
    if (!data || data.processed) {
        return false
    }
    data.processed = true
    return true
}

export default checkData