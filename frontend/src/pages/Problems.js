import ProblemFilterFrame from "../components/concrete/filter/ProblemFilter"
import ProblemsTable from "../components/concrete/table/ProblemsTable"
import Pagination from "../components/util/Pagination"

function Problems({ data }) {
    return (
        <div className="w-full flex justify-center">
            <div className="flex justify-center w-full max-w-7xl">
                <div className="w-full flex flex-col overflow-x-auto">
                    <div className="w-full px-4 space-y-2">
                        <ProblemFilterFrame />
                        <ProblemsTable problems={data.problems} />
                        <Pagination paginationData={data.paginationData} />
                    </div>
                </div>
            </div>
        </div>
    )
}

export default Problems
