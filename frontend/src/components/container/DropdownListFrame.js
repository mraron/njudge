import { useState } from "react"
import { SVGDropdownListArrow } from "../svg/SVGs"
import RoundedFrame from "./RoundedFrame"

function DropdownElem({ text, isOpen, onClick }) {
    return (
        <span
            className="max-w-fit flex link items-center emph-med no-underline"
            onClick={onClick}>
            <SVGDropdownListArrow isOpen={isOpen} />
            <span className="truncate">{text}</span>
        </span>
    )
}

function DropdownList({ tree, leaf: Leaf }) {
    const isRoot = !tree.title
    const [isOpen, setOpen] = useState(isRoot)

    const children = tree.children
        ? tree.children.map((child, index) => (
              <li className="mt-2" key={index}>
                  <DropdownList tree={child} leaf={Leaf} />
              </li>
          ))
        : []
    const innerNode = (
        <DropdownElem
            text={tree.title}
            isOpen={isOpen}
            onClick={() => setOpen(!isOpen)}
        />
    )
    const leafNode = <Leaf data={tree} />
    const isLeaf = !tree.children || tree.children.length === 0

    return (
        <div>
            {!isRoot && !isLeaf && innerNode}
            {!isRoot && isLeaf && leafNode}
            <ul
                className={`${isOpen ? "" : "hidden"} ${
                    isRoot ? "" : "ml-5"
                } mb-4`}>
                {children}
            </ul>
        </div>
    )
}

function DropdownListFrame({ title, tree, leaf: Leaf }) {
    return (
        <RoundedFrame title={title}>
            <div className="px-8 pt-4 pb-2 bg-grey-850 rounded-b-container">
                <DropdownList tree={tree} leaf={Leaf} />
            </div>
        </RoundedFrame>
    )
}

export default DropdownListFrame
