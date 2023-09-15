import RoundedFrame from "./RoundedFrame";

function RoundedTable({children, title, titleComponent}) {
    return (
        <RoundedFrame title={title} titleComponent={titleComponent}>
            <div
                className={`flex flex-col w-full overflow-x-auto ${title || titleComponent ? "rounded-bl-md rounded-br-md" : "rounded-md"} text-table`}>
                <table className="table-fixed divide-y divide-indigo-600 bg-grey-850 border-collapse">
                    {children}
                </table>
            </div>
        </RoundedFrame>
    );
}

export default RoundedTable;