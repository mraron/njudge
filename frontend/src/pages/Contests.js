import ProfileSideBar from "../components/concrete/other/ProfileSidebar";
import ContestList from "../components/concrete/other/ContestList";
import React from "react";

function Contests() {
    return (
        <div className="w-full flex justify-center">
            <div className="flex justify-center w-full max-w-7xl lg:space-x-3 px-4">
                <ProfileSideBar />
                <div className="w-full min-w-0">
                    <ContestList
                        contestData={[
                            [
                                "Online oktató programozóverseny #3",
                                "2023-12-23, 14:00",
                                true,
                            ],
                            [
                                "Online oktató programozóverseny #2",
                                "2023-08-23, 14:00",
                                false,
                            ],
                            [
                                "Online oktató programozóverseny #1",
                                "2023-04-23, 14:00",
                                false,
                            ],
                        ]}
                    />
                </div>
            </div>
        </div>
    );
}

export default Contests;
