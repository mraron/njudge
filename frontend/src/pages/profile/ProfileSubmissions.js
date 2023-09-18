import Pagination from "../../components/util/Pagination";
import SubmissionsTable from "../../components/concrete/table/SubmissionsTable";
import {matchPath, useLocation, useOutletContext} from "react-router-dom";
import React from "react";

function ProfileSubmissions() {
    const data = useOutletContext()
    const location = useLocation()

    if (!data || !matchPath(data.route, location.pathname)) {
        return
    }
    return (
        <div className="relative">
            <div className="flex flex-col w-full">
                <div className="mb-2">
                    <SubmissionsTable submissions={data.submissions}/>
                </div>
                <Pagination paginationData={data.paginationData}/>
            </div>
        </div>
    );
}

export default ProfileSubmissions;