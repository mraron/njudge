import { useState } from "react"
import { DropdownListArrow } from "../svg/SVGs";
import { Link } from "react-router-dom";

function DropdownList({ tree }) {
    const isRoot = !tree["title"];
    const [isOpen, setOpen] = useState(isRoot);

    const children = tree["children"]? tree["children"].map((child, index) => 
        <li className="mt-2" key={index}><DropdownList tree={child} /></li>
    ): []
    const arrow = <DropdownListArrow isOpen={isOpen} />
    const innerNode = 
        <a className="w-fit flex items-center cursor-pointer hover:text-indigo-300 font-medium transition-all duration-100" onClick={() => setOpen(!isOpen)} >
            {arrow}{tree["title"]}
        </a>
    const leafNode = 
        <a className="w-fit flex items-center cursor-pointer hover:text-indigo-300 transition-all duration-100">
            <Link to={tree["link"]}>{tree["title"]}</Link>
        </a>
    const isLeaf = !tree["children"] || tree["children"].length === 0
        
    return (
        <div>
            {!isRoot && !isLeaf && innerNode}
            {!isRoot && isLeaf && leafNode}
            <ul className={`${isOpen? "": "hidden"} ${isRoot? "": "ml-8"} mb-4`}>
                {children}
            </ul>
        </div>
    )
}

export default DropdownList;