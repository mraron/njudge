import Ranklist from "../../components/concrete/other/Ranklist";
import SVGTitleComponent from '../../svg/SVGTitleComponent';
import {SVGResults} from "../../svg/SVGs";
import Pagination from "../../components/util/Pagination";
import {useOutletContext} from "react-router-dom";

function ProblemRanklist() {
    const data = useOutletContext()
    const titleComponent = <SVGTitleComponent svg={<SVGResults/>} title="Eredmények"/>

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