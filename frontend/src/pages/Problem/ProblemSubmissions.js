import SubmissionsTable from "../../components/SubmissionsTable";
import Checkbox from "../../components/Checkbox"
import RoundedFrame from "../../components/RoundedFrame";
import Pagination from "../../components/Pagination";
import {useParams} from "react-router-dom";
import {useEffect, useState} from "react";
import PageLoadingAnimation from "../../components/PageLoadingAnimation";
import ProfileSideBar from "../../components/ProfileSidebar";

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

function ProblemSubmissions() {
    const {problem} = useParams()
    const [data, setData] = useState(null)

    useEffect(() => {
        fetch(`/api/v2/problemset/main/${problem}/submissions`)
            .then(res => res.json())
            .then(data => setData(data))
    }, []);
    let pageContent = <PageLoadingAnimation/>;
    if (data) {
        pageContent =
            <>
                <div className="mb-3">
                    <SubmissionFilterFrame />
                </div>
                <div className="mb-2">
                    <SubmissionsTable submissions={data.submissions} />
                </div>
                <Pagination current={1000} last={2000} />
            </>
    }
    return (
        <div className="relative">
            {pageContent}
        </div>
    )
}

export default ProblemSubmissions;