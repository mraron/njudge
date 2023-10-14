import { useEffect } from "react"
import { Link, useNavigate } from "react-router-dom"
import { useTranslation } from "react-i18next"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { routeMap } from "../../config/RouteConfig"
import Pagination from "../../components/util/Pagination"
import RoundedTable from "../../components/container/RoundedTable"

function RanklistRow({ result, maxScore, index }) {
    const { username, score, submissionID, accepted } = result
    return (
        <tr className={`divide-x divide-divide-def ${index % 2 === 0 ? "bg-grey-850" : "bg-grey-825"}`}>
            <td className="py-3">
                <Link className="link" to={routeMap.profile.replace(":user", username)}>
                    {username}
                </Link>
            </td>
            <td className="py-3 w-0">
                <div className="flex items-center justify-center">
                    {accepted && <FontAwesomeIcon icon="fa-check" className="w-4 h-4 highlight-green" />}
                    {!accepted && <FontAwesomeIcon icon="fa-xmark" className="w-4 h-4 highlight-red" />}
                </div>
            </td>
            <td className="py-3 w-0 sm:w-32 text-center">
                <Link className="link whitespace-nowrap" to={routeMap.submission.replace(":id", submissionID)}>
                    {score} / {maxScore}
                </Link>
            </td>
        </tr>
    )
}

function Ranklist({ ranklist }) {
    const { t } = useTranslation()

    const rows = ranklist.results.map((item, index) => (
        <RanklistRow result={item} maxScore={ranklist.maxScore} index={index} key={index} />
    ))
    return (
        <RoundedTable cls="overflow-hidden">
            <thead>
                <tr>
                    <th>{t("problem_ranklist.username")}</th>
                    <th colSpan={2}>{t("problem_ranklist.result")}</th>
                </tr>
            </thead>
            <tbody className="bg-grey-850">{rows}</tbody>
        </RoundedTable>
    )
}

function ProblemRanklist({ data }) {
    const { t } = useTranslation()
    const navigate = useNavigate()

    useEffect(() => {
        if (!data.isPublic) {
            window.flash("flash.ranklist_not_public", "failure")
            navigate(-1)
        }
    })
    return (
        data.isPublic && (
            <div className="space-y-2">
                <Ranklist ranklist={data.ranklist} title={t("problem_ranklist.ranklist")} />
                <Pagination paginationData={data.paginationData} />
            </div>
        )
    )
}

export default ProblemRanklist
