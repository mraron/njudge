import ProfileSideBar from '../components/concrete/other/ProfileSidebar'
import DropdownListFrame from '../components/container/DropdownListFrame'
import React from "react";
import {Link} from "react-router-dom";
import {SVGCorrectSimple, SVGWrongSimple} from "../svg/SVGs";

function ProblemLeaf({data}) {
    return (
        <Link className="w-fit flex items-center cursor-pointer hover:text-indigo-300 transition-all duration-100"
              to={data.link}>
            <div className="w-5 mr-2">
                {data.solvedStatus === 1 && <SVGWrongSimple cls="w-5 h-5 text-red-500 shrink-0" />}
                {data.solvedStatus === 2 && <SVGCorrectSimple cls="w-5 h-5 text-indigo-500 shrink-0" />}
                {data.solvedStatus === 3 && <SVGCorrectSimple cls="w-5 h-5 text-green-500 shrink-0" />}
            </div>
            <span>{data.title}</span>
        </Link>
    )
}

function Archive({data}) {
    const categoriesContent = data.categories.map((item, index) =>
        <div className="mb-3" key={index}>
            <DropdownListFrame title={item.title} tree={{"children": item.children}} leaf={ProblemLeaf}/>
        </div>
    )
    return (
        <div className="relative w-full flex justify-center">
            <div className="flex justify-center w-full max-w-7xl">
                <div className="ml-0 lg:ml-4">
                    <ProfileSideBar/>
                </div>
                <div className="w-full px-4 lg:pl-3">
                    {categoriesContent}
                </div>
            </div>
        </div>
    );
}

export default Archive;