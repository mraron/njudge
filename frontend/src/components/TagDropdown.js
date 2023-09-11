import React, { useState, useEffect } from 'react';
import TextBoxDropdown from './TextBoxDropdown';
import { SVGClose } from '../svg/SVGs';

function Tag({ title, onClick }) {
    const [hovered, setHovered] = useState(false);
    const handleClick = (event) => {
        event.stopPropagation();
        onClick();
    };
    return (
        <span className={`whitespace-nowrap flex items-center cursor-pointer text-sm px-2 py-1 border-1 rounded m-1 hover:bg-grey-700 ${hovered ? "hover:bg-red-600 hover:border-red-500" : "bg-grey-725 border-grey-650"} transition-all duration-200`}>
            {title}
            <span className={`ml-3 rounded-full p-1 hover:bg-red-800 transition-all duration-200`} onMouseOver={() => setHovered(true)} onMouseLeave={() => setHovered(false)} onClick={handleClick}>
                <SVGClose size={"h-2 w-2"} />
            </span>
        </span>
    );
}

function TagDropdown({ id, label, itemNames, initTags, onChange }) {
    const [tagTitles, setTagTitles] = useState(initTags || []);
    useEffect(() => onChange(tagTitles), [onChange, tagTitles]);
    
    const tags = tagTitles.map((title, index) => {
        const handleRemoveTag = () => {
            setTagTitles(prevTitles => prevTitles.filter(tagTitle => tagTitle !== title));
        };
        return (
            <Tag title={title} onClick={handleRemoveTag} key={index} />
        );
    });
    const handleAddTag = (selected, title) => {
        setTagTitles(prevTitles => prevTitles.includes(title)? prevTitles: prevTitles.concat([title]));
    };
    return (
        <div>
            <TextBoxDropdown id={id} label={label} itemNames={itemNames.filter(itemName => !tagTitles.includes(itemName))} onClick={handleAddTag} />
            <div className={`${tags.length > 0? "mt-2": ""} flex flex-wrap`}>
                {tags}
            </div>
        </div>
    )
}

export default TagDropdown;