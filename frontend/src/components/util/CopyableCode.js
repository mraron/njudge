import {SVGCopy} from "../../svg/SVGs";
import React from "react";

function CopyableCode({text}) {
    const handleCopy = () => {
        navigator.clipboard.writeText(text)
        window.flash("info.successful_copy", "success")
    }
    return (
        <div className="w-full flex items-center">
            <button
                className="h-9 w-9 mr-2 rounded-md border-1 bg-grey-800 border-grey-725 hover:bg-grey-775 transition duration-200 relative"
                aria-label="Copy" onClick={handleCopy}>
                <SVGCopy
                    cls={`absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 w-4 h-4 `}/>
            </button>
            <div className="w-full flex items-center px-3 py-2 border-1 border-grey-725 rounded-md bg-grey-875 font-mono">
                {text}
            </div>
        </div>
    )
}

export default CopyableCode