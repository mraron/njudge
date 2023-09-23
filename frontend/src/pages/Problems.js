import { ProblemFilterFrame } from "../components/concrete/other/ProblemFilter";
import ProblemsTable from "../components/concrete/table/ProblemsTable";
import ProfileSideBar from "../components/concrete/other/ProfileSidebar";
import Pagination from "../components/util/Pagination";

function Problems({ data }) {
    return (
        <div className="relative w-full flex justify-center">
            <div className="flex justify-center w-full max-w-7xl">
                <div className="ml-0 lg:ml-4">
                    <ProfileSideBar />
                </div>
                <div className="w-full flex flex-col overflow-x-auto">
                    <div className="w-full px-4 lg:pl-3">
                        <div className="mb-2">
                            <ProblemFilterFrame />
                        </div>
                        <div className="mb-2">
                            <ProblemsTable problems={data.problems} />
                        </div>
                        <Pagination paginationData={data.paginationData} />
                    </div>
                </div>
            </div>
        </div>
    );
}

export default Problems;
