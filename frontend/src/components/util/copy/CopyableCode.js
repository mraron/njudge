import CopyButton from "./CopyButton"
import RoundedFrame from "../../container/RoundedFrame"
import { useState } from "react"

function CopyableCode({ text, maxHeight = "auto", titleComponent, cls }) {
    const [isHovered, setHovered] = useState(false)
    return (
        <RoundedFrame cls={`${cls} overflow-auto shadow-none`} titleComponent={titleComponent}>
            <div
                onMouseEnter={() => setHovered(true)}
                onMouseLeave={() => setHovered(false)}
                className="relative bg-codebgcol flex"
                style={{ minHeight: "3.75rem", maxHeight: maxHeight }}>
                <pre className="w-full px-4 py-3 overflow-x-auto">{text}</pre>
                <div className="absolute top-2.5 right-2.5 duration-200">
                    <CopyButton text={text} isVisible={isHovered} />
                </div>
            </div>
        </RoundedFrame>
    )
}

export default CopyableCode
