import { Link } from "react-router-dom";
import { SVGCorrectSimple, SVGWrongSimple } from "../../svg/SVGs";
import { routeMap } from "../../../config/RouteConfig";
import RoundedFrame from "../../container/RoundedFrame";

function RanklistRow({ result, maxScore }) {
    const { username, score, submissionID, accepted } = result;
    return (
        <div className="w-full flex items-center justify-between py-3 px-6">
            <Link
                className="link"
                to={routeMap.profile.replace(":user", username)}>
                {username}
            </Link>
            <div className="flex">
                {accepted && (
                    <SVGCorrectSimple cls="w-5 h-5 text-green-600 mr-2" />
                )}
                {!accepted && (
                    <SVGWrongSimple cls="w-5 h-5 text-red-600 mr-2" />
                )}
                <Link
                    className="link whitespace-nowrap"
                    to={routeMap.submission.replace(":id", submissionID)}>
                    {score} / {maxScore}
                </Link>
            </div>
        </div>
    );
}

function Ranklist({ ranklist, title, titleComponent }) {
    const rows = ranklist.results.map((item, index) => (
        <RanklistRow result={item} maxScore={ranklist.maxScore} key={index} />
    ));
    return (
        <RoundedFrame
            title={title}
            titleComponent={titleComponent}
            cls="overflow-hidden">
            <div className="divide-y divide-dividecol bg-grey-850">{rows}</div>
        </RoundedFrame>
    );
}

export default Ranklist;
