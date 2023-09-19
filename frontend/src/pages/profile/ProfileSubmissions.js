import Pagination from "../../components/util/Pagination";
import SubmissionsTable from "../../components/concrete/table/SubmissionsTable";
import React from "react";

function ProfileSubmissions({ data }) {
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