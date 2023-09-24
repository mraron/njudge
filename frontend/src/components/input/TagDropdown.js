import { useEffect, useState } from "react";
import { SVGClose } from "../svg/SVGs";
import TextBoxDropdown from "./TextBoxDropdown";
import _ from "lodash";

function Tag({ title, onClick }) {
    const [hovered, setHovered] = useState(false);
    const handleClick = (event) => {
        event.stopPropagation();
        onClick();
    };
    return (
        <span
            className={`whitespace-nowrap flex items-center cursor-pointer text-sm px-2 py-1 border rounded m-1 hover:bg-grey-700 ${
                hovered
                    ? "hover:bg-red-600 hover:border-red-500"
                    : "bg-grey-725 border-grey-650"
            } transition-all duration-200`}>
            {title}
            <span
                className={`ml-3 rounded-full p-1 hover:bg-red-800 transition-all duration-200`}
                onMouseOver={() => setHovered(true)}
                onMouseLeave={() => setHovered(false)}
                onClick={handleClick}>
                <SVGClose cls="h-2 w-2" />
            </span>
        </span>
    );
}

function TagDropdown({ id, label, itemNames, initTags = [], onChange }) {
    const [tags, setTags] = useState(initTags);
    useEffect(() => onChange(tags), [onChange, tags]);

    const tagsContent = tags.map((tag, index) => {
        const handleRemoveTag = () => {
            setTags(prevTags =>
                prevTags.filter(prevTag => prevTag !== tag),
            );
        };
        return <Tag title={itemNames[tag]} onClick={handleRemoveTag} key={index} />;
    });
    const handleAddTag = (selected, title) => {
        const remaining = _.range(itemNames.length).filter(
            index => !tags.includes(index),
        )
        const tag = remaining[selected]
        setTags(prevTags =>
            prevTags.includes(tag) ? prevTags : [...prevTags, tag],
        );
    };
    return (
        <div>
            <TextBoxDropdown
                id={id}
                label={label}
                itemNames={itemNames.filter(
                    (itemName, index) => !tags.includes(index),
                )}
                onClick={handleAddTag}
            />
            <div
                className={`${
                    tagsContent.length > 0 ? "mt-2" : ""
                } flex flex-wrap`}>
                {tagsContent}
            </div>
        </div>
    );
}

export default TagDropdown;
