import {Route, Routes, useLocation} from "react-router-dom";
import React, {useEffect, useState} from "react";
import Menubar from "./components/Menubar";
import Main from "./pages/Main";
import Contests from "./pages/Contests";
import Info from "./pages/Info";
import Archive from "./pages/Archive";
import Submissions from "./pages/Submissions";
import Problems from "./pages/Problems";
import Submission from "./pages/Submission";
import Profile from "./pages/Profile/Profile";
import ProfileMain from "./pages/Profile/ProfileMain";
import ProfileSubmissions from "./pages/Profile/ProfileSubmissions";
import ProfileSettings from "./pages/Profile/ProfileSettings";
import Problem from "./pages/Problem/Problem";
import ProblemStatement from "./pages/Problem/ProblemStatement";
import ProblemSubmit from "./pages/Problem/ProblemSubmit";
import ProblemSubmissions from "./pages/Problem/ProblemSubmissions";
import ProblemRanklist from "./pages/Problem/ProblemRanklist";
import Login from "./pages/Auth/Login";
import Register from "./pages/Auth/Register";
import NotFound from "./pages/Error/NotFound";
import PageLoadingAnimation from "./components/PageLoadingAnimation";
import {updatePageData} from "./util/UpdatePageData";
import FadeIn from "./components/FadeIn";
import {routeMap} from "./config/RouteConfig";

const routesToFetch = [
    routeMap.main,
    routeMap.archive,
    routeMap.submissions,
    routeMap.problems,
    routeMap.submission,
    routeMap.login,
    routeMap.register
]

function RoutingComponent() {
    const location = useLocation()
    const [data, setData] = useState(null)
    const [loadingCount, setLoadingCount] = useState(0)
    const abortController = new AbortController();

    useEffect(() => {
        let isMounted = true
        updatePageData(location, routesToFetch, abortController, setData, setLoadingCount, () => isMounted)

        return () => {
            isMounted = false
            abortController.abort()
        }
    }, [location]);

    let pageContent = null
    if (loadingCount === 0)
        pageContent =
            <Routes location={location} key={location.pathname}>
                <Route path={routeMap.main} element={<FadeIn>
                    <Main data={data} />
                </FadeIn>}  />
                <Route path={routeMap.contests} element={<FadeIn>
                    <Contests data={data} />
                </FadeIn>} />
                <Route path={routeMap.info} element={<FadeIn>
                    <Info data={data} />
                </FadeIn>} />
                <Route path={routeMap.archive} element={<FadeIn>
                    <Archive data={data} />
                </FadeIn>} />
                <Route path={routeMap.submissions} element={<FadeIn>
                    <Submissions data={data} />
                </FadeIn>} />
                <Route path={routeMap.problems} element={<FadeIn>
                    <Problems data={data} />
                </FadeIn>} />
                <Route path={routeMap.submission} element={<FadeIn>
                    <Submission data={data} />
                </FadeIn>} />
                <Route path={routeMap.login} element={<FadeIn>
                    <Login data={data} />
                </FadeIn>} />
                <Route path={routeMap.register} element={<FadeIn>
                    <Register data={data} />
                </FadeIn>} />
                <Route path={routeMap.profile} element={<FadeIn>
                    <Profile />
                </FadeIn>} >
                    <Route path={routeMap.profile} element={<FadeIn>
                        <ProfileMain />
                    </FadeIn>} />
                    <Route path={routeMap.profileSubmissions} element={<FadeIn>
                        <ProfileSubmissions />
                    </FadeIn>} />
                    <Route path={routeMap.profileSettings} element={<FadeIn>
                        <ProfileSettings />
                    </FadeIn>} />
                </Route>
                <Route path={routeMap.problem} element={<FadeIn>
                    <Problem />
                </FadeIn>} >
                    <Route path={routeMap.problem} element={<FadeIn>
                        <ProblemStatement />
                    </FadeIn>} />
                    <Route path={routeMap.problemSubmit} element={<FadeIn>
                        <ProblemSubmit />
                    </FadeIn>} />
                    <Route path={routeMap.problemSubmissions} element={<FadeIn>
                        <ProblemSubmissions />
                    </FadeIn>} />
                    <Route path={routeMap.problemRanklist} element={<FadeIn>
                        <ProblemRanklist />
                    </FadeIn>} />
                </Route>
                <Route path="*" element={<FadeIn>
                    <NotFound />
                </FadeIn>} />
            </Routes>

    return (
        <>
            <div className="pb-20">
                <Menubar />
            </div>
            <div className="relative w-full">
                <PageLoadingAnimation isVisible={loadingCount !== 0} />
                {pageContent}
            </div>
        </>
    )
}

export default RoutingComponent