import {SVGDoubleRightArrow} from "../../svg/SVGs"
import RoundedFrame from "../container/RoundedFrame"
import {useLocation, useNavigate} from "react-router-dom";
import UpdateQueryString from "../../util/updateQueryString";

function Pagination({paginationData}) {
    const {currentPage, lastPage} = paginationData
    const location = useLocation()
    const navigate = useNavigate()
    const handlePageChanged = (page) => {
        UpdateQueryString(location, navigate, ["page"], [page])
    }
    const cls = "px-3 py-1.5 text-sm transition duration-200 border-default border hover:bg-grey-750 text-center"
    return (
        <RoundedFrame>
            <div className="flex justify-center p-3 overflow-x-auto">
                <button aria-label="First" className={`${cls} border-r-0 rounded-l-md`} onClick={() => handlePageChanged(1)}>
                    <SVGDoubleRightArrow cls="w-[1.32rem] h-[1.32rem] rotate-180"/>
                </button>
                {currentPage >= 3 && <button className={`${cls} hidden lg:block border-r-0`}
                                             onClick={() => handlePageChanged(currentPage - 2)}>{currentPage - 2}</button>}
                {currentPage >= 2 && <button className={`${cls} border-r-0`}
                                             onClick={() => handlePageChanged(currentPage - 1)}>{currentPage - 1}</button>}
                <button
                    className="px-3 py-1.5 text-sm bg-grey-700 border-1 border-grey-600 hover:border-grey-550 hover:bg-grey-650 transition duration-200 text-center">{currentPage}</button>
                {currentPage <= lastPage - 1 && <button className={`${cls} border-l-0`}
                                                        onClick={() => handlePageChanged(currentPage + 1)}>{currentPage + 1}</button>}
                {currentPage <= lastPage - 2 && <button className={`${cls} hidden lg:block border-l-0`}
                                                        onClick={() => handlePageChanged(currentPage + 2)}>{currentPage + 2}</button>}
                <button aria-label="Last" className={`${cls} border-l-0 rounded-r-md`} onClick={() => handlePageChanged(lastPage)}>
                    <SVGDoubleRightArrow cls="w-[1.32rem] h-[1.32rem] rotate-0"/>
                </button>
            </div>
        </RoundedFrame>
    )
}

export default Pagination