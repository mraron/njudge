import Pagination from "../../components/Pagination";
import SubmissionsTable from "../../components/SubmissionsTable";
import {useOutletContext} from "react-router-dom";
import React from "react";
import checkData from "../../util/CheckData";

function ProfileSubmissions() {
    const data = useOutletContext()
    if (!checkData(data)) {
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