import React from 'react';
import MapDataFrame from './MapDataFrame';
import { SVGDots, SVGRecent, SVGStatistics } from '../svg/SVGs';
import RoundedFrame from './RoundedFrame';
import SVGTitleComponent from '../svg/SVGTitleComponent';
import RoundedTable from './RoundedTable';

function ProfilePicture({ src }) {
    return (
        <img alt="avatar" className="object-contain border-1 border-default" src={src} />
    );
}

export function ProfileFrame({ src, username, rating }) {
    return (
        <div className="flex flex-col items-center">
            <div className="flex flex-col items-center p-8 pb-4 border-1 border-default rounded-md bg-grey-825 w-full">
                <ProfilePicture src={src} />
                <div className="flex justify-center items-center w-full">
                    <span className="mt-2 text-md font-medium truncate">
                        <a href="#">{username}</a>
                    </span>
                    <span className="mt-2 text-2xl font-semibold text-indigo-500 mx-2">
                        &#8226;
                    </span>
                    <span className="mt-2 text-md truncate">
                        {rating}
                    </span>
                </div>
            </div>
        </div>
    );
}

export function ProfileData({ rating, score,  solved }) {
    const titleComponent = <SVGTitleComponent svg={<SVGStatistics cls="w-6 h-6 mr-2" />} title="Statisztikák" />
    return (
        <MapDataFrame data={[
            ["Értékelés", `${rating}`],
            ["Pontszám", `${score}`],
            ["Megoldott feladatok", `${solved}`]
        ]} titleComponent={titleComponent} />
    );
}

function SubmissionsFrame({ titleComponent, submissions }) {
    const rows = submissions.map((item, index) => 
        <tr className="divide-x divide-default" key={index}>
            <td className="padding-td-default">
                <a className="link" href={item[2]}>{item[0]}</a>
            </td>
            <td className="padding-td-default">
                {item[1]}
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
    const titleComponent = <SVGTitleComponent svg={<SVGRecent />} title="Utolsó beküldések" />
    return (
        <div className="w-full hidden lg:flex justify-center">
            <div className="flex flex-col bg-grey-900 w-80">
                <div className="mb-3">
                    <ProfileFrame
                        src="/assets/profile.webp"
                        username="dbence"
                        rating={2350} />
                </div>
                <div className="mb-3">
                    <ProfileData rating={2350} score={65.4} solved={314} />
                </div>
                <SubmissionsFrame titleComponent={titleComponent} submissions={[
                    ["31415", "2023-09-06, 14:23:42", "#"], ["92653", "2023-09-06, 14:23:42", "#"], ["58979", "2023-09-06, 14:23:42", "#"]]} />
            </div>
        </div>
    );
}

export default ProfileSideBar;