import RoundedTable from './RoundedTable'

function MapData({ data, title, titleComponent }) {
    const rows = data.map((pair, index) =>     
        <tr className="divide-x divide-grey-700" key={index}>
            <td className="padding-td-default bg-grey-800 font-medium align-top">{pair[0]}</td>
            <td className="padding-td-default bg-grey-825">{pair[1]}</td>
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

export default MapData;