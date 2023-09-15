import {ProfileDataFrame, ProfilePictureFrame} from "../../components/ProfileSidebar";
import TagListFrame from "../../components/TagListFrame";
import SVGTitleComponent from "../../svg/SVGTitleComponent";
import {SVGCorrectSimple, SVGWrongSimple} from "../../svg/SVGs";
import {useOutletContext} from "react-router-dom";
import React from "react";
import checkData from "../../util/CheckData";

function ProfileMain() {
    const data = useOutletContext()
    if (!checkData(data)) {
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
                    <ProfilePictureFrame src="/assets/profile.webp" profileData={data.profileData}/>
                </div>
                <ProfileDataFrame profileData={data.profileData}/>
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