import { Link } from "react-router-dom"
import { useTranslation } from "react-i18next"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import RoundedTable from "../../container/RoundedTable"
import { SVGSpinner } from "../../svg/SVGs"
import { routeMap } from "../../../config/RouteConfig"

function Submission({ submission }) {
    const {
        id,
        date,
        user,
        problem,
        language,
        verdictName,
        verdictType,
        score,
        maxScore,
        time,
        memory,
    } = submission
    return (
        <tr>
            <td className=" w-0 text-center">
                <Link
                    className="link"
                    to={routeMap.submission.replace(":id", submission.id)}>
                    {id}
                </Link>
            </td>
            <td>{date}</td>
            <td>
                <Link
                    className="link"
                    to={routeMap.profile.replace(":user", submission.user)}>
                    {user}
                </Link>
            </td>
            <td>
                <Link className="link" to={problem.href}>
                    {problem.text}
                </Link>
            </td>
            <td>{language}</td>
            <td colSpan={maxScore === 0.0 ? 2 : 1}>
                <div className="flex items-center">
                    {verdictType === 0 && <SVGSpinner cls="w-4 h-4 mr-3" />}
                    {verdictType === 1 && (
                        <FontAwesomeIcon
                            icon="fa-xmark"
                            className="w-4 h-4 highlight-red mr-3"
                        />
                    )}
                    {verdictType === 2 && (
                        <FontAwesomeIcon
                            icon="fa-check"
                            className="w-4 h-4 highlight-yellow mr-3"
                        />
                    )}
                    {verdictType === 3 && (
                        <FontAwesomeIcon
                            icon="fa-check"
                            className="w-4 h-4 highlight-green mr-3"
                        />
                    )}
                    <span className="whitespace-nowrap">{verdictName}</span>
                </div>
            </td>
            {maxScore !== 0.0 && (
                <td className="w-0 text-center">
                    <span className="whitespace-nowrap">
                        {score} / {maxScore}
                    </span>
                </td>
            )}
            <td>{time} ms</td>
            <td>{memory} KiB</td>
        </tr>
    )
}

function SubmissionsTable({ submissions }) {
    const { t } = useTranslation()
    const submissionsContent = submissions.map((item, index) => (
        <Submission submission={item} key={index} />
    ))
    return (
        <RoundedTable>
            <thead>
                <tr>
                    <th>{t("submissions_table.id")}</th>
                    <th>{t("submissions_table.date")}</th>
                    <th>{t("submissions_table.user")}</th>
                    <th>{t("submissions_table.problem")}</th>
                    <th>{t("submissions_table.language")}</th>
                    <th colSpan={2}>{t("submissions_table.verdict")}</th>
                    <th>{t("submissions_table.time")}</th>
                    <th>{t("submissions_table.memory")}</th>
                </tr>
            </thead>
            <tbody>{submissionsContent}</tbody>
        </RoundedTable>
    )
}

export default SubmissionsTable
