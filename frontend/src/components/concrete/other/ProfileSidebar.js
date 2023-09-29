import { useContext } from "react";
import { useTranslation } from "react-i18next";
import { Link } from "react-router-dom";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faLineChart } from "@fortawesome/free-solid-svg-icons";
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
            <div className="flex flex-col items-center p-6 pb-4">
                <Link to={profileRoute}>
                    <img
                        alt="avatar"
                        className="object-contain border border-bordercol hover:border-grey-450"
                        src={userData.pictureSrc}
                    />
                </Link>
                <div className="mt-2 flex justify-center items-center w-full">
                    <Link
                        className="link truncate font-medium no-underline"
                        to={profileRoute}>
                        {userData.username}
                    </Link>
                    <span className="text-2xl font-semibold text-indigo-600 mx-2">
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
                <FontAwesomeIcon icon={faLineChart} className="w-4 h-4 mr-3" />
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
            labelColWidth="12rem"
        />
    );
}

function SubmissionsFrame({ titleComponent, submissions }) {
    const rows = submissions.map((item, index) => (
        <tr className="divide-x divide-dividecol" key={index}>
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
                    to={item.problem.href}>
                    {item.problem.text}
                </Link>
            </td>
        </tr>
    ));
    return (
        <RoundedTable titleComponent={titleComponent}>
            <tbody className="divide-y divide-dividecol">{rows}</tbody>
        </RoundedTable>
    );
}

function ProfileSideBar() {
    const { t } = useTranslation();
    const { userData, isLoggedIn } = useContext(UserContext);
    const titleComponent = (
        <SVGTitleComponent
            svg={
                <FontAwesomeIcon
                    icon="fa-regular fa-clock"
                    className="w-4 h-4 mr-3"
                />
            }
            title={t("profile_sidebar.last_submissions")}
        />
    );
    return (
        isLoggedIn && (
            <div className="w-full hidden lg:flex justify-center">
                <div className="flex flex-col bg-grey-900 w-72">
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
