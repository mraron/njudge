import RoundedTable from './RoundedTable'

function MapDataFrame({data, title, titleComponent}) {
    const rows = data.map((pair, index) =>
        <tr className="divide-x divide-default" key={index}>
            <td className="padding-td-default bg-grey-800 font-medium align-top whitespace-nowrap w-0">{pair[0]}</td>
            <td className="padding-td-default bg-grey-825 break-words" style={{maxWidth: 0}}>
                {pair[1]}
            </td>
        </tr>
    );
    return (
        <RoundedTable title={title} titleComponent={titleComponent}>
            <tbody className="divide-y divide-default">
                {rows}
            </tbody>
        </RoundedTable>
    );
}

export default MapDataFrame;