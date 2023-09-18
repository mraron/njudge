import Cookies from 'js-cookie';

export async function login(username, password) {
    const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            username: username,
            password: password
        })
    }
    try {
        const response = await fetch("/api/v2/user/auth/login/", requestOptions)
        const data = await response.json()
        if (!response.ok) {
            return {success: false, message: data.error}
        }
        console.log(data.authToken)
        Cookies.set("authToken", data.authToken, { expires: 7, secure: true })
        return {success: true, message: "A bejelentkezés sikeres."}
    } catch(error) {
        console.error(error)
    }
}

export function logout() {
    if (Cookies.get("authToken")) {
        Cookies.remove("authToken")
        return true
    }
    return false
}

export async function authenticate() {
    try {
        const response = await fetch("/api/v2/user/auth/")
        const data = await response.json()
        return data.userData
    } catch (error) {
        console.error(error)
    }
}