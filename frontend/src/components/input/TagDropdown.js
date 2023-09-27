import _ from "lodash";
import { useEffect, useState } from "react";
import TextBoxDropdown from "./TextBoxDropdown";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

function Tag({ title, onClick }) {
    const [hovered, setHovered] = useState(false);
    const handleClick = (event) => {
        event.stopPropagation();
        onClick();
    };
    return (
        <span
            className={`tag flex items-center ${
                hovered
                    ? "text-white hover:bg-red-500 dark:hover:bg-red-600 hover:border-transparent"
                    : "bg-grey-725"
            }`}>
            {title}
            <span
                className={`flex ml-3 rounded-full p-1 hover:bg-red-700 dark:hover:bg-red-800`}
                onMouseOver={() => setHovered(true)}
                onMouseLeave={() => setHovered(false)}
                onClick={handleClick}>
                <FontAwesomeIcon icon="fa-close" className="h-3 w-3" />
            </span>
        </span>
    );
}

function TagDropdown({ id, label, itemNames, initTags = [], onChange }) {
    const [tags, setTags] = useState(initTags);
    useEffect(() => {
        if (onChange) {
            onChange(tags);
        }
    }, [onChange, tags]);

    const tagsContent = tags.map((tag, index) => {
        const handleRemoveTag = () => {
            setTags((prevTags) =>
                prevTags.filter((prevTag) => prevTag !== tag),
            );
        };
        return (
            <Tag title={itemNames[tag]} onClick={handleRemoveTag} key={index} />
        );
    });
    const handleAddTag = (selected, title) => {
        const remaining = _.range(itemNames.length).filter(
            (index) => !tags.includes(index),
        );
        const tag = remaining[selected];
        setTags((prevTags) =>
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
