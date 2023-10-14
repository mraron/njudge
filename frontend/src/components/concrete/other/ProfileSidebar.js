import { useContext } from "react"
import { useTranslation } from "react-i18next"
import { Link } from "react-router-dom"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { faLineChart } from "@fortawesome/free-solid-svg-icons"
import MapDataFrame from "../../container/MapDataFrame"
import RoundedTable from "../../container/RoundedTable"
import UserContext from "../../../contexts/user/UserContext"
import RoundedFrame, { SVGTitleComponent } from "../../container/RoundedFrame"
import { routeMap } from "../../../config/RouteConfig"

export function ProfilePictureFrame({ userData }) {
    const profileRoute = routeMap.profile.replace(":user", encodeURIComponent(userData.username))
    return (
        <RoundedFrame cls="bg-frame-bg">
            <div className="flex flex-col items-center p-5 border-b border-border-def">
                <Link
                    to={profileRoute}
                    className="flex justify-center items-center w-full aspect-square bg-grey-875 border border-border-def hover:border-grey-450">
                    <img alt="avatar" className="object-contain" src={userData.pictureSrc} />
                </Link>
            </div>
            <div className="px-6 py-3 flex justify-center items-center w-full">
                <Link className="link truncate" to={profileRoute}>
                    {userData.username}
                </Link>
                <span className="text-xl emph-strong text-indigo-600 mx-2">&#8226;</span>
                <span className="truncate">{userData.rating}</span>
            </div>
        </RoundedFrame>
    )
}

export function ProfileDataFrame({ userData }) {
    const { t } = useTranslation()
    const titleComponent = (
        <SVGTitleComponent
            icon={<FontAwesomeIcon icon={faLineChart} className="w-4 h-4 mr-3" />}
            title={t("profile_sidebar.stats")}
        />
    )
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
    )
}

function SubmissionsFrame({ titleComponent, submissions }) {
    const rows = submissions.map((item, index) => (
        <tr key={index}>
            <td className="py-3">
                <Link className="link" to={routeMap.submission.replace(":id", item.id)}>
                    {item.id}
                </Link>
            </td>
            <td className="py-3">
                <Link className="link" to={item.problem.href}>
                    {item.problem.text}
                </Link>
            </td>
        </tr>
    ))
    return (
        <RoundedTable titleComponent={titleComponent}>
            <tbody>{rows}</tbody>
        </RoundedTable>
    )
}

function ProfileSideBar() {
    const { t } = useTranslation()
    const { userData, isLoggedIn } = useContext(UserContext)
    const titleComponent = (
        <SVGTitleComponent
            icon={<FontAwesomeIcon icon="fa-regular fa-clock" className="w-4 h-4 mr-3" />}
            title={t("profile_sidebar.last_submissions")}
        />
    )
    return (
        isLoggedIn && (
            <div className="hidden lg:flex flex-col w-72 shrink-0 space-y-3">
                <ProfilePictureFrame userData={userData} />
                <ProfileDataFrame userData={userData} />
                <SubmissionsFrame titleComponent={titleComponent} submissions={userData.lastSubmissions} />
            </div>
        )
    )
}

export default ProfileSideBar
