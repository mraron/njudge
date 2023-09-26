import { Link } from "react-router-dom";
import ProfileSideBar from "../components/concrete/other/ProfileSidebar";
import DropdownListFrame from "../components/container/DropdownListFrame";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import React from "react";

function ProblemLeaf({ data }) {
    return (
        <span className="w-fit flex items-center">
            <div className="w-4 mr-2">
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
            <Link className="link no-underline" to={data.link}>
                {data.title}
            </Link>
        </span>
    );
}

function Archive({ data }) {
    const categoriesContent = data.categories.map((item, index) => (
        <div className="mb-3" key={index}>
            <DropdownListFrame
                title={item.title}
                tree={{ children: item.children }}
                leaf={ProblemLeaf}
            />
        </div>
    ));
    return (
        <div className="relative w-full flex justify-center">
            <div className="flex justify-center w-full max-w-7xl">
                <div className="ml-0 lg:ml-4">
                    <ProfileSideBar />
                </div>
                <div className="w-full px-4 lg:pl-3">{categoriesContent}</div>
            </div>
        </div>
    );
}

export default Archive;
