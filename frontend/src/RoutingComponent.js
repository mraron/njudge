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
import ProblemRanklist from "./pages/Problem/ProblemRankings";
import Login from "./pages/Auth/Login";
import Register from "./pages/Auth/Register";
import NotFound from "./pages/Error/NotFound";
import PageLoadingAnimation from "./components/PageLoadingAnimation";
import {routeMap} from "./config/RouteConfig";

function RoutingComponent() {
    const location = useLocation()
    const [data, setData] = useState(null)
    const [loadingCount, setLoadingCount] = useState(0)

    useEffect(() => {
        setLoadingCount(arg => arg + 1)
        fetch(`/api/v2${location.pathname}${location.search}`)
            .then(res => res.json())
            .then(data => {
                setData(data)
                setLoadingCount(arg => arg - 1)
            })
            .catch((error) => {
                console.error('Network error:', error);
                setLoadingCount((count) => count - 1);
            });
    }, [location]);
    let pageContent = <PageLoadingAnimation />
    if (loadingCount === 0 && data) {
        pageContent =
            <Routes>
                <Route path={routeMap.main} element={<Main data={data} />}  />
                <Route path={routeMap.contests} element={<Contests data={data} />} />
                <Route path={routeMap.info} element={<Info data={data} />} />
                <Route path={routeMap.archive} element={<Archive data={data} />} />
                <Route path={routeMap.submissions} element={<Submissions data={data} />} />
                <Route path={routeMap.problems} element={<Problems data={data} />} />
                <Route path={routeMap.submission} element={<Submission data={data} />} />
                <Route path={routeMap.profile} element={<Profile data={data} />} >
                    <Route path={routeMap.profile} element={<ProfileMain data={data} />} />
                    <Route path={routeMap.profileSubmissions} element={<ProfileSubmissions data={data} />} />
                    <Route path={routeMap.profileSettings} element={<ProfileSettings data={data} />} />
                </Route>
                <Route path={routeMap.problem} element={<Problem data={data} />} >
                    <Route path={routeMap.problem} element={<ProblemStatement data={data} />} />
                    <Route path={routeMap.problemSubmit} element={<ProblemSubmit data={data} />} />
                    <Route path={routeMap.problemSubmissions} element={<ProblemSubmissions data={data} />} />
                    <Route path={routeMap.problemRanklist} element={<ProblemRanklist data={data} />} />
                </Route>
                <Route path={routeMap.login} element={<Login data={data} />} />
                <Route path={routeMap.register} element={<Register data={data} />} />
                <Route path="*" element={<NotFound data={data} />} />
            </Routes>
    }
    return (
        <>
            <div className="pb-20">
                <Menubar />
            </div>
            {pageContent}
        </>
    )
}

export default RoutingComponent