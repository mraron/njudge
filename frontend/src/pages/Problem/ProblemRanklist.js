import Rankings from "../../components/Rankings";
import SVGTitleComponent from '../../svg/SVGTitleComponent';
import {SVGResults} from "../../svg/SVGs";
import Pagination from "../../components/Pagination";

function ProblemRanklist() {
    const titleComponent = <SVGTitleComponent svg={<SVGResults/>} title="Eredmények"/>
    const data = [
        ["dbence", "50 / 50", "5669"],
        ["dbence", "50 / 50", "5669"],
        ["vpeti", "48 / 50", "5669"],
        ["vpeti", "48 / 50", "5669"],
        ["gonterarmin", "2 / 50", "5669"],
        ["gonterarmin", "2 / 50", "5669"],
    ]
    return (
        <div>
            <div className="mb-2">
                <Rankings data={data} titleComponent={titleComponent}/>
            </div>
            <Pagination paginationData={{currentPage: 1, lastPage: 200}}/>
        </div>
    )
}

export default ProblemRanklist;