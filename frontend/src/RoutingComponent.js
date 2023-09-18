import {Route, Routes, useLocation} from "react-router-dom";
import React, {useContext, useEffect, useState} from "react";
import Menubar from "./components/concrete/other/Menubar";
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
import {updateData} from "./util/UpdateData";
import FadeIn from "./components/util/FadeIn";
import {routeMap} from "./config/RouteConfig";
import UserContext from "./contexts/user/UserContext";
import {authenticate} from "./util/User";

const routesToFetch = [
    routeMap.main,
    routeMap.info,
    routeMap.contests,
    routeMap.archive,
    routeMap.submissions,
    routeMap.problems,
    routeMap.submission
]

function RoutingComponent() {
    const {setUserData, setLoggedIn} = useContext(UserContext)
    const location = useLocation()
    const [data, setData] = useState(null)
    const [loadingCount, setLoadingCount] = useState(0)
    const abortController = new AbortController();

    useEffect(() => {
        let isMounted = true
        updateData(
            location,
            routesToFetch,
            abortController,
            setData,
            setLoadingCount,
            setUserData,
            setLoggedIn,
            () => isMounted
        )
        return () => {
            isMounted = false
            abortController.abort()
        }
    }, [location]);

    let pageContent = null
    if (loadingCount === 0)
        pageContent =
            <Routes key={location.pathname}>
                <Route path={routeMap.main} element={<FadeIn>
                    <Main data={data}/>
                </FadeIn>}/>
                <Route path={routeMap.contests} element={<FadeIn>
                    <Contests data={data}/>
                </FadeIn>}/>
                <Route path={routeMap.info} element={<FadeIn>
                    <Info data={data}/>
                </FadeIn>}/>
                <Route path={routeMap.archive} element={<FadeIn>
                    <Archive data={data}/>
                </FadeIn>}/>
                <Route path={routeMap.submissions} element={<FadeIn>
                    <Submissions data={data}/>
                </FadeIn>}/>
                <Route path={routeMap.problems} element={<FadeIn>
                    <Problems data={data}/>
                </FadeIn>}/>
                <Route path={routeMap.submission} element={<FadeIn>
                    <Submission data={data}/>
                </FadeIn>}/>
                <Route path={routeMap.login} element={<FadeIn>
                    <Login data={data}/>
                </FadeIn>}/>
                <Route path={routeMap.register} element={<FadeIn>
                    <Register data={data}/>
                </FadeIn>}/>
                <Route path={routeMap.profile} element={<Profile/>}>
                    <Route path={routeMap.profile} element={<ProfileMain/>}/>
                    <Route path={routeMap.profileSubmissions} element={<ProfileSubmissions/>}/>
                    <Route path={routeMap.profileSettings} element={<ProfileSettings/>}/>
                </Route>
                <Route path={routeMap.problem} element={<Problem/>}>
                    <Route path={routeMap.problem} element={<ProblemStatement/>}/>
                    <Route path={routeMap.problemSubmit} element={<ProblemSubmit/>}/>
                    <Route path={routeMap.problemSubmissions} element={<ProblemSubmissions/>}/>
                    <Route path={routeMap.problemRanklist} element={<ProblemRanklist/>}/>
                </Route>
                <Route path="*" element={<FadeIn>
                    <NotFound/>
                </FadeIn>}/>
            </Routes>

    return (
        <div className="relative w-full">
            <PageLoadingAnimation isVisible={loadingCount !== 0}/>
            {pageContent}
        </div>
    )
}

export default RoutingComponent