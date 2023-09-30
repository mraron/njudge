import { useTranslation } from "react-i18next";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import Ranklist from "../../components/concrete/other/Ranklist";
import SVGTitleComponent from "../../components/svg/SVGTitleComponent";
import Pagination from "../../components/util/Pagination";

function ProblemRanklist({ data }) {
    const { t } = useTranslation();
    return (
        <div>
            <div className="mb-2">
                <Ranklist
                    ranklist={data.ranklist}
                    title={t("problem_ranklist.ranklist")}
                />
            </div>
            <Pagination paginationData={data.paginationData} />
        </div>
    );
}

export default ProblemRanklist;
