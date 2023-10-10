import Pagination from "../components/util/Pagination"
import SubmissionsTable from "../components/concrete/table/SubmissionsTable"
import WidePage from "./wrappers/WidePage";

function Submissions({ data }) {
    return (
        <WidePage>
            <div className="w-full space-y-2">
                <SubmissionsTable submissions={data.submissions} />
                <Pagination paginationData={data.paginationData} />
            </div>
        </WidePage>
    )
}

export default Submissions
