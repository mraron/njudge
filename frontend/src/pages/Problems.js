import ProblemFilterFrame from "../components/concrete/filter/ProblemFilter"
import ProblemsTable from "../components/concrete/table/ProblemsTable"
import Pagination from "../components/util/Pagination"
import WidePage from "./wrappers/WidePage"

function Problems({ data }) {
    return (
        <WidePage>
            <div className="w-full space-y-2">
                <ProblemFilterFrame />
                <ProblemsTable problems={data.problems} />
                <Pagination paginationData={data.paginationData} />
            </div>
        </WidePage>
    )
}

export default Problems
