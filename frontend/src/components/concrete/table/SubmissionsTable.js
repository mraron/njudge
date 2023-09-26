import RoundedTable from "../../container/RoundedTable";
import { SVGSpinner } from "../../svg/SVGs";
import { Link } from "react-router-dom";
import { routeMap } from "../../../config/RouteConfig";
import { useTranslation } from "react-i18next";
import React from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

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
    } = submission;
    console.log(problem);
    return (
        <tr className={"divide-x divide-dividecol"}>
            <td className="padding-td-default w-0">
                <Link
                    className="link"
                    to={routeMap.submission.replace(":id", submission.id)}>
                    {id}
                </Link>
            </td>
            <td className="padding-td-default">{date}</td>
            <td className="padding-td-default">
                <Link
                    className="link"
                    to={routeMap.profile.replace(":user", submission.user)}>
                    {user}
                </Link>
            </td>
            <td className="padding-td-default">
                <Link className="link" to={problem.href}>
                    {problem.text}
                </Link>
            </td>
            <td className="padding-td-default">{language}</td>
            <td
                className="padding-td-default"
                colSpan={maxScore === 0.0 ? 2 : 1}>
                <div className="flex items-center">
                    {verdictType === 0 && (
                        <SVGSpinner cls="w-4 h-4 shrink-0 mr-2" />
                    )}
                    {verdictType === 1 && (
                        <FontAwesomeIcon
                            icon="fa-xmark"
                            className="w-4 h-4 highlight-red mr-2"
                        />
                    )}
                    {verdictType === 2 && (
                        <FontAwesomeIcon
                            icon="fa-check"
                            className="w-4 h-4 highlight-yellow mr-2"
                        />
                    )}
                    {verdictType === 3 && (
                        <FontAwesomeIcon
                            icon="fa-check"
                            className="w-4 h-4 highlight-green mr-2"
                        />
                    )}
                    <span className="whitespace-nowrap">{verdictName}</span>
                </div>
            </td>
            {maxScore !== 0.0 && (
                <td className="padding-td-default w-0 text-center">
                    <span className="whitespace-nowrap">
                        {score} / {maxScore}
                    </span>
                </td>
            )}
            <td className="padding-td-default">{time} ms</td>
            <td className="padding-td-default">{memory} KiB</td>
        </tr>
    );
}

function SubmissionsTable({ submissions }) {
    const { t } = useTranslation();
    const submissionsContent = submissions.map((item, index) => (
        <Submission submission={item} key={index} />
    ));
    return (
        <RoundedTable>
            <thead className="bg-grey-800">
                <tr className="divide-x divide-dividecol">
                    <th className="padding-td-default">
                        {t("submissions_table.id")}
                    </th>
                    <th className="padding-td-default">
                        {t("submissions_table.date")}
                    </th>
                    <th className="padding-td-default">
                        {t("submissions_table.user")}
                    </th>
                    <th className="padding-td-default">
                        {t("submissions_table.problem")}
                    </th>
                    <th className="padding-td-default">
                        {t("submissions_table.language")}
                    </th>
                    <th className="padding-td-default" colSpan={2}>
                        {t("submissions_table.verdict")}
                    </th>
                    <th className="padding-td-default">
                        {t("submissions_table.time")}
                    </th>
                    <th className="padding-td-default">
                        {t("submissions_table.memory")}
                    </th>
                </tr>
            </thead>
            <tbody className="divide-y divide-dividecol">
                {submissionsContent}
            </tbody>
        </RoundedTable>
    );
}

export default SubmissionsTable;
