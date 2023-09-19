import React, {useContext} from 'react';
import MapDataFrame from '../../container/MapDataFrame';
import {SVGRecent, SVGStatistics} from '../../../svg/SVGs';
import SVGTitleComponent from '../../../svg/SVGTitleComponent';
import RoundedTable from '../../container/RoundedTable';
import {routeMap} from "../../../config/RouteConfig";
import {Link} from "react-router-dom";
import UserContext from "../../../contexts/user/UserContext";
import {useTranslation} from "react-i18next";

export function ProfilePictureFrame({userData}) {
    const profileRoute = routeMap.profile.replace(":user", encodeURIComponent(userData.username))
    return (
        <div className="flex flex-col items-center">
            <div className="flex flex-col items-center p-8 pb-4 border-1 border-default rounded-md bg-grey-825 w-full">
                <Link to={profileRoute}>
                    <img alt="avatar"
                         className="object-contain border-1 border-default hover:border-grey-450 transition duration-200"
                         src={userData.pictureSrc}/>
                </Link>
                <div className="flex justify-center items-center w-full">
                    <Link className="mt-2 text-md font-medium truncate hover:text-indigo-200 transition duration-200"
                          to={profileRoute}>
                        {userData.username}
                    </Link>
                    <span className="mt-2 text-2xl font-semibold text-indigo-500 mx-2">
                        &#8226;
                    </span>
                    <span className="mt-2 text-md truncate">
                        {userData.rating}
                    </span>
                </div>
            </div>
        </div>
    );
}

export function ProfileDataFrame({userData}) {
    const {t} = useTranslation()
    const titleComponent = <SVGTitleComponent svg={<SVGStatistics cls="w-6 h-6 mr-2"/>} title={t("profile_sidebar.stats")}/>
    return (
        <MapDataFrame data={[
            [t("profile_sidebar.rating"), `${userData.rating}`],
            [t("profile_sidebar.score"), `${userData.score}`],
            [t("profile_sidebar.num_solved"), `${userData.numSolved}`]
        ]} titleComponent={titleComponent}/>
    );
}

function SubmissionsFrame({titleComponent, submissions}) {
    const rows = submissions.map((item, index) =>
        <tr className="divide-x divide-default" key={index}>
            <td className="padding-td-default">
                <Link className="link" to={routeMap.submission.replace(":id", item.id)}>{item.id}</Link>
            </td>
            <td className="padding-td-default">
                <Link className="link" to={routeMap.problem.replace(":problem", item.problem)}>{item.problem}</Link>
            </td>
        </tr>
    )
    return (
        <RoundedTable titleComponent={titleComponent}>
            <tbody className="divide-y divide-default">
            {rows}
            </tbody>
        </RoundedTable>
    )
}

function ProfileSideBar() {
    const {t} = useTranslation()
    const {userData, isLoggedIn} = useContext(UserContext)
    const titleComponent = <SVGTitleComponent svg={<SVGRecent/>} title={t("profile_sidebar.last_submissions")}/>
    return (
        isLoggedIn &&
        <div className="w-full hidden lg:flex justify-center">
            <div className="flex flex-col bg-grey-900 w-80">
                <div className="mb-3">
                    <ProfilePictureFrame userData={userData}/>
                </div>
                <div className="mb-3">
                    <ProfileDataFrame userData={userData}/>
                </div>
                <SubmissionsFrame titleComponent={titleComponent} submissions={userData.lastSubmissions}/>
            </div>
        </div>
    );
}

export default ProfileSideBar;