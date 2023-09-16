import ProfileSideBar from '../components/concrete/other/ProfileSidebar'
import DropdownListFrame from '../components/container/DropdownListFrame'
import checkData from "../util/CheckData";
import React from "react";

function Archive({data}) {
    if (!checkData(data)) {
        return
    }
    window.flash("Sikeres belépés", "success")
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