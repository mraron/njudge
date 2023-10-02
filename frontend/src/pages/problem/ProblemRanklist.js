import { useTranslation } from "react-i18next";
import Ranklist from "../../components/concrete/other/Ranklist";
import Pagination from "../../components/util/Pagination";

function ProblemRanklist({ data }) {
    const { t } = useTranslation();
    return (
        <div className="space-y-2">
            <Ranklist
                ranklist={data.ranklist}
                title={t("problem_ranklist.ranklist")}
            />
            <Pagination paginationData={data.paginationData} />
        </div>
    );
}

export default ProblemRanklist;
