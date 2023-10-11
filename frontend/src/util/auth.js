import { apiRoute } from "../config/RouteConfig";
import Cookies from "js-cookie";
import fetchWithCredentials from "./fetchWithCredentials";

export async function login(username, password) {
    const requestOptions = {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
            username: username,
            password: password,
        }),
    }
    try {
        const response = await fetchWithCredentials(apiRoute("/user/auth/login/"), requestOptions)
        const data = await response.json()
        return { ...data, success: response.ok }
    } catch (error) {
        console.error(error)
    }
}

export async function register(username, email, password, passwordConfirm) {
    const requestOptions = {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
            username: username,
            email: email,
            password: password,
            passwordConfirm: passwordConfirm,
        }),
    }
    try {
        const response = await fetchWithCredentials(apiRoute("/user/auth/register/"), requestOptions)
        const data = await response.json()
        return { ...data, success: response.ok }
    } catch (error) {
        console.error(error)
    }
}

export async function change_password(email) {
    const requestOptions = {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
            email: email,
        }),
    }
    try {
        const response = await fetchWithCredentials(apiRoute("/user/auth/forgotten_password/"), requestOptions)
        return response.ok
    } catch (error) {
        console.error(error)
        return false
    }
}

export async function reset_password(user, token, password, passwordConfirm) {
    const requestOptions = {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
            password: password,
            passwordConfirm: passwordConfirm,
        }),
    }
    try {
        const response = await fetchWithCredentials(
            apiRoute(`/user/auth/reset_password/${user}/${token}/`),
            requestOptions,
        )
        const data = await response.json()
        return { ...data, success: response.ok }
    } catch (error) {
        console.error(error)
    }
}

export async function verify(token) {
    const response = await fetchWithCredentials(apiRoute(`/user/auth/verify/${token}/`))
    const data = await response.json()
    const result = { ...data, success: response.ok }
    if (response.ok) {
        Cookies.set("authToken", data.authToken, { expires: 7, secure: true })
    }
    return result
}

export async function logout() {
    try {
        const response = await fetchWithCredentials(apiRoute("/user/auth/logout/"))
        return response.ok
    } catch (error) {
        console.error(error)
        return false
    }
}

export async function authenticate() {
    try {
        const response = await fetchWithCredentials(apiRoute("/user/auth/"))
        const data = await response.json()
        return data.userData
    } catch (error) {
        console.error(error)
    }
}
