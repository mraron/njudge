import { useEffect, useState } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import TextBoxDropdown from "./TextBoxDropdown";
import _ from "lodash";

function Tag({ title, onClick }) {
    const [hovered, setHovered] = useState(false)
    const handleClick = (event) => {
        event.stopPropagation()
        onClick()
    }
    return (
        <span
            className={`tag m-1 flex items-center space-x-3 ${
                hovered ? "text-white hover:bg-red-500 dark:hover:bg-red-600 hover:border-transparent" : ""
            }`}>
            <span>{title}</span>
            <span
                className={`flex rounded-full p-1 hover:bg-red-700 dark:hover:bg-red-800`}
                onMouseOver={() => setHovered(true)}
                onMouseLeave={() => setHovered(false)}
                onClick={handleClick}>
                <FontAwesomeIcon icon="fa-close" className="h-3 w-3" />
            </span>
        </span>
    )
}

function TagDropdown({ id, label, items, initTags = [], onChange }) {
    const [tags, setTags] = useState(initTags)

    useEffect(() => {
        onChange?.(tags)
    }, [onChange, tags])

    const tagsContent = tags.map((tag, index) => {
        const handleRemoveTag = () => {
            setTags((prevTags) => prevTags.filter((prevTag) => prevTag !== tag))
        }
        return <Tag title={items[tag]} onClick={handleRemoveTag} key={index} />
    })
    const handleAddTag = (selected, text) => {
        const remaining = _.range(items.length).filter((index) => !tags.includes(index))
        const tag = remaining[selected]
        setTags((prevTags) => (prevTags.includes(tag) ? prevTags : [...prevTags, tag]))
    }
    return (
        <div>
            <TextBoxDropdown
                id={id}
                label={label}
                items={items.filter((item, index) => !tags.includes(index))}
                onClick={handleAddTag}
            />
            <div className={`${tagsContent.length > 0 ? "mt-2" : ""} flex flex-wrap`}>{tagsContent}</div>
        </div>
    )
}

export default TagDropdown
