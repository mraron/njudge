import React from "react"
import { Link } from "react-router-dom"
import { routeMap } from "../../../config/RouteConfig"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import RoundedTable from "../../container/RoundedTable"

function RanklistRow({ result, maxScore, index }) {
    const { username, score, submissionID, accepted } = result
    return (
        <tr
            className={`divide-x divide-dividecol ${
                index % 2 === 0 ? "bg-grey-850" : "bg-grey-825"
            }`}>
            <td className="py-2.5">
                <Link
                    className="link"
                    to={routeMap.profile.replace(":user", username)}>
                    {username}
                </Link>
            </td>
            <td className="py-2.5 w-0 sm:w-44">
                <div className="flex items-center justify-center">
                    {accepted && (
                        <FontAwesomeIcon
                            icon="fa-check"
                            className="w-3.5 h-3.5 highlight-green mr-2"
                        />
                    )}
                    {!accepted && (
                        <FontAwesomeIcon
                            icon="fa-xmark"
                            className="w-3.5 h-3.5 highlight-red mr-2"
                        />
                    )}
                    <Link
                        className="link whitespace-nowrap"
                        to={routeMap.submission.replace(":id", submissionID)}>
                        {score} / {maxScore}
                    </Link>
                </div>
            </td>
        </tr>
    )
}

function Ranklist({ ranklist }) {
    const rows = ranklist.results.map((item, index) => (
        <RanklistRow
            result={item}
            maxScore={ranklist.maxScore}
            index={index}
            key={index}
        />
    ))
    return (
        <RoundedTable cls="overflow-hidden">
            <thead className="bg-grey-800">
                <tr className="divide-x divide-dividecol">
                    <th>Név</th>
                    <th>Eredmény</th>
                </tr>
            </thead>
            <tbody className="divide-y divide-dividecol bg-grey-850">
                {rows}
            </tbody>
        </RoundedTable>
    )
}

export default Ranklist
