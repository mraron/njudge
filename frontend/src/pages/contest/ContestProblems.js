import { Link } from "react-router-dom"
import { useTranslation } from "react-i18next"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { SVGEllipsis } from "../../components/svg/SVGs"
import RoundedTable from "../../components/container/RoundedTable"
import React from "react"

function Problem({ problem }) {
    const { id, title, status } = problem
    return (
        <tr>
            <td className="text-center">{id}</td>
            <td>
                <Link className="link" to={title.href}>
                    {title.text}
                </Link>
            </td>
            <td className="text-center">
                {status === 0 && <SVGEllipsis cls="w-4 h-4 text-grey-300" />}
                {status === 1 && <FontAwesomeIcon icon="fa-xmark" className="w-4 h-4 highlight-red" />}
                {status === 2 && <FontAwesomeIcon icon="fa-check" className="w-4 h-4 highlight-yellow" />}
                {status === 3 && <FontAwesomeIcon icon="fa-check" className="w-4 h-4 highlight-green" />}
            </td>
        </tr>
    )
}

function ContestProblems({ data }) {
    const { t } = useTranslation()
    const problemsContent = data.problems.map((item, index) => <Problem problem={item} key={index} />)
    return (
        <RoundedTable>
            <thead>
                <tr>
                    <th className="w-0">{t("contest_problems.id")}</th>
                    <th>{t("contest_problems.title")}</th>
                    <th className="w-0">{t("contest_problems.status")}</th>
                </tr>
            </thead>
            <tbody>{problemsContent}</tbody>
        </RoundedTable>
    )
}

export default ContestProblems
