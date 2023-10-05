import ProfileSideBar from "../components/concrete/other/ProfileSidebar";
import ContestList from "../components/concrete/other/ContestList";
import React from "react";

function Contests({ data }) {
    return (
        <div className="w-full flex justify-center">
            <div className="flex justify-center w-full max-w-7xl lg:space-x-3 px-4">
                <ProfileSideBar />
                <div className="w-full min-w-0">
                    <ContestList contests={data.contests} />
                </div>
            </div>
        </div>
    );
}

export default Contests;
