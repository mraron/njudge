import React from "react";
import {SVGWrongSimple} from "../../../svg/SVGs";
import SVGTitleComponent from "../../../svg/SVGTitleComponent";
import RoundedFrame from "../../container/RoundedFrame";
import CopyableCode from "../../util/copy/CopyableCode";

function CompileErrorFrame({message}) {
    const titleComponent =
        <SVGTitleComponent title="Fordítási hiba" svg={<SVGWrongSimple cls="w-6 h-6 mr-2 text-red-500" />} />

    return (
        <CopyableCode text={message} titleComponent={titleComponent} maxHeight="16rem"/>
    )
}

export default CompileErrorFrame