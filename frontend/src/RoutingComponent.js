import { useEffect } from "react"
import { Route, Routes, useLocation } from "react-router-dom"
import { useTranslation } from "react-i18next"

import Home from "./pages/Home"
import Contests from "./pages/Contests"
import Info from "./pages/Info"
import Archive from "./pages/Archive"
import Submissions from "./pages/Submissions"
import Problems from "./pages/Problems"
import Submission from "./pages/Submission"
import Profile from "./pages/profile/Profile"
import ProfileMain from "./pages/profile/ProfileMain"
import ProfileSubmissions from "./pages/profile/ProfileSubmissions"
import ProfileSettings from "./pages/profile/ProfileSettings"
import Problem from "./pages/problem/Problem"
import ProblemStatement from "./pages/problem/ProblemStatement"
import ProblemSubmit from "./pages/problem/ProblemSubmit"
import ProblemSubmissions from "./pages/problem/ProblemSubmissions"
import ProblemRanklist from "./pages/problem/ProblemRanklist"
import Login from "./pages/auth/Login"
import Register from "./pages/auth/Register"
import NotFound from "./pages/error/NotFound"
import Logout from "./pages/auth/Logout"
import Verify from "./pages/auth/Verify"
import Admin from "./pages/auth/Admin"
import UpdatePage from "./pages/wrappers/UpdatedPage"

import { findRouteIndex } from "./util/findRouteIndex"
import { routeMap } from "./config/RouteConfig"
import extractParams from "./util/extractParams"
import ForgottenPassword from "./pages/auth/ForgottenPassword"
import ResetPassword from "./pages/auth/ResetPassword"
import Contest from "./pages/contest/Contest"
import ContestProblems from "./pages/contest/ContestProblems"
import ContestSubmissions from "./pages/contest/ContestSubmissions"
import ContestRanklist from "./pages/contest/ContestRanklist"

const titles = {
    [routeMap.home]: "home.page_title",
    [routeMap.contests]: "contests.page_title",
    [routeMap.archive]: "archive.page_title",
    [routeMap.submissions]: "submissions.page_title",
    [routeMap.problems]: "problems.page_title",
    [routeMap.info]: "info.page_title",
    [routeMap.login]: "login.page_title",
    [routeMap.register]: "register.page_title",
    [routeMap.forgotten_password]: "forgotten_password.page_title",
    [routeMap.reset_password]: "reset_password.page_title",
    [routeMap.profile]: "profile_main.page_title",
    [routeMap.profileSubmissions]: "profile_submissions.page_title",
    [routeMap.profileSettings]: "profile_settings.page_title",
    [routeMap.problem]: "problem_statement.page_title",
    [routeMap.problemSubmit]: "problem_submit.page_title",
    [routeMap.problemSubmissions]: "problem_submissions.page_title",
    [routeMap.problemRanklist]: "problem_ranklist.page_title",
    [routeMap.contestProblems]: "contest_problems.page_title",
    [routeMap.contestSubmissions]: "contest_submissions.page_title",
    [routeMap.contestRanklist]: "contest_ranklist.page_title",
    [routeMap.submission]: "submission.page_title",
}
const routes = Object.keys(titles)

function RoutingComponent() {
    const { i18n, t } = useTranslation()
    const location = useLocation()

    useEffect(() => {
        const routeIndex = findRouteIndex(routes, location.pathname)
        if (routeIndex !== -1) {
            const params = extractParams(location.pathname, routes[routeIndex])
            document.title = t(titles[routes[routeIndex]], { params })
        } else {
            document.title = "njudge"
        }
    }, [location, i18n.language])

    return (
        <div className="w-full">
            <Routes>
                <Route path={routeMap.home} element={<UpdatePage key={location.key} page={Home} />} />
                <Route path={routeMap.contests} element={<UpdatePage key={location.key} page={Contests} />} />
                <Route path={routeMap.info} element={<UpdatePage key={location.key} page={Info} />} />
                <Route path={routeMap.archive} element={<UpdatePage key={location.key} page={Archive} />} />
                <Route path={routeMap.submissions} element={<UpdatePage key={location.key} page={Submissions} />} />
                <Route path={routeMap.problems} element={<UpdatePage key={location.key} page={Problems} />} />
                <Route path={routeMap.submission} element={<UpdatePage key={location.key} page={Submission} />} />
                <Route path={routeMap.login} element={<UpdatePage key={location.key} page={Login} />} />
                <Route path={routeMap.register} element={<UpdatePage key={location.key} page={Register} />} />
                <Route
                    path={routeMap.forgotten_password}
                    element={<UpdatePage key={location.key} page={ForgottenPassword} />}
                />
                <Route
                    path={routeMap.reset_password}
                    element={<UpdatePage key={location.key} page={ResetPassword} />}
                />
                <Route path={routeMap.verify} element={<UpdatePage key={location.key} page={Verify} />} />
                <Route path={routeMap.logout} element={<UpdatePage key={location.key} page={Logout} />} />
                <Route path={routeMap.admin} element={<UpdatePage key={location.key} page={Admin} />} />
                <Route path={routeMap.profile} element={<Profile />}>
                    <Route index element={<UpdatePage key={location.key} page={ProfileMain} />} />
                    <Route
                        path={routeMap.profileSubmissions}
                        element={<UpdatePage key={location.key} page={ProfileSubmissions} />}
                    />
                    <Route
                        path={routeMap.profileSettings}
                        element={<UpdatePage key={location.key} page={ProfileSettings} />}
                    />
                </Route>
                <Route path={routeMap.problem} element={<Problem />}>
                    <Route index element={<UpdatePage key={location.key} page={ProblemStatement} />} />
                    <Route
                        path={routeMap.problemSubmit}
                        element={<UpdatePage key={location.key} page={ProblemSubmit} />}
                    />
                    <Route
                        path={routeMap.problemSubmissions}
                        element={<UpdatePage key={location.key} page={ProblemSubmissions} />}
                    />
                    <Route
                        path={routeMap.problemRanklist}
                        element={<UpdatePage key={location.key} page={ProblemRanklist} />}
                    />
                </Route>
                <Route path={routeMap.contest} element={<Contest />}>
                    <Route index element={<UpdatePage key={location.key} page={ContestProblems} />} />
                    <Route
                        path={routeMap.contestSubmissions}
                        element={<UpdatePage key={location.key} page={ContestSubmissions} />}
                    />
                    <Route
                        path={routeMap.contestRanklist}
                        element={<UpdatePage key={location.key} page={ContestRanklist} />}
                    />
                </Route>
                <Route path="*" element={<UpdatePage page={NotFound} />} />
            </Routes>
        </div>
    )
}

export default RoutingComponent
