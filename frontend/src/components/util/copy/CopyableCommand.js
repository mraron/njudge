import React from "react";
import CopyButton from "./CopyButton";

function CopyableCommand({text}) {
    return (
        <div className="flex items-start">
            <div className="mr-2">
                <CopyButton text={text} />
            </div>
            <div className={`w-full code-default border border-default rounded-md flex items-center`}>
                <div className="w-full px-3 py-2 font-mono whitespace-nowrap text-sm">
                    {text}
                </div>
            </div>
        </div>
    )
}

export default CopyableCommand