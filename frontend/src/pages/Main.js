import React from 'react';
import ProfileSideBar from '../components/concrete/other/ProfileSidebar'
import PostFrame from '../components/container/PostFrame'
import {SVGSpinner} from '../svg/SVGs';
import '../index.css';
import checkData from "../util/CheckData";
import {matchPath, useLocation} from "react-router-dom";

function Main({data}) {
    const location = useLocation()
    if (!data || !matchPath(data.route, location.pathname)) {
        return
    }
    return (
        <div className="relative w-full flex justify-center">
            <div className="flex justify-center w-full max-w-7xl">
                <SVGSpinner cls="hidden w-16 h-16 absolute top-1/2 left-1/2"/>
                <div className="ml-0 lg:ml-4">
                    <ProfileSideBar/>
                </div>
                <div className="w-full px-4 lg:pl-3">
                    <PostFrame posts={data.posts}/>
                </div>
            </div>
        </div>
    );
}

export default Main;