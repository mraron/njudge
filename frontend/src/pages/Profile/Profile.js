import TabFrame from '../../components/TabFrame'
import { Outlet } from 'react-router-dom';

const routes = [
    ["Profil", "/user/profile/"], 
    ["Beküldések", "/user/profile/submissions/"],
    ["Beállítások", "/user/profile/settings/"]];

function Profile() {
	return (
        <div className="text-white">
            <div className="w-full flex justify-center">
                <div className="flex justify-center w-full max-w-7xl">
                    <div className="w-full px-4">
                        <TabFrame routes={routes}>
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