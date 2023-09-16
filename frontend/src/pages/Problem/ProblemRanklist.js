import Ranklist from "../../components/Ranklist";
import SVGTitleComponent from '../../svg/SVGTitleComponent';
import {SVGResults} from "../../svg/SVGs";
import Pagination from "../../components/Pagination";
import {useOutletContext} from "react-router-dom";
import checkData from "../../util/CheckData";

function ProblemRanklist() {
    const data = useOutletContext()
    if (!checkData(data)) {
        return
    }
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