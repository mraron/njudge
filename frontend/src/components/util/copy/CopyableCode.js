import React from "react";
import CopyButton from "./CopyButton";

function CopyableCode({text, maxHeight, titleComponent}) {
    maxHeight ||= "auto"
    return (
        <div className="bg-grey-800 border-1 rounded-md flex flex-col border-default w-full">
            {titleComponent}
            <div className={`relative w-full code-default rounded-md flex`} style={{maxHeight: maxHeight, minHeight: "3.75rem"}}>
                <pre className="w-full px-4 py-3 pr-16 overflow-x-auto text-sm">
                    {text}
                </pre>
                <div className="absolute top-2.5 right-2.5 opacity-75 hover:opacity-100 transition duration-200">
                    <CopyButton text={text} />
                </div>
            </div>
        </div>
    )
}

export default CopyableCode