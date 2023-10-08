import CopyButton from "./CopyButton"
import RoundedFrame from "../../container/RoundedFrame"
import { useState } from "react"

function CopyableCode({ text, maxHeight = "auto", cls }) {
    const [isHovered, setHovered] = useState(false)
    return (
        <RoundedFrame cls={`${cls} overflow-auto shadow-none bg-codebgcol`}>
            <div
                onMouseEnter={() => setHovered(true)}
                onMouseLeave={() => setHovered(false)}
                className="relative flex"
                style={{ minHeight: "3rem", maxHeight: maxHeight }}>
                <pre className="w-full px-4 py-3 overflow-x-auto">{text}</pre>
                <div className="absolute top-1.5 right-1.5">
                    <CopyButton text={text} isVisible={isHovered} />
                </div>
            </div>
        </RoundedFrame>
    )
}

export default CopyableCode
