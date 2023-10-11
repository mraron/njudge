import { useLocation, useNavigate } from "react-router-dom";
import RoundedFrame from "../container/RoundedFrame";
import UpdateQueryString from "../../util/updateQueryString";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { useTranslation } from "react-i18next";

function Pagination({ paginationData }) {
    const { t } = useTranslation()
    const { currentPage, lastPage } = paginationData
    const location = useLocation()
    const navigate = useNavigate()
    const handlePageChanged = (page) => {
        UpdateQueryString({
            location: location,
            navigate: navigate,
            args: ["page"],
            values: [page],
        })
    }
    const cls = "flex justify-center items-center px-3 py-1.5 text-sm border-bormedcol border hover:bg-grey-775"
    return (
        <RoundedFrame cls="overflow-hidden">
            <div className="flex justify-center p-4 overflow-x-auto">
                <button
                    aria-label={t("aria_label.first")}
                    className={`${cls} border-r-0 rounded-l-md`}
                    onClick={() => handlePageChanged(1)}>
                    <FontAwesomeIcon icon="fa-angles-left" className="w-2.5 h-2.5" />
                </button>
                {currentPage >= 3 && (
                    <button
                        className={`${cls} hidden lg:block border-r-0`}
                        onClick={() => handlePageChanged(currentPage - 2)}>
                        {currentPage - 2}
                    </button>
                )}
                {currentPage >= 2 && (
                    <button className={`${cls} border-r-0`} onClick={() => handlePageChanged(currentPage - 1)}>
                        {currentPage - 1}
                    </button>
                )}
                <button className="px-3 py-1.5 text-sm bg-btncol border border-bormedcol text-center">
                    {currentPage}
                </button>
                {currentPage <= lastPage - 1 && (
                    <button className={`${cls} border-l-0`} onClick={() => handlePageChanged(currentPage + 1)}>
                        {currentPage + 1}
                    </button>
                )}
                {currentPage <= lastPage - 2 && (
                    <button
                        className={`${cls} hidden lg:block border-l-0`}
                        onClick={() => handlePageChanged(currentPage + 2)}>
                        {currentPage + 2}
                    </button>
                )}
                <button
                    aria-label={t("aria_label.last")}
                    className={`${cls} border-l-0 rounded-r-md`}
                    onClick={() => handlePageChanged(lastPage)}>
                    <FontAwesomeIcon icon="fa-angles-right" className="w-2.5 h-2.5" />
                </button>
            </div>
        </RoundedFrame>
    )
}

export default Pagination
