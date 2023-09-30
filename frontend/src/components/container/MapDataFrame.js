import RoundedTable from "./RoundedTable";

function MapDataFrame({ data, title, titleComponent, labelColWidth }) {
    const rows = data.map((pair, index) => (
        <tr className="divide-x divide-dividecol" key={index}>
            <td
                className="padding-td-default bg-grey-800 whitespace-nowrap"
                style={{ width: labelColWidth || "0" }}>
                {pair[0]}
            </td>
            <td className="padding-td-default bg-grey-850 break-words">
                {pair[1]}
            </td>
        </tr>
    ));
    return (
        <RoundedTable title={title} titleComponent={titleComponent}>
            <tbody className="divide-y divide-dividecol">{rows}</tbody>
        </RoundedTable>
    );
}

export default MapDataFrame;
