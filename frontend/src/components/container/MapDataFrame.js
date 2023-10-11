import RoundedTable from "./RoundedTable";

function MapDataFrame({ data, title, titleComponent, labelColWidth }) {
    const rows = data.map((pair, index) => (
        <tr key={index}>
            <td className="padding-td-mapdata bg-grey-850 whitespace-nowrap" style={{ width: labelColWidth || "0" }}>
                {pair[0]}
            </td>
            <td className="padding-td-mapdata bg-grey-850 break-words">{pair[1]}</td>
        </tr>
    ))
    return (
        <RoundedTable title={title} titleComponent={titleComponent}>
            <tbody>{rows}</tbody>
        </RoundedTable>
    )
}

export default MapDataFrame
