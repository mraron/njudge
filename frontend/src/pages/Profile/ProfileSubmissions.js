import Pagination from "../../components/Pagination";
import SubmissionsTable from "../../components/SubmissionsTable";
import {matchPath} from "react-router-dom";
import {routeMap} from "../../config/RouteConfig";

function ProfileSubmissions({ data }) {
    if (!data || !matchPath(routeMap.profileSubmissions, data.route)) {
        return <></>
    }
    return (
        <div className="relative">
            <div className="flex flex-col w-full">
                <div className="mb-2">
                    <SubmissionsTable submissions={data.submissions} />
                </div>
                <Pagination paginationData={data.paginationData} />
            </div>
        </div>
    );
}

export default ProfileSubmissions;