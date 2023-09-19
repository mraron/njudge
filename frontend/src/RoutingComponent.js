import {matchPath, Route, Routes, useLocation} from "react-router-dom";
import React, {useContext, useEffect, useLayoutEffect, useState} from "react";
import Main from "./pages/Main";
import Contests from "./pages/Contests";
import Info from "./pages/Info";
import Archive from "./pages/Archive";
import Submissions from "./pages/Submissions";
import Problems from "./pages/Problems";
import Submission from "./pages/Submission";
import Profile from "./pages/profile/Profile";
import ProfileMain from "./pages/profile/ProfileMain";
import ProfileSubmissions from "./pages/profile/ProfileSubmissions";
import ProfileSettings from "./pages/profile/ProfileSettings";
import Problem from "./pages/problem/Problem";
import ProblemStatement from "./pages/problem/ProblemStatement";
import ProblemSubmit from "./pages/problem/ProblemSubmit";
import ProblemSubmissions from "./pages/problem/ProblemSubmissions";
import ProblemRanklist from "./pages/problem/ProblemRanklist";
import Login from "./pages/auth/Login";
import Register from "./pages/auth/Register";
import NotFound from "./pages/error/NotFound";
import PageLoadingAnimation from "./components/util/PageLoadingAnimation";
import {updateData} from "./util/updateData";
import FadeIn from "./components/util/FadeIn";
import {routeMap} from "./config/RouteConfig";
import UserContext from "./contexts/user/UserContext";
import {findRouteIndex} from "./util/findRouteIndex";
import Logout from "./pages/auth/Logout";
import UpdatePage from "./pages/wrappers/UpdatedPage";

const routesToFetch = [
    routeMap.main,
    routeMap.contests,
    routeMap.archive,
    routeMap.submissions,
    routeMap.problems,
    routeMap.submission
]

function RoutingComponent() {
    const location = useLocation()
    return (
        <div className="w-full">
            <Routes>
                <Route path={routeMap.main} element={<UpdatePage key={location.key} page={Main} />}/>
                <Route path={routeMap.contests} element={<UpdatePage key={location.key} page={Contests} />}/>
                <Route path={routeMap.info} element={<UpdatePage key={location.key} page={Info} />}/>
                <Route path={routeMap.archive} element={<UpdatePage key={location.key} page={Archive} />}/>
                <Route path={routeMap.submissions} element={<UpdatePage key={location.key} page={Submissions} />}/>
                <Route path={routeMap.problems} element={<UpdatePage key={location.key} page={Problems} />}/>
                <Route path={routeMap.submission} element={<UpdatePage key={location.key} page={Submission} />}/>
                <Route path={routeMap.login} element={<UpdatePage key={location.key} page={Login} />}/>
                <Route path={routeMap.register} element={<UpdatePage key={location.key} page={Register} />}/>
                <Route path={routeMap.logout} element={<UpdatePage key={location.key} page={Logout} />}/>
                <Route path={routeMap.profile} element={<FadeIn><Profile /></FadeIn>}>
                    <Route index element={<UpdatePage key={location.key} page={ProfileMain} />}/>
                    <Route path={routeMap.profileSubmissions} element={<UpdatePage key={location.key} page={ProfileSubmissions}/>}/>
                    <Route path={routeMap.profileSettings} element={<UpdatePage key={location.key} page={ProfileSettings}/>}/>
                </Route>
                <Route path={routeMap.problem} element={<FadeIn><Problem /></FadeIn>}>
                    <Route index element={<UpdatePage key={location.key} page={ProblemStatement} />}/>
                    <Route path={routeMap.problemSubmit} element={<UpdatePage key={location.key} page={ProblemSubmit} />}/>
                    <Route path={routeMap.problemSubmissions} element={<UpdatePage key={location.key} page={ProblemSubmissions} />}/>
                    <Route path={routeMap.problemRanklist} element={<UpdatePage key={location.key} page={ProblemRanklist} />}/>
                </Route>
                <Route path="*" element={<NotFound/>}/>
            </Routes>
        </div>
    )
}

export default RoutingComponent