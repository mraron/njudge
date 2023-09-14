import SubmissionsTable from "../../components/SubmissionsTable";
import Checkbox from "../../components/Checkbox"
import RoundedFrame from "../../components/RoundedFrame";
import Pagination from "../../components/Pagination";
import {matchPath, useLocation, useParams} from "react-router-dom";
import {useEffect, useState} from "react";
import PageLoadingAnimation from "../../components/PageLoadingAnimation";
import ProfileSideBar from "../../components/ProfileSidebar";
import {updatePageData} from "../../util/UpdatePageData";
import {routeMap} from "../../config/RouteConfig";

function SubmissionFilterFrame() {
    return (
        <RoundedFrame>
            <div className="px-6 py-4 flex flex-col sm:flex-row items-start sm:items-center justify-between">
                <div className="mb-2 sm:mb-0">
                    <Checkbox label="Teljes megoldások" />
                </div>
                <Checkbox label="Saját beküldéseim" />
            </div>
        </RoundedFrame>
    )
}

function ProblemSubmissions({ data }) {
    if (!data || !matchPath(routeMap.problemSubmissions, data.route)) {
        return <></>
    }
    return (
        <div className="relative">
            <div className="mb-2">
                <SubmissionFilterFrame />
            </div>
            <div className="mb-2">
                <SubmissionsTable submissions={data.submissions} />
            </div>
            <Pagination paginationData={data.paginationData} />
        </div>
    )
}

export default ProblemSubmissions;