import { useContext, useState } from "react"
import { useLocation, useNavigate } from "react-router-dom"
import { useTranslation } from "react-i18next"
import TextBox from "../../input/TextBox"
import TextBoxDropdown from "../../input/TextBoxDropdown"
import TagDropdown from "../../input/TagDropdown"
import DropdownFrame from "../../container/DropdownFrame"
import updateQueryString from "../../../util/updateQueryString"
import JudgeDataContext from "../../../contexts/judgeData/JudgeDataContext"
import queryString from "query-string"
import { parseInt } from "lodash"
import Button from "../../basic/Button"

function ProblemFilter() {
    const { t } = useTranslation()
    const { judgeData } = useContext(JudgeDataContext)
    const location = useLocation()
    const navigate = useNavigate()

    const parseTitle = (title) => {
        return title ? title : ""
    }
    const parseCategory = (category) => {
        if (category === undefined) {
            return -1
        }
        const categoryInt = parseInt(category)
        return judgeData.categories.findIndex(
            (item) => item.value === categoryInt,
        )
    }
    const parseTags = (tags) => {
        if (tags === undefined) {
            return []
        }
        const tokens = tags.split(",").map(parseInt)
        if (
            tokens.some(
                (elem) =>
                    isNaN(elem) || elem <= -1 || elem >= judgeData.tags.length,
            )
        ) {
            return []
        }
        return tokens
    }
    const qData = queryString.parse(location.search)
    const [title, setTitle] = useState(parseTitle(qData.title))
    const [tags, setTags] = useState(parseTags(qData.tags))
    const [category, setCategory] = useState(parseCategory(qData.category))

    const handleTitleChange = (newText) => {
        setTitle(newText)
    }
    const handleCategoryChange = (selected, newText) => {
        setCategory(selected)
    }
    const handleTagsChange = (tags) => {
        setTags(tags)
    }
    const handleSubmit = () => {
        updateQueryString({
            location: location,
            navigate: navigate,
            args: ["title", "tags", "category"],
            values: [
                title,
                tags.join(","),
                category === -1 ? -1 : judgeData.categories[category].value,
            ],
            validArgs: ["title", "tags", "category"],
        })
    }
    const handleReset = () => {
        updateQueryString({
            location: location,
            navigate: navigate,
            validArgs: [],
        })
    }
    return (
        <div className="w-full">
            <div className="space-y-4 mb-5">
                <TextBox
                    id="filterTitle"
                    label={t("problem_filter.title")}
                    initText={title}
                    onChange={handleTitleChange}
                />
                <TagDropdown
                    id="filterTags"
                    label={t("problem_filter.tags")}
                    fillSelected={false}
                    itemNames={judgeData.tags.map(t)}
                    initTags={tags}
                    onChange={handleTagsChange}
                />
                <TextBoxDropdown
                    id="filterCategory"
                    label={t("problem_filter.category")}
                    initText={
                        category === -1
                            ? ""
                            : judgeData.categories[category].label
                    }
                    initSelected={category}
                    fillSelected={true}
                    itemNames={judgeData.categories.map((x) => x.label)}
                    onChange={handleCategoryChange}
                />
            </div>
            <div className="flex justify-center space-x-2">
                <Button color="indigo" minWidth="8rem" onClick={handleSubmit}>
                    {t("problem_filter.filter")}
                </Button>
                <Button color="gray" minWidth="8rem" onClick={handleReset}>
                    {t("problem_filter.reset")}
                </Button>
            </div>
        </div>
    )
}

export function ProblemFilterFrame() {
    const { t } = useTranslation()
    return (
        <DropdownFrame title={t("problem_filter.filter")}>
            <div className="px-8 py-6">
                <ProblemFilter />
            </div>
        </DropdownFrame>
    )
}

export default ProblemFilter
