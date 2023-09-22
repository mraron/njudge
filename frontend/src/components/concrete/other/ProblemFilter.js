import React, {useState} from 'react';
import TextBox from '../../input/TextBox';
import TextBoxDropdown from '../../input/TextBoxDropdown';
import TagDropdown from '../../input/TagDropdown';
import {useLocation, useNavigate} from 'react-router-dom';
import DropdownFrame from "../../container/DropdownFrame";
import queryString from "query-string";
import {useTranslation} from "react-i18next";
import updateQueryString from "../../../util/updateQueryString";

function ProblemFilter() {
    const {t} = useTranslation()
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
        updateQueryString(location, navigate, ["title", "tags", "category"], [title, tags.join(","), category[0]], ["title", "tags", "category"])
    };
    const handleReset = () => {
        navigate(location.pathname)
    }
    return (
        <div className="w-full">
            <div className="mb-4">
                <TextBox id="filterTitle" label={t("problem_filter.title")} initText={title} onChange={handleTitleChange}/>
            </div>
            <div className="mb-4">
                <TagDropdown id="filterTags" label={t("problem_filter.tags")} fillSelected={false} itemNames={[
                    "matematika",
                    "mohó",
                    "dinamikus programozás",
                    "adatszerkezetek",
                ]} initTags={tags} onChange={handleTagsChange}/>
            </div>
            <div className="mb-5">
                <TextBoxDropdown id="filterCategory" label={t("problem_filter.category")} initText={category[1]} initSelected={category[0]}
                                 fillSelected={true} itemNames={[
                    "IOI-CEOI Válogató 2023",
                    "IOI-CEOI Válogató 2023 − 1. forduló",
                    "IOI-CEOI Válogató 2023 − 2. forduló",
                    "IOI-CEOI Válogató 2023 − 3. forduló"
                ]} onChange={handleCategoryChange}/>
            </div>
            <div className="flex justify-center">
                <button className="mr-1 btn-indigo padding-btn-default w-32" onClick={handleSubmit}>{t("problem_filter.search")}</button>
                <button className="ml-1 btn-gray padding-btn-default w-32" onClick={handleReset}>{t("problem_filter.reset")}</button>
            </div>
        </div>
    )
}

export function ProblemFilterFrame() {
    const {t} = useTranslation()
    return (
        <DropdownFrame title={t("problem_filter.filter")}>
            <div className="px-8 py-6">
                <ProblemFilter/>
            </div>
        </DropdownFrame>
    )
}

export default ProblemFilter;