import Pagination from "../../components/Pagination";
import SubmissionsTable from "../../components/SubmissionsTable";
import {matchPath, useOutletContext} from "react-router-dom";
import {routeMap} from "../../config/RouteConfig";
import React from "react";

function ProfileSubmissions() {
    const data = useOutletContext()
    if (!data || data.processed) {
        return <></>
    }
    data.processed = true
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