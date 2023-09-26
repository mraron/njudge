import Cookies from "js-cookie";
import { apiRoute } from "../config/RouteConfig";

export async function login(username, password) {
    const requestOptions = {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
            username: username,
            password: password,
        }),
    };
    try {
        const response = await fetch(
            apiRoute("/user/auth/login/"),
            requestOptions,
        );
        const data = await response.json();
        return { ...data, success: response.ok };
    } catch (error) {
        console.error(error);
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
    };
    try {
        const response = await fetch(
            apiRoute("/user/auth/register/"),
            requestOptions,
        );
        const data = await response.json();
        return { ...data, success: response.ok };
    } catch (error) {
        console.error(error);
    }
}

export async function verify(token) {
    const response = await fetch(apiRoute(`/user/auth/verify/${token}/`));
    const data = await response.json();
    const result = { ...data, success: response.ok };
    if (response.ok) {
        Cookies.set("authToken", data.authToken, { expires: 7, secure: true });
    }
    return result;
}

export async function logout() {
    try {
        const response = await fetch(apiRoute("/user/auth/logout/"));
        return response.ok;
    } catch (error) {
        console.error(error);
        return false;
    }
}

export async function authenticate() {
    try {
        const response = await fetch(apiRoute("/user/auth/"));
        const data = await response.json();
        return data.userData;
    } catch (error) {
        console.error(error);
    }
}
