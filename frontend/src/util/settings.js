export async function saveSettings(showUnsolved, hideSolved) {
    const requestOptions = {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({
            showUnsolved: showUnsolved,
            hideSolved: hideSolved,
        })
    }
    const response = await fetch("/api/v2/user/profile/save_settings/other/", requestOptions)
    const data = await response.json()
    return {...data, success: response.ok}
}

export async function changePassword(oldPw, newPw, newPwConfirm) {
    const requestOptions = {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({
            oldPw: oldPw,
            newPw: newPw,
            newPwConfirm: newPwConfirm
        })
    }
    const response = await fetch("/api/v2/user/profile/save_settings/change_password/", requestOptions)
    const data = await response.json()
    return {...data, success: response.ok}
}