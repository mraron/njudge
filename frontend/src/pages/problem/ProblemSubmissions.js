import React, {useContext, useState} from "react";
import SubmissionsTable from "../../components/concrete/table/SubmissionsTable";
import Checkbox from "../../components/input/Checkbox"
import Pagination from "../../components/util/Pagination";
import {useLocation, useNavigate} from "react-router-dom";
import UpdateQueryString from "../../util/updateQueryString";
import {useTranslation} from "react-i18next";
import DropdownFrame from "../../components/container/DropdownFrame";
import queryString from "query-string";
import UserContext from "../../contexts/user/UserContext";

function SubmissionFilterFrame() {
    const {t} = useTranslation()
    const {isLoggedIn} = useContext(UserContext)
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
        <DropdownFrame title={t("problem_submissions.filter")}>
            <div className="px-8 py-6 flex flex-col">
                {isLoggedIn && <div className="mb-3">
                    <Checkbox label={t("problem_submissions.own_solutions")} onChange={setOnlyOwn} initChecked={onlyOwn}/>
                </div>}
                <div className="mb-5">
                    <Checkbox label={t("problem_submissions.full_solutions")} onChange={setOnlyAC} initChecked={onlyAC}/>
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