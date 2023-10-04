import Pagination from "../../components/util/Pagination";
import SubmissionsTable from "../../components/concrete/table/SubmissionsTable";
import SubmissionFilterFrame from "../../components/concrete/other/SubmissionFilter";

function ProblemSubmissions({ data }) {
    return (
        <div className="relative space-y-2">
            <SubmissionFilterFrame />
            <SubmissionsTable submissions={data.submissions} />
            <Pagination paginationData={data.paginationData} />
        </div>
    );
}

export default ProblemSubmissions;
