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
    admin: "/user/admin/",
};

const apiRoot = "http://localhost:5555/api/v2";

export function apiRoute(route) {
    console.log(`${apiRoot}${route}`);
    return `${apiRoot}${route}`;
}
