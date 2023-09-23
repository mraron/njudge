import { useTranslation } from "react-i18next";
import { SVGResults } from "../../components/svg/SVGs";
import Ranklist from "../../components/concrete/other/Ranklist";
import SVGTitleComponent from "../../components/svg/SVGTitleComponent";
import Pagination from "../../components/util/Pagination";

function ProblemRanklist({ data }) {
    const { t } = useTranslation();
    const titleComponent = (
        <SVGTitleComponent
            svg={<SVGResults />}
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
