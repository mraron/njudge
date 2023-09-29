export const routeMap = {
    home: "/",
    contests: "/contests/",
    info: "/info/",
    archive: "/archive/",
    submissions: "/problemset/status/",
    problems: "/problemset/main/",
    submission: "/submission/:id/",
    profile: "/user/profile/:user/",
    profileSubmissions: "/user/profile/:user/submissions/",
    profileSettings: "/user/profile/:user/settings/",
    problem: "/problemset/main/:problem/",
    problemSubmit: "/problemset/main/:problem/submit/",
    problemSubmissions: "/problemset/main/:problem/submissions/",
    problemRanklist: "/problemset/main/:problem/ranklist/",
    login: "/user/login/",
    register: "/user/register/",
    verify: "/user/verify/:token/",
    logout: "/user/logout/",
    forgotten_password: "/user/forgotten_password/",
    reset_password: "/user/reset_password/:user/:token/",
    admin: "/user/admin/",
};

const apiRoot = "https://127.0.0.1:5619/api/v2";

export function apiRoute(route) {
    console.log(`${apiRoot}${route}`);
    return `${apiRoot}${route}`;
}
