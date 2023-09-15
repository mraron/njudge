import React from 'react';
import ProfileSideBar from '../components/ProfileSidebar'
import PostFrame from '../components/PostFrame'
import { SVGSpinner } from '../svg/SVGs';
import '../index.css';
import {matchPath} from "react-router-dom";
import {routeMap} from "../config/RouteConfig";

function Main({ data }) {
    if (!data || data.processed) {
        return <></>
    }
    data.processed = true
    return (
        <div className="relative w-full flex justify-center">
            <div className="flex justify-center w-full max-w-7xl">
                <SVGSpinner cls="hidden w-16 h-16 absolute top-1/2 left-1/2"/>
                <div className="ml-0 lg:ml-4">
                    <ProfileSideBar
                        src="https://st3.depositphotos.com/6672868/13701/v/450/depositphotos_137014128-stock-illustration-user-profile-icon.jpg"
                        username="dbence"
                        score="2550"/>
                </div>
                <div className="w-full px-4 lg:pl-3">
                    <PostFrame posts={data.posts} />
                </div>
            </div>
        </div>
    );
}

export default Main;