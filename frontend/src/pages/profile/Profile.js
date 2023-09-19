import React, {useContext} from "react";
import TabFrame from '../../components/container/TabFrame'
import {Outlet, useParams} from 'react-router-dom';
import {routeMap} from "../../config/RouteConfig";
import UserContext from "../../contexts/user/UserContext";

function Profile({ data }) {
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
    const routes = routePatterns.map(item => item.replace(":user", user))

    if (!isLoggedIn || userData.username !== user) {
        routeLabels.pop()
        routePatterns.pop()
    }
    return (
        <div className="flex justify-center">
            <div className="w-full max-w-7xl">
                <div className="w-full px-4">
                    <TabFrame routes={routes} routePatterns={routePatterns} routeLabels={routeLabels}>
                        <div className="relative w-full">
                            <Outlet/>
                        </div>
                    </TabFrame>
                </div>
            </div>
        </div>
    );
}

export default Profile;