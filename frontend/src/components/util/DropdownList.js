import { useState } from "react";
import { SVGDropdownListArrow } from "../svg/SVGs";

function DropdownElem({ text, isOpen, onClick }) {
    return (
        <span
            className="max-w-fit flex link items-center dark:font-medium no-underline"
            onClick={onClick}>
            <SVGDropdownListArrow isOpen={isOpen} />
            <span className="truncate">{text}</span>
        </span>
    );
}

function DropdownList({ tree, leaf: Leaf }) {
    const isRoot = !tree.title;
    const [isOpen, setOpen] = useState(isRoot);

    const children = tree.children
        ? tree.children.map((child, index) => (
              <li className="mt-2" key={index}>
                  <DropdownList tree={child} leaf={Leaf} />
              </li>
          ))
        : [];
    const innerNode = (
        <DropdownElem
            text={tree.title}
            isOpen={isOpen}
            onClick={() => setOpen(!isOpen)}
        />
    );
    const leafNode = <Leaf data={tree} />;
    const isLeaf = !tree.children || tree.children.length === 0;

    return (
        <div>
            {!isRoot && !isLeaf && innerNode}
            {!isRoot && isLeaf && leafNode}
            <ul
                className={`${isOpen ? "" : "hidden"} ${
                    isRoot ? "" : "ml-6"
                } mb-4`}>
                {children}
            </ul>
        </div>
    );
}

export default DropdownList;
