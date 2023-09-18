import TabFrame from '../../components/container/TabFrame'
import {Outlet, useLocation, useParams} from 'react-router-dom';
import React, {useContext, useEffect, useState} from "react";
import {updateData, updatePageData} from "../../util/UpdateData";
import FadeIn from "../../components/util/FadeIn";
import {routeMap} from "../../config/RouteConfig";
import PageLoadingAnimation from "../../components/util/PageLoadingAnimation";
import UserContext from "../../contexts/user/UserContext";

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
    const routesToFetch = [
        routeMap.profile,
        routeMap.profileSubmissions
    ]
    const {user} = useParams()
    const {userData, isLoggedIn} = useContext(UserContext)
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

    if (!isLoggedIn || userData.username !== user) {
        routeLabels.pop()
        routePatterns.pop()
    }
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