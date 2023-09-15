import TabFrame from '../../components/TabFrame'
import {Outlet, useLocation, useParams} from 'react-router-dom';
import {routeMap} from "../../config/RouteConfig";
import React, {useEffect, useState} from "react";
import {findRouteIndex} from "../../util/RouteUtil";
import {updatePageData} from "../../util/UpdatePageData";
import PageLoadingAnimation from "../../components/PageLoadingAnimation";

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

    useEffect(() => {
        const fullPath = location.pathname + location.search
        if (findRouteIndex(routesToFetch, location.pathname) !== -1) {
            updatePageData(fullPath, setData, setLoadingCount)
        }
    }, [location]);

    let pageContent =  <PageLoadingAnimation />
    if (loadingCount === 0 && data) {
        pageContent = <Outlet data={data} />
    }
	return (
        <div className="flex justify-center">
            <div className="w-full max-w-7xl">
                <div className="w-full px-4">
                    <TabFrame routes={routes} routePatterns={routePatterns} routeLabels={routeLabels}>
                        <div className="w-full">
                            {pageContent}
                        </div>
                    </TabFrame>
                </div>
            </div>
        </div>
	);
}

export default Profile;