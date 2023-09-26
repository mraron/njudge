import { useTranslation } from "react-i18next";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import Ranklist from "../../components/concrete/other/Ranklist";
import SVGTitleComponent from "../../components/svg/SVGTitleComponent";
import Pagination from "../../components/util/Pagination";

function ProblemRanklist({ data }) {
    const { t } = useTranslation();
    const titleComponent = (
        <SVGTitleComponent
            svg={
                <FontAwesomeIcon
                    icon="fa-ranking-star"
                    className="w-5 h-5 mr-2"
                />
            }
            title={t("problem_ranklist.ranklist")}
        />
    );

    return (
        <div>
            <div className="mb-2">
                <Ranklist
                    ranklist={data.ranklist}
                    titleComponent={titleComponent}
                />
            </div>
            <Pagination paginationData={data.paginationData} />
        </div>
    );
}

export default ProblemRanklist;
