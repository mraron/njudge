import { useEffect, useState } from "react"
import TextBox from "./TextBox"

function DropdownItem({ text, onClick }) {
    return (
        <li
            className="cursor-pointer px-4 py-2.5 flex items-center hover:bg-grey-825 border-grey-750"
            onMouseDown={onClick}>
            {text}
        </li>
    )
}

function TextBoxDropdown({ id, label, items, fillSelected, initText = "", initSelected = -1, onChange, onClick }) {
    const [focused, setFocused] = useState(false)
    const [selected, setSelected] = useState(initSelected)
    const [text, setText] = useState(initText)

    useEffect(() => {
        onChange?.(selected, text)
    }, [selected, text])

    const handleFocus = () => {
        setFocused(true)
    }
    const handleBlur = () => {
        setFocused(false)
    }
    const handleTextChange = (newText) => {
        setSelected(items.map((item) => item.toLowerCase()).indexOf(newText.toLowerCase()))
        setText(newText)
    }
    const itemsContent = items
        .filter((item) => item.toLowerCase().includes(text.toLowerCase()))
        .map((item, index) => {
            const handleClick = () => {
                if (fillSelected) {
                    setText(item)
                }
                setSelected(index)
                onClick?.(index, item)
            }
            return <DropdownItem text={item} onClick={handleClick} key={index} />
        })
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
            <div className={`z-10 absolute overflow-hidden inset-x-0 ${focused ? "max-h-52" : "max-h-0"}`}>
                <div
                    className={`rounded-sm max-h-52 overflow-y-auto border-border-def ${
                        items.length > 0 ? "border" : ""
                    }`}>
                    <ul className={`divide-y divide-divide-def bg-grey-875 text-sm`}>{itemsContent}</ul>
                </div>
            </div>
        </div>
    )
}

export default TextBoxDropdown
