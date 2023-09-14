import Pagination from '../components/Pagination';
import ProfileSideBar from '../components/ProfileSidebar'
import SubmissionsTable from '../components/SubmissionsTable';
import {matchPath} from "react-router-dom";
import {routeMap} from "../config/RouteConfig";

function Submissions({ data }) {
    if (!data || !matchPath(routeMap.submissions, data.route)) {
        return <></>
    }
    return (
        <div className="relative w-full flex justify-center">
            <div className="flex justify-center w-full max-w-7xl">
                <div className="ml-0 lg:ml-4">
                    <ProfileSideBar
                        src="https://st3.depositphotos.com/6672868/13701/v/450/depositphotos_137014128-stock-illustration-user-profile-icon.jpg"
                        username="dbence"
                        score="2550"/>
                </div>
                <div className="w-full px-4 lg:pl-3 overflow-x-auto">
                    <div className="mb-2">
                        <SubmissionsTable submissions={data.submissions}/>
                    </div>
                    <Pagination paginationData={data.paginationData}/>
                </div>
            </div>
        </div>
    );
}

export default Submissions;