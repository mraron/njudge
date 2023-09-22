import {ProfileDataFrame, ProfilePictureFrame} from "../../components/concrete/other/ProfileSidebar";
import TagListFrame, {LinkTag} from "../../components/container/TagListFrame";
import SVGTitleComponent from "../../svg/SVGTitleComponent";
import {SVGCorrectSimple, SVGWrongSimple} from "../../svg/SVGs";
import React from "react";
import {routeMap} from "../../config/RouteConfig";
import {useTranslation} from "react-i18next";

const makeProblemLink = (problem) => {
    return {"text": problem, "href": routeMap.problem.replace(":problem", problem)}
}

function ProfileMain({data}) {
    const {t} = useTranslation()
    const titleComponentCorrect = <SVGTitleComponent svg={<SVGCorrectSimple cls="w-6 h-6 text-green-500 mr-2"/>}
                                                     title={t("profile_main.solved_problems")}/>
    const titleComponentWrong = <SVGTitleComponent svg={<SVGWrongSimple cls="w-6 h-6 text-red-500 mr-2"/>}
                                                   title={t("profile_main.unsolved_problems")}/>

    return (
        <div className="flex flex-col sm:flex-row w-full items-start">
            <div className="w-full sm:w-80 mb-3 shrink-0">
                <div className="mb-3">
                    <ProfilePictureFrame src="/assets/profile.webp" userData={data.userData}/>
                </div>
                <ProfileDataFrame userData={data.userData}/>
            </div>
            <div className="w-full mb-3 sm:ml-3">
                <div className="mb-3">
                    <TagListFrame titleComponent={titleComponentCorrect} tags={data.solved.map(makeProblemLink)}
                                  tag={LinkTag}/>
                </div>
                <div className="mb-3">
                    <TagListFrame titleComponent={titleComponentWrong} tags={data.unsolved.map(makeProblemLink)}
                                  tag={LinkTag}/>
                </div>
            </div>
        </div>
    );
}

export default ProfileMain;