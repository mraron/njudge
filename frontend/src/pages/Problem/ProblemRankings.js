import Rankings from "../../components/Result";
import SVGTitleComponent from '../../svg/SVGTitleComponent';
import { SVGResults } from "../../svg/SVGs";

function ProblemRankings() {
    const titleComponent = <SVGTitleComponent svg={<SVGResults />} title="Eredmények" />
    const data = [
        ["dbence", "50 / 50", "5669"],
        ["dbence", "50 / 50", "5669"],
        ["vpeti", "48 / 50", "5669"],
        ["vpeti", "48 / 50", "5669"],
        ["gonterarmin", "2 / 50", "5669"],
        ["gonterarmin", "2 / 50", "5669"],
    ]
    return (
        <Rankings data={data} titleComponent={titleComponent} emphasize={false} />
    )
}

export default ProblemRankings;