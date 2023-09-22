import SubmissionsTable from "../../components/concrete/table/SubmissionsTable";
import Checkbox from "../../components/input/Checkbox"
import RoundedFrame from "../../components/container/RoundedFrame";
import Pagination from "../../components/util/Pagination";
import {useLocation, useNavigate} from "react-router-dom";
import UpdateQueryString from "../../util/updateQueryString";
import queryString from "query-string";
import {useTranslation} from "react-i18next";
import DropdownFrame from "../../components/container/DropdownFrame";
import React, {useState} from "react";

function SubmissionFilterFrame() {
    const {t} = useTranslation()
    const location = useLocation()
    const navigate = useNavigate()

    const qData = queryString.parse(location.search)
    const [onlyAC, setOnlyAC] = useState(qData["ac"] === "1")
    const [onlyOwn, setOnlyOwn] = useState(qData["own"] === "1")

    const handleReset = () => {
        navigate(location.pathname)
    }
    const handleSubmit = () => {
        UpdateQueryString(location, navigate, ["ac", "own"], [onlyAC? 1: 0, onlyOwn? 1: 0], ["ac", "own"])
    }
    return (
        <DropdownFrame title="Szűrés">
            <div className="px-8 py-6 flex flex-col">
                <div className="mb-3">
                    <Checkbox label={t("problem_submissions.full_solutions")} onChange={setOnlyAC} initChecked={onlyAC}/>
                </div>
                <div className="mb-5">
                    <Checkbox label={t("problem_submissions.own_solutions")} onChange={setOnlyOwn} initChecked={onlyOwn}/>
                </div>
                <div className="flex justify-center">
                    <button className="mr-1 btn-indigo padding-btn-default w-32" onClick={handleSubmit}>{t("problem_filter.search")}</button>
                    <button className="ml-1 btn-gray padding-btn-default w-32" onClick={handleReset}>{t("problem_filter.reset")}</button>
                </div>
            </div>
        </DropdownFrame>
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