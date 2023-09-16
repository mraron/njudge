import React from 'react';
import MapDataFrame from './MapDataFrame';
import {SVGRecent, SVGStatistics} from '../svg/SVGs';
import SVGTitleComponent from '../svg/SVGTitleComponent';
import RoundedTable from './RoundedTable';
import {routeMap} from "../config/RouteConfig";
import {Link} from "react-router-dom";

function ProfilePicture({src}) {
    return (
        <img alt="avatar" className="object-contain border-1 border-default" src={src}/>
    );
}

export function ProfilePictureFrame({profileData}) {
    return (
        <div className="flex flex-col items-center">
            <div className="flex flex-col items-center p-8 pb-4 border-1 border-default rounded-md bg-grey-825 w-full">
                <ProfilePicture src={profileData.pictureSrc}/>
                <div className="flex justify-center items-center w-full">
                    <span className="mt-2 text-md font-medium truncate">
                        <a href="#">{profileData.username}</a>
                    </span>
                    <span className="mt-2 text-2xl font-semibold text-indigo-500 mx-2">
                        &#8226;
                    </span>
                    <span className="mt-2 text-md truncate">
                        {profileData.rating}
                    </span>
                </div>
            </div>
        </div>
    );
}

export function ProfileDataFrame({profileData}) {
    const titleComponent = <SVGTitleComponent svg={<SVGStatistics cls="w-6 h-6 mr-2"/>} title="Statisztikák"/>
    return (
        <MapDataFrame data={[
            ["Értékelés", `${profileData.rating}`],
            ["Pontszám", `${profileData.score}`],
            ["Megoldott feladatok", `${profileData.numSolved}`]
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

function ProfileSideBar({profileData}) {
    const titleComponent = <SVGTitleComponent svg={<SVGRecent/>} title="Utolsó beküldések"/>
    return (
        <div className="w-full hidden lg:flex justify-center">
            <div className="flex flex-col bg-grey-900 w-80">
                <div className="mb-3">
                    <ProfilePictureFrame profileData={profileData}/>
                </div>
                <div className="mb-3">
                    <ProfileDataFrame profileData={profileData}/>
                </div>
                <SubmissionsFrame titleComponent={titleComponent} submissions={profileData.lastSubmissions}/>
            </div>
        </div>
    );
}

export default ProfileSideBar;