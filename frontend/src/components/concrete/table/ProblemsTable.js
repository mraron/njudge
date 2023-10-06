import React, { useContext } from "react"
import { useTranslation } from "react-i18next"
import { Link } from "react-router-dom"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { SVGEllipsis } from "../../svg/SVGs"
import RoundedTable from "../../container/RoundedTable"
import UserContext from "../../../contexts/user/UserContext"
import { routeMap } from "../../../config/RouteConfig"
import OrderedColumnTitle from "../../util/OrderedColumnTitle"
import Tag from "../../basic/Tag"

function Problem(data) {
    const { t } = useTranslation()
    const { isLoggedIn } = useContext(UserContext)
    const { problem, solvedStatus, title, category, tags, solverCount } = data.problem
    const tagsContent = tags.map((item, index) => <Tag key={index}>{t(item)}</Tag>)
    return (
        <tr>
            {isLoggedIn && (
                <td className=" w-0">
                    <div className="flex items-center justify-center">
                        {solvedStatus === 0 && (
                            <SVGEllipsis cls="w-5 h-5 text-grey-300" title={t("solved_status.not_tried")} />
                        )}
                        {solvedStatus === 1 && (
                            <FontAwesomeIcon
                                icon="fa-xmark"
                                className="w-4 h-4 highlight-red"
                                title={t("solved_status.wrong")}
                            />
                        )}
                        {solvedStatus === 2 && (
                            <FontAwesomeIcon
                                icon="fa-check"
                                className="w-4 h-4 highlight-yellow"
                                title={t("solved_status.partially_correct")}
                            />
                        )}
                        {solvedStatus === 3 && (
                            <FontAwesomeIcon
                                icon="fa-check"
                                className="w-4 h-4 highlight-green"
                                title={t("solved_status.correct")}
                            />
                        )}
                    </div>
                </td>
            )}
            <td>{problem}</td>
            <td>
                <Link
                    className="link"
                    to={routeMap.problem.replace(":problemset", data.problemset).replace(":problem", problem)}>
                    {title}
                </Link>
            </td>
            <td>
                <Link className="link" to={category.href}>
                    {category.text}
                </Link>
            </td>
            <td>
                <div className="flex flex-wrap -m-1.5">{tagsContent}</div>
            </td>
            <td className=" w-0">
                <Link
                    className="link flex items-center justify-center"
                    to={`${routeMap.problemSubmissions
                        .replace(":problemset", data.problemset)
                        .replace(":problem", problem)}?ac=1`}>
                    <FontAwesomeIcon icon="fa-user" className="w-3.5 h-3.5 mr-1" />
                    <span>{solverCount}</span>
                </Link>
            </td>
        </tr>
    )
}

function ProblemsTable({ problemset = "main", problems }) {
    const { t } = useTranslation()
    const { isLoggedIn } = useContext(UserContext)

    const problemsContent = problems.map((item, index) => (
        <Problem problemset={problemset} problem={item} key={index} />
    ))
    return (
        <RoundedTable>
            <thead>
                <tr>
                    <th colSpan={isLoggedIn ? 2 : 1}>{t("problems_table.id")}</th>
                    <th>{t("problems_table.title")}</th>
                    <th>{t("problems_table.category")}</th>
                    <th>{t("problems_table.tags")}</th>
                    <th>
                        <OrderedColumnTitle text={t("problems_table.solved")} label="solver_count" />
                    </th>
                </tr>
            </thead>
            <tbody>{problemsContent}</tbody>
        </RoundedTable>
    )
}

export default ProblemsTable
