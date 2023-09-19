import TabFrame from '../../components/container/TabFrame'
import {matchPath, Outlet, useLocation, useParams} from 'react-router-dom';
import React, {useEffect, useLayoutEffect, useState} from "react";
import {updatePageData} from "../../util/updateData";
import FadeIn from "../../components/util/FadeIn";
import PageLoadingAnimation from "../../components/util/PageLoadingAnimation";
import {routeMap} from "../../config/RouteConfig";
import {findRouteIndex} from "../../util/findRouteIndex";

const routeLabels = [
    "Leírás",
    "Beküld",
    "Beküldések",
    "Eredmények"
]
const routePatterns = [
    routeMap.problem,
    routeMap.problemSubmit,
    routeMap.problemSubmissions,
    routeMap.problemRanklist
]
const routesToFetch = [
    routeMap.problem,
    routeMap.problemSubmissions,
    routeMap.problemRanklist
]

function Problem() {
    const {problem} = useParams()
    const location = useLocation()
    const [data, setData] = useState(null)
    const [isLoading, setLoading] = useState(false)
    const routes = routePatterns.map(item => item.replace(":problem", problem))
    const abortController = new AbortController();

    useLayoutEffect(() => {
        setLoading(true)
    }, [location]);

    useEffect(() => {
        let isMounted = true
        updatePageData(location, routesToFetch, abortController, setData, () => isMounted).then(() =>
            setLoading(false)
        )
        return () => {
            isMounted = false
            abortController.abort()
        }
    }, [location]);

    let pageContent = null
    if (findRouteIndex(routesToFetch, location.pathname) === -1 || !isLoading && data && matchPath(data.route, location.pathname)) {
        pageContent = <FadeIn><Outlet context={data}/></FadeIn>
    }
    return (
        <div className="flex justify-center">
            <div className="w-full max-w-7xl">
                <div className="w-full px-4">
                    <TabFrame routes={routes} routeLabels={routeLabels} routePatterns={routePatterns}>
                        <div className="relative w-full">
                            <PageLoadingAnimation isVisible={isLoading}/>
                            {pageContent}
                        </div>
                    </TabFrame>
                </div>
            </div>
        </div>
    );
}

export default Problem;