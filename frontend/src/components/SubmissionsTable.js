import RoundedTable from './RoundedTable';
import { SVGCorrectSimple, SVGSpinner, SVGWrongSimple } from "../svg/SVGs";
import { Link } from "react-router-dom";
import {routeMap} from "../config/RouteConfig"

function Submission({ submission }) {
    const {id, date, user, problem, lang, verdict, verdictType, time, memory} = submission
    return (
        <tr className={"divide-x divide-default"}>
            <td className="padding-td-default">
                <Link className="link" to={routeMap.submission.replace(":id", submission.id)}>{id}</Link>
            </td>
            <td className="padding-td-default">
                {date}
            </td>
            <td className="padding-td-default">
                <Link className="link" to={routeMap.submission.replace(":user", submission.user)}>{user}</Link>
            </td>
            <td className="padding-td-default">
                <Link className="link" to={problem.link}>{problem.label}</Link>
            </td>
            <td className="padding-td-default">
                {lang}
            </td>
            <td className="padding-td-default">
                <div className="flex items-center">
                    {verdictType === 0 && <SVGSpinner cls="w-5 h-5 mr-2 shrink-0" />}
                    {verdictType === 1 && <SVGWrongSimple cls="w-5 h-5 text-red-500 mr-2 shrink-0" />}
                    {verdictType === 2 && <SVGCorrectSimple cls="w-5 h-5 text-green-500 mr-2 shrink-0"/>}
                    <span className="whitespace-nowrap">{verdict}</span>
                </div>
            </td>
            <td className="padding-td-default">
                {time} ms
            </td>
            <td className="padding-td-default">
                {memory} KiB
            </td>
        </tr>
    );
}

function SubmissionsTable({ submissions }) {
    const submissionsContent = submissions.map((item, index) =>
        <Submission submission={item} key={index} />
    );
    return (
        <RoundedTable>
            <thead className="bg-grey-800">
                <tr className="divide-x divide-default">
                    <th className="padding-td-default">ID</th>
                    <th className="padding-td-default">Dátum</th>
                    <th className="padding-td-default">Felhasználó</th>
                    <th className="padding-td-default">Feladat</th>
                    <th className="padding-td-default">Nyelv</th>
                    <th className="padding-td-default">Verdikt</th>
                    <th className="padding-td-default">Idő</th>
                    <th className="padding-td-default">Memória</th>
                </tr>
            </thead>
            <tbody className="divide-y divide-default">
                {submissionsContent}
            </tbody>
        </RoundedTable>
    );
}

export default SubmissionsTable;