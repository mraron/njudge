import RoundedTable from './RoundedTable'

function MapDataFrame({data, maxDataWidth, title, titleComponent}) {
    maxDataWidth ||= "auto"
    const rows = data.map((pair, index) =>
        <tr className="divide-x divide-grey-700" key={index}>
            <td className="padding-td-default bg-grey-800 font-medium align-top whitespace-nowrap">{pair[0]}</td>
            <td className="padding-td-default bg-grey-825 break-words" style={{maxWidth: maxDataWidth}}>
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