import {useState} from "react"
import {SVGDropdownListArrow} from "../../svg/SVGs";

function DropdownElem({isOpen, text, onClick}) {
    return (
        <span
            className="w-fit flex items-center cursor-pointer hover:text-indigo-300 font-medium transition-all duration-100"
            onClick={onClick}>
            <SVGDropdownListArrow isOpen={isOpen}/>{text}
        </span>
    )
}

function DropdownList({tree, leaf: Leaf}) {
    const isRoot = !tree.title;
    const [isOpen, setOpen] = useState(isRoot);

    const children = tree.children ? tree.children.map((child, index) =>
        <li className="mt-2" key={index}><DropdownList tree={child} leaf={Leaf}/></li>
    ) : []
    const innerNode = <DropdownElem isOpen={isOpen} text={tree.title} onClick={() => setOpen(!isOpen)}/>
    const leafNode = <Leaf data={tree}/>
    const isLeaf = !tree.children || tree.children.length === 0

    return (
        <div>
            {!isRoot && !isLeaf && innerNode}
            {!isRoot && isLeaf && leafNode}
            <ul className={`${isOpen ? "" : "hidden"} ${isRoot ? "" : "ml-6"} mb-4`}>
                {children}
            </ul>
        </div>
    )
}

export default DropdownList;