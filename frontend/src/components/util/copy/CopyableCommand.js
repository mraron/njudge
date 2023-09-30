import CopyButton from "./CopyButton";
import RoundedFrame from "../../container/RoundedFrame";
import { useState } from "react";

function CopyableCommand({ text, cls }) {
    const [isHovered, setHovered] = useState(false);
    return (
        <RoundedFrame cls={`${cls} overflow-auto shadow-none`}>
            <div
                className="p-0.5 relative bg-codebgcol flex"
                onMouseEnter={() => setHovered(true)}
                onMouseLeave={() => setHovered(false)}>
                <div className="flex items-center w-full px-4 py-3 overflow-x-auto whitespace-nowrap font-code">
                    {text}
                </div>
                <div className="absolute right-1.5 top-1/2 transform -translate-y-1/2 duration-200">
                    <CopyButton text={text} isVisible={isHovered} />
                </div>
            </div>
        </RoundedFrame>
    );
}

export default CopyableCommand;
