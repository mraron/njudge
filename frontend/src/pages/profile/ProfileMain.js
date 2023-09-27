import { useTranslation } from "react-i18next";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import SVGTitleComponent from "../../components/svg/SVGTitleComponent";
import TagListFrame, { LinkTag } from "../../components/container/TagListFrame";
import {
    ProfileDataFrame,
    ProfilePictureFrame,
} from "../../components/concrete/other/ProfileSidebar";

function ProfileMain({ data }) {
    const { t } = useTranslation();
    const titleComponentCorrect = (
        <SVGTitleComponent
            svg={
                <FontAwesomeIcon
                    icon="fa-check"
                    className="w-5 h-5 highlight-green mr-3"
                />
            }
            title={t("profile_main.solved_problems")}
        />
    );
    const titleComponentWrong = (
        <SVGTitleComponent
            svg={
                <FontAwesomeIcon
                    icon="fa-xmark"
                    className="w-5 h-5 highlight-red mr-3"
                />
            }
            title={t("profile_main.unsolved_problems")}
        />
    );
    return (
        <div className="flex flex-col sm:flex-row w-full items-start">
            <div className="w-full sm:w-72 mb-3 shrink-0">
                <div className="mb-3">
                    <ProfilePictureFrame
                        src="/assets/profile.webp"
                        userData={data.userData}
                    />
                </div>
                <ProfileDataFrame userData={data.userData} />
            </div>
            <div className="w-full mb-3 sm:ml-3">
                <div className="mb-3">
                    <TagListFrame
                        titleComponent={titleComponentCorrect}
                        tags={data.solved}
                        tag={LinkTag}
                    />
                </div>
                <div className="mb-3">
                    <TagListFrame
                        titleComponent={titleComponentWrong}
                        tags={data.unsolved}
                        tag={LinkTag}
                    />
                </div>
            </div>
        </div>
    );
}

export default ProfileMain;
