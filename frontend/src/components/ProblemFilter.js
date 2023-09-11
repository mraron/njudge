import React, { useState } from 'react';
import TextBox from './TextBox';
import TextBoxDropdown from './TextBoxDropdown';
import TagDropdown from './TagDropdown';
import { useLocation } from 'react-router-dom';
import { useNavigate } from 'react-router-dom';
import DropdownFrame from "./DropdownFrame";

function ProblemFilter() {
    const [title, setTitle] = useState("");
    const [tags, setTags] = useState([]);
    const [category, setCategory] = useState([-1, ""]);
    const location = useLocation();
    const navigate = useNavigate();

    const handleTitleChange = (newText) => {
        setTitle(newText);
    };
    const handleCategoryChange = (selected, newText) => {
        setCategory([selected, newText]);
    };
    const handleTagsChange = (tags) => {
        setTags(tags);
    };
    const handleSubmit = () => {
        const queryTitle = encodeURIComponent(title);
        const queryCategory = encodeURIComponent(category[0]);
        const queryTags = encodeURIComponent(tags.join(','));
        const queryString = `?title=${queryTitle}&category=${queryCategory}&tags=${queryTags}`;
        navigate(`${location.pathname}${queryString}`);
    };
    return (
        <div className="w-full">
            <div className="mb-4">
                <TextBox id="filterTitle" label="Feladatcím" initText={title} onChange={handleTitleChange} />
            </div>
            <div className="mb-4">
                <TagDropdown id="filterTags" label="Címkék" fillSelected={false} itemNames={[
                    "matematika",
                    "mohó",
                    "dinamikus programozás",
                    "adatszerkezetek",
                ]} initTags={tags} onChange={handleTagsChange} />
            </div>
            <div className="mb-5">
                <TextBoxDropdown id="filterCategory" label="Kategória" initText={category[1]} initSelected={category[0]} fillSelected={true} itemNames={[
                    "IOI-CEOI Válogató 2023",
                    "IOI-CEOI Válogató 2023 − 1. forduló",
                    "IOI-CEOI Válogató 2023 − 2. forduló",
                    "IOI-CEOI Válogató 2023 − 3. forduló"
                ]} onChange={handleCategoryChange} />
            </div>
            <div className="flex justify-center">
                <button className="mr-1 btn-indigo w-32" onClick={handleSubmit}>Keres</button>
                <button className="ml-1 btn-gray w-32">Visszaállít</button>
            </div>
        </div>
    )
}

export function ProblemFilterFrame() {
    return (
        <DropdownFrame>
            <div className="px-8 py-6 border-t border-default">
                <ProblemFilter />
            </div>
        </DropdownFrame>
    )
}

export default ProblemFilter;