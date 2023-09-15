import {useState} from "react";
import {SVGCheckmark} from "../svg/SVGs";

function Checkbox({id, label}) {
    const [checked, setChecked] = useState(false)
    const [hovered, setHovered] = useState(false)
    return (
        <label htmlFor={id} className="flex items-start justify-center max-w-fit" onMouseOver={() => setHovered(true)}
               onMouseLeave={() => setHovered(false)}>
            <div
                className={`flex border items-center justify-center ${checked ? `border-transparent ${hovered ? "bg-indigo-500" : "bg-indigo-600"}` : `${hovered ? "bg-grey-800" : "bg-grey-850"} border-default`} w-5 h-5 rounded-sm transition duration-200`}>
                <SVGCheckmark cls={`w-3.5 h-3.5 ${checked ? "opacity-100" : "opacity-0"} transition duration-200`}/>
                <input id={id} onChange={(event) => setChecked(event.target.checked)} className={`appearance-none`}
                       type="checkbox"/>
            </div>
            <span className="text-label ml-3">{label}</span>
        </label>
    )
}

export default Checkbox;