import { useEffect, useState } from "react";
import { SVGClose } from "../svg/SVGs";
import TextBoxDropdown from "./TextBoxDropdown";

function Tag({ title, onClick }) {
    const [hovered, setHovered] = useState(false);
    const handleClick = (event) => {
        event.stopPropagation();
        onClick();
    };
    return (
        <span
            className={`whitespace-nowrap flex items-center cursor-pointer text-sm px-2 py-1 border-1 rounded m-1 hover:bg-grey-700 ${
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

function TagDropdown({ id, label, itemNames, initTags, onChange }) {
    const [tags, setTags] = useState(initTags || []);
    useEffect(() => onChange(tags), [onChange, tags]);

    const tagsContent = tags.map((title, index) => {
        const handleRemoveTag = () => {
            setTags((prevTitles) =>
                prevTitles.filter((tagTitle) => tagTitle !== title),
            );
        };
        return <Tag title={title} onClick={handleRemoveTag} key={index} />;
    });
    const handleAddTag = (selected, title) => {
        setTags((prevTitles) =>
            prevTitles.includes(title) ? prevTitles : [...prevTitles, title],
        );
    };
    return (
        <div>
            <TextBoxDropdown
                id={id}
                label={label}
                itemNames={itemNames.filter(
                    (itemName) => !tags.includes(itemName),
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
