import { SVGDoubleLeftArrow, SVGLeftArrow } from "../svg/SVGs"
import RoundedFrame from "./RoundedFrame"

function Pagination({ current, last }) {
    return (
        <RoundedFrame>
            <div className="flex justify-center p-3 overflow-x-auto">
                <button className="px-3 py-1.5 text-sm transition duration-200 border-default border border-r-0 hover:bg-grey-750 text-center rounded-l-md">
                    <SVGDoubleLeftArrow cls="w-4 h-4" />
                </button>
                <button className="px-3 py-1.5 text-sm transition duration-200 border-default border border-r-0 hover:bg-grey-750 text-center">
                    <SVGLeftArrow cls="w-3 h-3" />
                </button>
                {current >= 3 && <button className="hidden lg:block px-3 py-1.5 text-sm transition duration-200 border-default border border-r-0 hover:bg-grey-750 text-center">{current - 2}</button>}
                {current >= 2 && <button className="px-3 py-1.5 text-sm transition duration-200 border-default border border-r-0 hover:bg-grey-750 text-center">{current - 1}</button>}
                <button className="px-3 py-1.5 text-sm font-medium bg-indigo-600 border-indigo-600 hover:bg-indigo-500 hover:border-indigo-500 transition duration-200 text-center">{current}</button>
                {current <= last - 1 && <button className="px-3 py-1.5 text-sm transition duration-200 border-default border border-l-0 hover:bg-grey-750 text-center">{current + 1}</button>}
                {current <= last - 2 && <button className="hidden lg:block px-3 py-1.5 text-sm transition duration-200 border-default border border-l-0 hover:bg-grey-750 text-center">{current + 2}</button>}
                <button className="px-3 py-1.5 text-sm transition duration-200 border-default border border-l-0 hover:bg-grey-750 text-center">
                    <SVGLeftArrow cls="w-3 h-3 rotate-180" />
                </button>
                <button className="px-3 py-1.5 text-sm transition duration-200 border-default border border-l-0 hover:bg-grey-750 text-center rounded-r-md">
                    <SVGDoubleLeftArrow cls="w-4 h-4 rotate-180" />
                </button>
            </div>
        </RoundedFrame>
    )
}

export default Pagination