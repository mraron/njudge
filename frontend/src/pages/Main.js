import React, {useEffect, useState} from 'react';
import ProfileSideBar from '../components/ProfileSidebar'
import PostFrame from '../components/PostFrame'
import { SVGSpinner } from '../svg/SVGs';
import '../index.css';
import PageLoadingAnimation from "../components/PageLoadingAnimation";

function Main() {
    const [data, setData] = useState(null)

    useEffect(() => {
        fetch("/api/v2/")
            .then(res => res.json())
            .then(data => setData(data))
    }, []);
    let pageContent = <PageLoadingAnimation/>;
    if (data) {
        pageContent =
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
    }
    return (
        <div className="relative w-full flex justify-center">
            {pageContent}
        </div>
    );
}

export default Main;