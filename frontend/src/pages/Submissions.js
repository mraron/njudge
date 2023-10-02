import Pagination from "../components/util/Pagination";
import SubmissionsTable from "../components/concrete/table/SubmissionsTable";

function Submissions({ data }) {
    return (
        <div className="w-full flex justify-center">
            <div className="flex justify-center w-full max-w-7xl">
                <div className="w-full px-4 overflow-x-auto space-y-2">
                    <SubmissionsTable submissions={data.submissions} />
                    <Pagination paginationData={data.paginationData} />
                </div>
            </div>
        </div>
    );
}

export default Submissions;
