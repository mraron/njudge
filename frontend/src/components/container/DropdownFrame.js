import { useState } from "react";
import { SVGDropdownFilterArrow } from "../svg/SVGs";
import RoundedFrame from "./RoundedFrame";

function DropdownFrame({ children, title }) {
    const [isOpen, setOpen] = useState(false);
    return (
        <RoundedFrame>
            <button
                onClick={() => setOpen(!isOpen)}
                className={`w-full h-12 ${
                    isOpen
                        ? "bg-grey-750 hover:bg-grey-725 rounded-tl-md rounded-tr-md"
                        : "bg-grey-800 hover:bg-grey-775 rounded-md"
                } border-bordercol flex items-center justify-center`}>
                <span className="font-medium mr-[0.3rem]">{title}</span>
                <SVGDropdownFilterArrow isOpen={isOpen} />
            </button>
            <div className={`${isOpen ? "" : "h-0 overflow-hidden"}`}>
                <div className="border-t border-bordercol">{children}</div>
            </div>
        </RoundedFrame>
    );
}

export default DropdownFrame;
