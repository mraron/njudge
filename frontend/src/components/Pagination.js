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
    const cls = "px-3 py-1.5 text-sm transition duration-200 border-default border hover:bg-grey-750 text-center"
    return (
        <RoundedFrame>
            <div className="flex justify-center p-3 overflow-x-auto">
                <button className={`${cls} border-r-0 rounded-l-md`} onClick={() => setPage(1)}>
                    <SVGDoubleLeftArrow cls="w-4 h-4" />
                </button>
                {current >= 2 &&
                    <button className={`${cls} border-r-0`} onClick={() => setPage(current - 1)}>
                        <SVGLeftArrow cls="w-3 h-3" />
                    </button>}
                {current >= 3 && <button className={`${cls} hidden lg:block border-r-0`} onClick={() => setPage(current - 2)}>{current - 2}</button>}
                {current >= 2 && <button className={`${cls} border-r-0`} onClick={() => setPage(current - 1)}>{current - 1}</button>}
                <button className="px-3 py-1.5 text-sm font-medium bg-indigo-600 border-indigo-600 hover:bg-indigo-500 hover:border-indigo-500 transition duration-200 text-center">{current}</button>
                {current <= last - 1 && <button className={`${cls} border-l-0`} onClick={() => setPage(current + 1)}>{current + 1}</button>}
                {current <= last - 2 && <button className={`${cls} hidden lg:block border-l-0`} onClick={() => setPage(current + 2)}>{current + 2}</button>}
                {current <= last - 1 &&
                    <button className={`${cls} border-l-0`} onClick={() => setPage(current + 1)}>
                        <SVGLeftArrow cls="w-3 h-3 rotate-180" />
                    </button>}
                <button className={`${cls} border-l-0 rounded-r-md`} onClick={() => setPage(last)}>
                    <SVGDoubleLeftArrow cls="w-4 h-4 rotate-180" />
                </button>
            </div>
        </RoundedFrame>
    )
}

export default Pagination