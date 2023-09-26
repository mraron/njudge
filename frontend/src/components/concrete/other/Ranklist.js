import { Link } from "react-router-dom";
import { routeMap } from "../../../config/RouteConfig";
import RoundedFrame from "../../container/RoundedFrame";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import React from "react";

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
                    <FontAwesomeIcon
                        icon="fa-check"
                        className="w-4 h-4 highlight-green mr-2"
                    />
                )}
                {!accepted && (
                    <FontAwesomeIcon
                        icon="fa-xmark"
                        className="w-4 h-4 highlight-red mr-2"
                    />
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
