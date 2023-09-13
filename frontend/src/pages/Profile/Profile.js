import TabFrame from '../../components/TabFrame'
import {Outlet, useParams} from 'react-router-dom';
import {routeMap} from "../../config/RouteConfig";

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

function Profile() {
    const {user} = useParams()
    const routes = routePatterns.map(item => item.replace(":user", user))
	return (
        <div className="text-white">
            <div className="w-full flex justify-center">
                <div className="flex justify-center w-full max-w-7xl">
                    <div className="w-full px-4">
                        <TabFrame routes={routes} routePatterns={routePatterns} routeLabels={routeLabels}>
                            <div className="w-full">
                                <Outlet />
                            </div>
                        </TabFrame>
                    </div>
                </div>
            </div>
        </div>
	);
}

export default Profile;