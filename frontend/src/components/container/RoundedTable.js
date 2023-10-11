import RoundedFrame from "./RoundedFrame";

function RoundedTable({ children, title, titleComponent, cls }) {
    return (
        <RoundedFrame title={title} titleComponent={titleComponent} cls={`${cls} overflow-hidden`}>
            <div
                className={`w-full overflow-x-auto ${
                    title || titleComponent ? "rounded-b-container" : "rounded-container"
                }`}>
                <table
                    className={`w-full divide-y divide-highlight bg-grey-850 border-collapse text-table overflow-x-auto ${
                        title || titleComponent ? "rounded-b-container" : "rounded-container"
                    }`}>
                    {children}
                </table>
            </div>
        </RoundedFrame>
    )
}

export default RoundedTable
