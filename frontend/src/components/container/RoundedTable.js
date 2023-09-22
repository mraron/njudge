import RoundedFrame from "./RoundedFrame";

function RoundedTable({children, title, titleComponent}) {
    return (
        <RoundedFrame title={title} titleComponent={titleComponent}>
            <div
                className={`w-full overflow-x-auto ${title || titleComponent ? "rounded-bl-md rounded-br-md" : "rounded-md"}`}>
                <table className="w-full divide-y divide-indigo-600 bg-grey-850 border-collapse text-table overflow-x-auto">
                    {children}
                </table>
            </div>
        </RoundedFrame>
    );
}

export default RoundedTable;