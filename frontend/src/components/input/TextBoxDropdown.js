import { useEffect, useState } from "react";
import TextBox from "./TextBox";

function DropdownItem({ itemName, onClick }) {
    return (
        <li
            className="cursor-pointer px-4 py-2 flex items-center hover:bg-framebgcol border-grey-750"
            onMouseDown={onClick}>
            {itemName}
        </li>
    );
}

function TextBoxDropdown({
    id,
    label,
    itemNames,
    fillSelected,
    initText = "",
    initSelected = -1,
    onChange,
    onClick,
}) {
    const [focused, setFocused] = useState(false);
    const [selected, setSelected] = useState(initSelected);
    const [text, setText] = useState(initText);

    useEffect(() => {
        if (onChange) onChange(selected, text);
    }, [selected, text]);

    const handleFocus = () => {
        setFocused(true);
    };
    const handleBlur = () => {
        setFocused(false);
    };
    const handleTextChange = (newText) => {
        setSelected(
            itemNames
                .map((itemName) => itemName.toLowerCase())
                .indexOf(newText.toLowerCase()),
        );
        setText(newText);
    };
    const items = itemNames
        .filter((itemName) =>
            itemName.toLowerCase().includes(text.toLowerCase()),
        )
        .map((itemName, index) => {
            const handleClick = () => {
                if (fillSelected) {
                    setText(itemName);
                }
                setSelected(index);
                if (onClick) {
                    onClick(index, itemName);
                }
            };
            return (
                <DropdownItem
                    itemName={itemName}
                    onClick={handleClick}
                    key={index}
                />
            );
        });
    return (
        <div className="relative">
            <TextBox
                id={id}
                label={label}
                initText={text}
                onChange={handleTextChange}
                onFocus={handleFocus}
                onBlur={handleBlur}
            />
            <div
                className={`z-10 absolute overflow-hidden inset-x-0 ${
                    focused ? "max-h-60" : "max-h-0"
                }`}>
                <div
                    className={`rounded-sm max-h-60 overflow-y-auto border-bordercol ${
                        items.length > 0 ? "border" : ""
                    }`}>
                    <ul className={`divide-y divide-dividecol bg-grey-875`}>
                        {items}
                    </ul>
                </div>
            </div>
        </div>
    );
}

export default TextBoxDropdown;
