import { useState } from "react";
import { SVGDropdownFrameArrow } from "../svg/SVGs";
import RoundedFrame from "./RoundedFrame";

function DropdownFrame({ children, title }) {
    const [isOpen, setOpen] = useState(false);
    return (
        <RoundedFrame>
            <button
                onClick={() => setOpen(!isOpen)}
                className={`w-full padding-btn-default ${
                    isOpen
                        ? "bg-grey-750 hover:bg-grey-725 rounded-t-container"
                        : "bg-framebgcol hover:bg-grey-775 rounded-container"
                } border-bordercol flex items-center justify-center`}>
                <span className="dark:font-medium mr-[0.3rem] truncate">
                    {title}
                </span>
                <SVGDropdownFrameArrow isOpen={isOpen} />
            </button>
            <div className={`${isOpen ? "" : "h-0 overflow-hidden"}`}>
                <div className="border-t border-bordercol">{children}</div>
            </div>
        </RoundedFrame>
    );
}

export default DropdownFrame;
