import RoundedTable from './RoundedTable';
import {SVGCorrectSimple, SVGWrong, SVGWrongSimple} from "../svg/SVGs";

function Submission({ id, date, user, problem, lang, verdict, time, mem }) {
    return (
        <tr className={"divide-x divide-default"}>
            <td className="padding-td-default">
                <span className="link">{id}</span>
            </td>
            <td className="padding-td-default">
                {date}
            </td>
            <td className="padding-td-default">
                <span className="link">{user}</span>
            </td>
            <td className="padding-td-default">
                <span className="link">{problem}</span>
            </td>
            <td className="padding-td-default">
                {lang}
            </td>
            <td className="padding-td-default">
                <div className="flex items-center">
                    {problem.includes("a") && <SVGWrongSimple cls="w-5 h-5 text-red-500 mr-2 shrink-0" />}
                    {!problem.includes("a") && <SVGCorrectSimple cls="w-5 h-5 text-indigo-500 mr-2 shrink-0"/>}
                    <span className="whitespace-nowrap">{verdict}</span>
                </div>
            </td>
            <td className="padding-td-default">
                {time}
            </td>
            <td className="padding-td-default">
                {mem}
            </td>
        </tr>
    );
}

function SubmissionsTable({ submissions }) {
    const submissionElems = submissions.map((item, index) =>
        <Submission id={item[0]} date={item[1]} user={item[2]} problem={item[3]} lang={item[4]} verdict={item[5]} time={item[6]} mem={item[7]} key={index} />
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
                {submissionElems}
            </tbody>
        </RoundedTable>
    );
}

export default SubmissionsTable;