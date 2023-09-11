import RoundedTable from './RoundedTable'

function RankingRow({ name, score, emphasize }) {
    return (
        <tr className="divide-x divide-grey-700">
            <td className={`padding-td-default bg-grey-800 ${emphasize? "font-medium": ""} align-top`}>
                <span className="link">{name}</span>
            </td>
            <td className="padding-td-default bg-grey-825 text-center whitespace-nowrap">
                <span className="link">{score}</span>
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
            <tbody className="divide-y divide-default">
                {rows}
            </tbody>
        </RoundedTable>
    );
}

export default Rankings;