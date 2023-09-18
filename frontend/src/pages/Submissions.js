import Pagination from '../components/util/Pagination';
import ProfileSideBar from '../components/concrete/other/ProfileSidebar'
import SubmissionsTable from '../components/concrete/table/SubmissionsTable';
import React from "react";
import checkData from "../util/CheckData";
import {matchPath, useLocation} from "react-router-dom";

function Submissions({data}) {
    return (
        <div className="relative w-full flex justify-center">
            <div className="flex justify-center w-full max-w-7xl">
                <div className="ml-0 lg:ml-4">
                    <ProfileSideBar/>
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