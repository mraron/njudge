import Ranklist from "../../components/concrete/other/Ranklist";
import SVGTitleComponent from '../../svg/SVGTitleComponent';
import {SVGResults} from "../../svg/SVGs";
import Pagination from "../../components/util/Pagination";
import {useTranslation} from "react-i18next";

function ProblemRanklist({data}) {
    const {t} = useTranslation()
    const titleComponent = <SVGTitleComponent svg={<SVGResults/>} title={t("problem_ranklist.ranklist")}/>

    return (
        <div>
            <div className="mb-2">
                <Ranklist ranklist={data.ranklist} titleComponent={titleComponent}/>
            </div>
            <Pagination paginationData={data.paginationData}/>
        </div>
    )
}

export default ProblemRanklist;