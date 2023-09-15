import RoundedTable from './RoundedTable'
import {SVGCorrectSimple, SVGWrongSimple} from "../svg/SVGs";

function RankingRow({ name, score }) {
    return (
        <tr className="divide-x divide-grey-700">
            <td className={`padding-td-default align-top`}>
                <span className="link">{name}</span>
            </td>
            <td className="padding-td-default sm:w-60">
                <div className="flex items-center">
                    {score[0] === '5' && <SVGCorrectSimple cls="w-5 h-5 text-green-500 mr-2" />}
                    {score[0] !== '5' && <SVGWrongSimple cls="w-5 h-5 text-red-500 mr-2" />}
                    <span className="link whitespace-nowrap">{score}</span>
                </div>
            </td>
        </tr>
    )
}

function Rankings({ data, title, titleComponent, emphasize }) {
    if (emphasize == null) {
        emphasize = true;
    }
    const rows = data.map((pair, index) => <RankingRow name={pair[0]} score={pair[1]} key={index} emphasize={emphasize} />)
    return (
        <RoundedTable title={title} titleComponent={titleComponent}>
            <tbody className="divide-y divide-default bg-grey-850">
                {rows}
            </tbody>
        </RoundedTable>
    );
}

export default Rankings;