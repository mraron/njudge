import TabFrame from '../../components/container/TabFrame'
import {Outlet, useLocation, useParams} from 'react-router-dom';
import React, {useEffect, useState} from "react";
import {updatePageData} from "../../util/UpdatePageData";
import FadeIn from "../../components/util/FadeIn";
import PageLoadingAnimation from "../../components/util/PageLoadingAnimation";
import {routeMap} from "../../config/RouteConfig";

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
const routesToFetch = routePatterns

function Problem() {
    const {problem} = useParams()
    const location = useLocation()
    const [data, setData] = useState(null)
    const [loadingCount, setLoadingCount] = useState(0)
    const routes = routePatterns.map(item => item.replace(":problem", problem))
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
    if (loadingCount === 0) {
        pageContent = <FadeIn><Outlet context={data}/></FadeIn>
    }
    return (
        <div className="flex justify-center">
            <div className="w-full max-w-7xl">
                <div className="w-full px-4">
                    <TabFrame routes={routes} routeLabels={routeLabels} routePatterns={routePatterns}>
                        <div className="relative w-full">
                            <PageLoadingAnimation isVisible={loadingCount !== 0}/>
                            {pageContent}
                        </div>
                    </TabFrame>
                </div>
            </div>
        </div>
    );
}

export default Problem;