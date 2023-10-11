import SubmissionsTable from "../../components/concrete/table/SubmissionsTable"
import Pagination from "../../components/util/Pagination"
import SubmissionFilterFrame from "../../components/concrete/filter/SubmissionFilter"

function ContestSubmissions({ data }) {
    return (
        <div className="flex flex-col space-y-2">
            <SubmissionFilterFrame optionOwn={data.isPublic} />
            <SubmissionsTable submissions={data.submissions} />
            <Pagination paginationData={data.paginationData} />
        </div>
    )
}

export default ContestSubmissions
