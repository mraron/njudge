import { Link } from "react-router-dom"
import ProfileSideBar from "../components/concrete/other/ProfileSidebar"
import DropdownListFrame from "../components/container/DropdownListFrame"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import React from "react"
import { SVGEllipsis } from "../components/svg/SVGs"
import { useTranslation } from "react-i18next"

function ProblemLeaf({ data }) {
    const { t } = useTranslation()
    return (
        <span className="ml-2 max-w-fit flex items-center">
            <div className="w-4 mr-2 flex justify-center items-center">
                {data.solvedStatus === 0 && (
                    <SVGEllipsis cls="w-4 h-4 text-grey-150" title={t("solved_status.not_tried")} />
                )}
                {data.solvedStatus === 1 && (
                    <FontAwesomeIcon
                        icon="fa-xmark"
                        className="w-4 h-4 highlight-red"
                        title={t("solved_status.wrong")}
                    />
                )}
                {data.solvedStatus === 2 && (
                    <FontAwesomeIcon
                        icon="fa-check"
                        className="w-4 h-4 highlight-yellow"
                        title={t("solved_status.partially_correct")}
                    />
                )}
                {data.solvedStatus === 3 && (
                    <FontAwesomeIcon
                        icon="fa-check"
                        className="w-4 h-4 highlight-green"
                        title={t("solved_status.correct")}
                    />
                )}
            </div>
            <Link className="link no-underline truncate" to={data.href}>
                {data.title}
            </Link>
        </span>
    )
}

function Archive({ data }) {
    const categoriesContent = data.categories.map((item, index) => (
        <DropdownListFrame key={index} title={item.title} tree={{ children: item.children }} leaf={ProblemLeaf} />
    ))
    return (
        <div className="w-full flex justify-center">
            <div className="flex justify-center w-full max-w-7xl lg:space-x-3 px-3">
                <ProfileSideBar />
                <div className="w-full min-w-0 space-y-3">{categoriesContent}</div>
            </div>
        </div>
    )
}

export default Archive
