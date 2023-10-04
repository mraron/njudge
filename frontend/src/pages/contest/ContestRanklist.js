import RoundedTable from "../../components/container/RoundedTable";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { Link } from "react-router-dom";
import { routeMap } from "../../config/RouteConfig";
import Pagination from "../../components/util/Pagination";
import DropdownFrame from "../../components/container/DropdownFrame";

function RanklistRow(data) {
    const {place, name, score, verdicts} = data.result
    const verdictsContent = verdicts.map((item, index) => {
            return (
                <td className="w-0 px-0" key={index}>
                    <span className="flex justify-center items-center">
                        {item.verdictType === 1 && <FontAwesomeIcon icon="fa-xmark" className="w-3.5 h-3.5 highlight-red" />}
                        {item.verdictType === 2 && <FontAwesomeIcon icon="fa-check" className="w-3.5 h-3.5 highlight-yellow" />}
                        {item.verdictType === 3 && <FontAwesomeIcon icon="fa-check" className="w-3.5 h-3.5 highlight-green" />}
                    </span>
                </td>
            )
        }
    )
    return (
        <tr className={`divide-x divide-dividecol ${data.index % 2 === 0? "bg-grey-850": "bg-grey-825"} hover:bg-grey-800 cursor-pointer`}>
            <td className="text-center">{place}</td>
            <td>
                <Link className="link" to={routeMap.profile.replace(":user", "dbence")}>{name}</Link>
            </td>
            <td className="w-0 text-center">{score}</td>
            {verdictsContent}
        </tr>
    )
}

function ContestRanklist({data}) {
    const ranklistContent = data.ranklist.map((item, index) =>
        <RanklistRow result={item} index={index} key={index} />
    )
    const problemsContent = data.problems.map((item, index) =>
        <th className="padding-td-default" key={index}>
            <Link className="link" to={item.href}>
                {item.text}
            </Link>
        </th>
    )
    return (
        <div className="space-y-2">
            <DropdownFrame title="Szűrés">...</DropdownFrame>
            <RoundedTable>
                <thead className="bg-grey-800">
                    <tr className="divide-x divide-dividecol">
                        <th className="w-0">#</th>
                        <th>Név</th>
                        <th className="w-0">=</th>
                        {problemsContent}
                    </tr>
                </thead>
                <tbody className="divide-y divide-dividecol">
                    {ranklistContent}
                </tbody>
            </RoundedTable>
            <Pagination paginationData={data.paginationData} />
        </div>
    )
}

export default ContestRanklist;
