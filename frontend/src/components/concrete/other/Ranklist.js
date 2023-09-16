import RoundedTable from '../../container/RoundedTable'
import {SVGCorrectSimple, SVGWrongSimple} from "../../../svg/SVGs";
import {routeMap} from "../../../config/RouteConfig";
import {Link} from "react-router-dom";

function RanklistRow({result, maxScore}) {
    const {username, score, submissionID, accepted} = result
    return (
        <tr className="divide-x divide-grey-700">
            <td className={`padding-td-default align-top`}>
                <Link className="link" to={routeMap.profile.replace(":user", username)}>{username}</Link>
            </td>
            <td className="padding-td-default sm:w-60">
                <div className="flex items-center">
                    {accepted && <SVGCorrectSimple cls="w-5 h-5 text-green-500 mr-2"/>}
                    {!accepted && <SVGWrongSimple cls="w-5 h-5 text-red-500 mr-2"/>}
                    <Link className="link whitespace-nowrap" to={routeMap.submission.replace(":id", submissionID)}>{score} / {maxScore}</Link>
                </div>
            </td>
        </tr>
    )
}

function Ranklist({ranklist, title, titleComponent}) {
    const rows = ranklist.results.map((item, index) => <RanklistRow result={item} maxScore={ranklist.maxScore} key={index}/>)
    return (
        <RoundedTable title={title} titleComponent={titleComponent}>
            <tbody className="divide-y divide-default bg-grey-850">
            {rows}
            </tbody>
        </RoundedTable>
    );
}

export default Ranklist;