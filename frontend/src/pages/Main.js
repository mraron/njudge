import React from 'react';
import ProfileSideBar from '../components/ProfileSidebar'
import PostFrame from '../components/PostFrame'
import {SVGSpinner} from '../svg/SVGs';
import '../index.css';
import checkData from "../util/CheckData";

function Main({data}) {
    if (!checkData(data)) {
        return
    }
    return (
        <div className="relative w-full flex justify-center">
            <div className="flex justify-center w-full max-w-7xl">
                <SVGSpinner cls="hidden w-16 h-16 absolute top-1/2 left-1/2"/>
                <div className="ml-0 lg:ml-4">
                    <ProfileSideBar profileData={data.profileData}/>
                </div>
                <div className="w-full px-4 lg:pl-3">
                    <PostFrame posts={data.posts}/>
                </div>
            </div>
        </div>
    );
}

export default Main;