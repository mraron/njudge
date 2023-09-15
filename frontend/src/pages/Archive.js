import ProfileSideBar from '../components/ProfileSidebar'
import DropdownListFrame from '../components/DropdownListFrame'
import React from "react";
import checkData from "../util/CheckData";

function Archive({data}) {
    if (!checkData(data)) {
        return
    }
    const categoriesContent = data.categories.map((item, index) =>
        <div className="mb-3" key={index}>
            <DropdownListFrame title={item.title} tree={{"children": item.children}}/>
        </div>
    )
    return (
        <div className="relative w-full flex justify-center">
            <div className="flex justify-center w-full max-w-7xl">
                <div className="ml-0 lg:ml-4">
                    <ProfileSideBar profileData={data.profileData}/>
                </div>
                <div className="w-full px-4 lg:pl-3">
                    {categoriesContent}
                </div>
            </div>
        </div>
    );
}

export default Archive;