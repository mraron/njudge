import Pagination from "../../components/util/Pagination"
import SubmissionsTable from "../../components/concrete/table/SubmissionsTable"

function ProfileSubmissions({ data }) {
    return (
        <div className="flex flex-col w-full">
            <div className="mb-2">
                <SubmissionsTable submissions={data.submissions} />
            </div>
            <Pagination paginationData={data.paginationData} />
        </div>
    )
}

export default ProfileSubmissions
