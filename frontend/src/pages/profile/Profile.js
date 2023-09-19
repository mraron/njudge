import TabFrame from '../../components/container/TabFrame'
import {matchPath, Outlet, useLocation, useParams} from 'react-router-dom';
import React, {useContext, useEffect, useLayoutEffect, useState} from "react";
import {updatePageData} from "../../util/updateData";
import FadeIn from "../../components/util/FadeIn";
import {routeMap} from "../../config/RouteConfig";
import PageLoadingAnimation from "../../components/util/PageLoadingAnimation";
import UserContext from "../../contexts/user/UserContext";
import {findRouteIndex} from "../../util/findRouteIndex";

const routesToFetch = [
    routeMap.profile,
    routeMap.profileSubmissions
]

function Profile() {
    let routeLabels = [
        "Profil",
        "Beküldések",
        "Beállítások"
    ]
    let routePatterns = [
        routeMap.profile,
        routeMap.profileSubmissions,
        routeMap.profileSettings
    ]
    const {user} = useParams()
    const {userData, isLoggedIn} = useContext(UserContext)
    const location = useLocation()
    const [data, setData] = useState(null)
    const [isLoading, setLoading] = useState(false)
    const routes = routePatterns.map(item => item.replace(":user", user))
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

    if (!isLoggedIn || userData.username !== user) {
        routeLabels.pop()
        routePatterns.pop()
    }
    let pageContent = null
    if (findRouteIndex(routesToFetch, location.pathname) === -1 || !isLoading && data && matchPath(data.route, location.pathname)) {
        pageContent = <FadeIn><Outlet context={data}/></FadeIn>
    }
    return (
        <div className="flex justify-center">
            <div className="w-full max-w-7xl">
                <div className="w-full px-4">
                    <TabFrame routes={routes} routePatterns={routePatterns} routeLabels={routeLabels}>
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

export default Profile;