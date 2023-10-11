import { useTranslation } from "react-i18next"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { SVGTitleComponent } from "../../components/container/RoundedFrame"
import TagListFrame, { LinkTag } from "../../components/container/TagListFrame"
import { ProfileDataFrame, ProfilePictureFrame } from "../../components/concrete/other/ProfileSidebar"

function ProfileMain({ data }) {
    const { t } = useTranslation()
    const titleComponentCorrect = (
        <SVGTitleComponent
            icon={<FontAwesomeIcon icon="fa-check" className="w-5 h-5 highlight-green mr-3" />}
            title={t("profile_main.solved_problems")}
        />
    )
    const titleComponentWrong = (
        <SVGTitleComponent
            icon={<FontAwesomeIcon icon="fa-xmark" className="w-5 h-5 highlight-red mr-3" />}
            title={t("profile_main.unsolved_problems")}
        />
    )
    return (
        <div className="flex flex-col sm:flex-row w-full items-start space-y-3 sm:space-y-0 sm:space-x-3">
            <div className="w-full sm:w-72 shrink-0 space-y-3">
                <ProfilePictureFrame src="/assets/profile.webp" userData={data.userData} />
                <ProfileDataFrame userData={data.userData} />
            </div>
            <div className="w-full min-w-0 space-y-3">
                <TagListFrame titleComponent={titleComponentCorrect} tags={data.solved} tag={LinkTag} />
                <TagListFrame titleComponent={titleComponentWrong} tags={data.unsolved} tag={LinkTag} />
            </div>
        </div>
    )
}

export default ProfileMain
