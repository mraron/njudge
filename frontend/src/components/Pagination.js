import { SVGDoubleLeftArrow, SVGLeftArrow } from "../svg/SVGs"
import RoundedFrame from "./RoundedFrame"
import {useLocation, useNavigate} from "react-router-dom";
import queryString from "query-string";

function Pagination({ current, last }) {
    const location = useLocation()
    const navigate = useNavigate()
    const setPage = (page) => {
        const qStringOld = location.search
        const qData = queryString.parse(qStringOld)
        qData.page = page

        const qStringNew = queryString.stringify(qData)
        navigate(`${location.pathname}?${qStringNew}`)
    }
    return (
        <RoundedFrame>
            <div className="flex justify-center p-3 overflow-x-auto">
                <button className="px-3 py-1.5 text-sm transition duration-200 border-default border border-r-0 hover:bg-grey-750 text-center rounded-l-md" onClick={() => setPage(1)}>
                    <SVGDoubleLeftArrow cls="w-4 h-4" />
                </button>
                {current >= 2 &&
                    <button className="px-3 py-1.5 text-sm transition duration-200 border-default border border-r-0 hover:bg-grey-750 text-center" onClick={() => setPage(current - 1)}>
                        <SVGLeftArrow cls="w-3 h-3" />
                    </button>}
                {current >= 3 && <button className="hidden lg:block px-3 py-1.5 text-sm transition duration-200 border-default border border-r-0 hover:bg-grey-750 text-center" onClick={() => setPage(current - 2)}>{current - 2}</button>}
                {current >= 2 && <button className="px-3 py-1.5 text-sm transition duration-200 border-default border border-r-0 hover:bg-grey-750 text-center" onClick={() => setPage(current - 1)}>{current - 1}</button>}
                <button className="px-3 py-1.5 text-sm font-medium bg-indigo-600 border-indigo-600 hover:bg-indigo-500 hover:border-indigo-500 transition duration-200 text-center">{current}</button>
                {current <= last - 1 && <button className="px-3 py-1.5 text-sm transition duration-200 border-default border border-l-0 hover:bg-grey-750 text-center" onClick={() => setPage(current + 1)}>{current + 1}</button>}
                {current <= last - 2 && <button className="hidden lg:block px-3 py-1.5 text-sm transition duration-200 border-default border border-l-0 hover:bg-grey-750 text-center" onClick={() => setPage(current + 2)}>{current + 2}</button>}
                {current <= last - 1 &&
                    <button className="px-3 py-1.5 text-sm transition duration-200 border-default border border-l-0 hover:bg-grey-750 text-center" onClick={() => setPage(current + 1)}>
                        <SVGLeftArrow cls="w-3 h-3 rotate-180" />
                    </button>}
                <button className="px-3 py-1.5 text-sm transition duration-200 border-default border border-l-0 hover:bg-grey-750 text-center rounded-r-md" onClick={() => setPage(last)}>
                    <SVGDoubleLeftArrow cls="w-4 h-4 rotate-180" />
                </button>
            </div>
        </RoundedFrame>
    )
}

export default Pagination