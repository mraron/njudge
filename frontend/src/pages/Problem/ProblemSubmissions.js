import SubmissionsTable from "../../components/concrete/table/SubmissionsTable";
import Checkbox from "../../components/input/Checkbox"
import RoundedFrame from "../../components/container/RoundedFrame";
import Pagination from "../../components/util/Pagination";
import {useLocation, useNavigate, useOutletContext} from "react-router-dom";
import checkData from "../../util/CheckData";
import UpdateQueryString from "../../util/UpdateQueryString";
import queryString from "query-string";

function SubmissionFilterFrame() {
    const location = useLocation()
    const navigate = useNavigate()
    const handleCheckboxFullClicked = (checked) => {
        UpdateQueryString(location, navigate, "ac", checked? 1: 0)
    }
    const handleCheckboxOwnClicked = (checked) => {
        UpdateQueryString(location, navigate, "own", checked? 1: 0)
    }
    const qData = queryString.parse(location.search)
    const accepted = qData["ac"] === "1"
    const own = qData["own"] === "1"
    return (
        <RoundedFrame>
            <div className="px-6 py-4 flex flex-col sm:flex-row items-start sm:items-center justify-between">
                <div className="mb-2 sm:mb-0">
                    <Checkbox label="Teljes megoldások" onChange={handleCheckboxFullClicked} initChecked={accepted} />
                </div>
                <Checkbox label="Saját beküldéseim" onChange={handleCheckboxOwnClicked} initChecked={own} />
            </div>
        </RoundedFrame>
    )
}

function ProblemSubmissions() {
    const data = useOutletContext()
    if (!checkData(data)) {
        return
    }
    return (
        <div className="relative">
            <div className="mb-2">
                <SubmissionFilterFrame/>
            </div>
            <div className="mb-2">
                <SubmissionsTable submissions={data.submissions}/>
            </div>
            <Pagination paginationData={data.paginationData}/>
        </div>
    )
}

export default ProblemSubmissions;