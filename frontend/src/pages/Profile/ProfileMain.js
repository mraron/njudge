import { ProfileData, ProfileFrame } from "../../components/ProfileSidebar";
import TagListFrame from "../../components/TagListFrame";
import SVGTitleComponent from "../../svg/SVGTitleComponent";
import {SVGCorrect, SVGCorrectSimple, SVGWrong, SVGWrongSimple} from "../../svg/SVGs";

function ProfileMain() {
    const titleComponentCorrect = <SVGTitleComponent svg={<SVGCorrectSimple cls="w-6 h-6 text-green-500 mr-2" />} title="Megoldott feladatok" />
    const titleComponentWrong = <SVGTitleComponent svg={<SVGWrongSimple cls="w-6 h-6 text-red-500 mr-2" />} title="Megpróbált feladatok" />
    return (
        <div className="flex flex-col sm:flex-row w-full items-start">
            <div className="w-full sm:w-80 mb-3 shrink-0">
                <div className="mb-3">
                    <ProfileFrame
                        src="/assets/profile.webp"
                        username="dbence"
                        rating={2350}/>
                </div>
                <ProfileData rating={2350} score={65.4} solved={187} />
            </div>
            <div className="w-full mb-3 sm:ml-3">
                <div className="mb-3">
                    <TagListFrame titleComponent={titleComponentCorrect} tagNames={[
                        "KK23_tomjerry",
                        "KK23_swaps",
                        "KK23_speeding",
                        "KK23_snacks",
                        "KK23_rusco",
                        "KK23_tomjerry",
                        "KK23_swaps",
                        "KK23_speeding",
                        "KK23_snacks",
                        "KK23_rusco",
                    ]} />
                </div>
                <div className="mb-3">
                    <TagListFrame titleComponent={titleComponentWrong} tagNames={[
                        "KK23_tomjerry",
                        "KK23_swaps",
                        "KK23_speeding",
                        "KK23_snacks",
                        "KK23_rusco",
                        "KK23_tomjerry",
                        "KK23_swaps",
                        "KK23_speeding",
                        "KK23_snacks",
                        "KK23_rusco",
                    ]} />
                </div>
            </div>
        </div>
    );
}

export default ProfileMain;