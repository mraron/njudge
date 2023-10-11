import { useState } from "react"
import { SVGDropdownFrameArrow } from "../svg/SVGs"
import RoundedFrame from "./RoundedFrame"

function DropdownFrame({ children, title }) {
    const [isOpen, setOpen] = useState(false)
    return (
        <RoundedFrame>
            <button
                onClick={() => setOpen(!isOpen)}
                className={`w-full px-8 py-2.5 ${
                    isOpen
                        ? "bg-grey-750 hover:bg-grey-725 rounded-t-container"
                        : "bg-framebgcol hover:bg-grey-775 rounded-container"
                } border-bordefcol flex items-center justify-center`}>
                <span className="emph-med mr-[0.3rem] truncate">{title}</span>
                <SVGDropdownFrameArrow isOpen={isOpen} />
            </button>
            <div className={`${isOpen ? "" : "h-0 overflow-hidden"}`}>
                <div className="border-t border-bordefcol">{children}</div>
            </div>
        </RoundedFrame>
    )
}

export default DropdownFrame
