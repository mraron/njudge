import SubmissionsTable from "../../components/concrete/table/SubmissionsTable";
import Checkbox from "../../components/input/Checkbox"
import RoundedFrame from "../../components/container/RoundedFrame";
import Pagination from "../../components/util/Pagination";
import {useLocation, useNavigate} from "react-router-dom";
import UpdateQueryString from "../../util/updateQueryString";
import queryString from "query-string";
import {useTranslation} from "react-i18next";

function SubmissionFilterFrame() {
    const {t} = useTranslation()
    const location = useLocation()
    const navigate = useNavigate()
    const handleCheckboxFullClicked = (checked) => {
        UpdateQueryString(location, navigate, "ac", checked ? 1 : 0, ["ac", "own"])
    }
    const handleCheckboxOwnClicked = (checked) => {
        UpdateQueryString(location, navigate, "own", checked ? 1 : 0, ["ac", "own"])
    }
    const qData = queryString.parse(location.search)
    const accepted = qData["ac"] === "1"
    const own = qData["own"] === "1"
    return (
        <RoundedFrame>
            <div className="px-6 py-4 flex flex-col sm:flex-row items-start sm:items-center justify-between">
                <div className="mb-2 sm:mb-0">
                    <Checkbox label={t("problem_submissions.full_solutions")} onChange={handleCheckboxFullClicked} initChecked={accepted}/>
                </div>
                <Checkbox label={t("problem_submissions.own_solutions")} onChange={handleCheckboxOwnClicked} initChecked={own}/>
            </div>
        </RoundedFrame>
    )
}

function ProblemSubmissions({ data }) {
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