import {ProfilePictureFrame, ProfileDataFrame} from "../../components/concrete/other/ProfileSidebar";
import TagListFrame from "../../components/container/TagListFrame";
import SVGTitleComponent from "../../svg/SVGTitleComponent";
import {SVGCorrectSimple, SVGWrongSimple} from "../../svg/SVGs";
import {matchPath, useLocation, useOutletContext} from "react-router-dom";
import React from "react";

function ProfileMain() {
    const data = useOutletContext()
    const location = useLocation()

    if (!data || !matchPath(data.route, location.pathname)) {
        return
    }
    const titleComponentCorrect = <SVGTitleComponent svg={<SVGCorrectSimple cls="w-6 h-6 text-green-500 mr-2"/>}
                                                     title="Megoldott feladatok"/>
    const titleComponentWrong = <SVGTitleComponent svg={<SVGWrongSimple cls="w-6 h-6 text-red-500 mr-2"/>}
                                                   title="Megpróbált feladatok"/>

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
                    <TagListFrame titleComponent={titleComponentCorrect} tagNames={data.solved}/>
                </div>
                <div className="mb-3">
                    <TagListFrame titleComponent={titleComponentWrong} tagNames={data.unsolved}/>
                </div>
            </div>
        </div>
    );
}

export default ProfileMain;