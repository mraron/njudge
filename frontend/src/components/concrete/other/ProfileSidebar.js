import { useContext } from "react";
import { useTranslation } from "react-i18next";
import { Link } from "react-router-dom";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faClock, faLineChart } from "@fortawesome/free-solid-svg-icons";
import MapDataFrame from "../../container/MapDataFrame";
import SVGTitleComponent from "../../svg/SVGTitleComponent";
import RoundedTable from "../../container/RoundedTable";
import UserContext from "../../../contexts/user/UserContext";
import RoundedFrame from "../../container/RoundedFrame";
import { routeMap } from "../../../config/RouteConfig";

export function ProfilePictureFrame({ userData }) {
    const profileRoute = routeMap.profile.replace(
        ":user",
        encodeURIComponent(userData.username),
    );
    return (
        <RoundedFrame>
            <div className="flex flex-col items-center p-8 pb-4">
                <Link to={profileRoute} className="invert dark:invert-0">
                    <img
                        alt="avatar"
                        className="object-contain border border-default hover:border-grey-450 transition duration-200"
                        src={userData.pictureSrc}
                    />
                </Link>
                <div className="mt-2 flex justify-center items-center w-full">
                    <Link
                        className="link truncate font-medium no-underline"
                        to={profileRoute}>
                        {userData.username}
                    </Link>
                    <span className="text-2xl font-semibold text-indigo-500 mx-2 invert dark:invert-0">
                        &#8226;
                    </span>
                    <span className="truncate">{userData.rating}</span>
                </div>
            </div>
        </RoundedFrame>
    );
}

export function ProfileDataFrame({ userData }) {
    const { t } = useTranslation();
    const titleComponent = (
        <SVGTitleComponent
            svg={
                <FontAwesomeIcon icon={faLineChart} className="w-5 h-5 mr-2" />
            }
            title={t("profile_sidebar.stats")}
        />
    );
    return (
        <MapDataFrame
            data={[
                [t("profile_sidebar.rating"), `${userData.rating}`],
                [t("profile_sidebar.score"), `${userData.score}`],
                [t("profile_sidebar.num_solved"), `${userData.numSolved}`],
            ]}
            titleComponent={titleComponent}
            labelColWidth="14rem"
        />
    );
}

function SubmissionsFrame({ titleComponent, submissions }) {
    const rows = submissions.map((item, index) => (
        <tr className="divide-x divide-default" key={index}>
            <td className="padding-td-default">
                <Link
                    className="link"
                    to={routeMap.submission.replace(":id", item.id)}>
                    {item.id}
                </Link>
            </td>
            <td className="padding-td-default">
                <Link
                    className="link"
                    to={routeMap.problem.replace(":problem", item.problem)}>
                    {item.problem}
                </Link>
            </td>
        </tr>
    ));
    return (
        <RoundedTable titleComponent={titleComponent}>
            <tbody className="divide-y divide-default">{rows}</tbody>
        </RoundedTable>
    );
}

function ProfileSideBar() {
    const { t } = useTranslation();
    const { userData, isLoggedIn } = useContext(UserContext);
    const titleComponent = (
        <SVGTitleComponent
            svg={<FontAwesomeIcon icon={faClock} className="w-5 h-5 mr-2" />}
            title={t("profile_sidebar.last_submissions")}
        />
    );
    return (
        isLoggedIn && (
            <div className="w-full hidden lg:flex justify-center">
                <div className="flex flex-col bg-grey-900 w-80">
                    <div className="mb-3">
                        <ProfilePictureFrame userData={userData} />
                    </div>
                    <div className="mb-3">
                        <ProfileDataFrame userData={userData} />
                    </div>
                    <SubmissionsFrame
                        titleComponent={titleComponent}
                        submissions={userData.lastSubmissions}
                    />
                </div>
            </div>
        )
    );
}

export default ProfileSideBar;
