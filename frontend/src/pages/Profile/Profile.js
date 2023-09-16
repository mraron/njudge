import TabFrame from '../../components/container/TabFrame'
import {Outlet, useLocation, useParams} from 'react-router-dom';
import React, {useEffect, useState} from "react";
import {updatePageData} from "../../util/UpdatePageData";
import FadeIn from "../../components/util/FadeIn";
import {routeMap} from "../../config/RouteConfig";
import PageLoadingAnimation from "../../components/util/PageLoadingAnimation";

const routeLabels = [
    "Profil",
    "Beküldések",
    "Beállítások"
]
const routePatterns = [
    routeMap.profile,
    routeMap.profileSubmissions,
    routeMap.profileSettings
]
const routesToFetch = routePatterns

function Profile() {
    const {user} = useParams()
    const location = useLocation()
    const [data, setData] = useState(null)
    const [loadingCount, setLoadingCount] = useState(0)
    const routes = routePatterns.map(item => item.replace(":user", user))
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
                    <TabFrame routes={routes} routePatterns={routePatterns} routeLabels={routeLabels}>
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

export default Profile;