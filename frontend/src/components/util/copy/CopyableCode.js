import CopyButton from "./CopyButton"
import RoundedFrame from "../../container/RoundedFrame"
import { useState } from "react"

function CopyableCode({ text, maxHeight = "auto", title, titleComponent, cls }) {
    const [isHovered, setHovered] = useState(false)
    return (
        <RoundedFrame title={title} titleComponent={titleComponent} cls={`${cls} shadow-none`}>
            <div
                onMouseEnter={() => setHovered(true)}
                onMouseLeave={() => setHovered(false)}
                className="relative flex bg-code-bg"
                style={{ minHeight: "3rem", maxHeight: maxHeight }}>
                <pre className="w-full px-4 py-3 overflow-x-auto ">{text}</pre>
                <div className="absolute top-2 right-2">
                    <CopyButton text={text} isVisible={isHovered} />
                </div>
            </div>
        </RoundedFrame>
    )
}

export default CopyableCode
