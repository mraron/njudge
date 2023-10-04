import { Link } from "react-router-dom";
import ProfileSideBar from "../components/concrete/other/ProfileSidebar";
import DropdownListFrame from "../components/container/DropdownListFrame";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import React from "react";

function ProblemLeaf({ data }) {
    return (
        <span className="ml-2 max-w-fit flex items-center">
            <div className="w-4 mr-3 flex justify-center items-center">
                {data.solvedStatus === 1 && (
                    <FontAwesomeIcon
                        icon="fa-xmark"
                        className="w-4 h-4 highlight-red"
                    />
                )}
                {data.solvedStatus === 2 && (
                    <FontAwesomeIcon
                        icon="fa-check"
                        className="w-4 h-4 highlight-yellow"
                    />
                )}
                {data.solvedStatus === 3 && (
                    <FontAwesomeIcon
                        icon="fa-check"
                        className="w-4 h-4 highlight-green"
                    />
                )}
            </div>
            <Link className="link no-underline truncate" to={data.href}>
                {data.title}
            </Link>
        </span>
    );
}

function Archive({ data }) {
    const categoriesContent = data.categories.map((item, index) => (
        <DropdownListFrame
            key={index}
            title={item.title}
            tree={{ children: item.children }}
            leaf={ProblemLeaf}
        />
    ));
    return (
        <div className="w-full flex justify-center">
            <div className="flex justify-center w-full max-w-7xl lg:space-x-3 px-4">
                <ProfileSideBar />
                <div className="w-full min-w-0 space-y-3">
                    {categoriesContent}
                </div>
            </div>
        </div>
    );
}

export default Archive;
