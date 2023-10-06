import { Link } from "react-router-dom"
import RoundedFrame from "./RoundedFrame"
import Tag from "../basic/Tag"

export function DefaultTag({ data }) {
    return <Tag cls="w-28 justify-center">{data}</Tag>
}

export function LinkTag({ data }) {
    return (
        <Link to={data.href} className="m-1">
            <Tag
                cls="w-28 justify-center hover:bg-indigo-200 hover:border-indigo-400 dark:hover:bg-indigo-800 dark:hover:bg-indigo-600"
                addMargin={false}>
                {data.text}
            </Tag>
        </Link>
    )
}

function TagListFrame({ title, titleComponent, tag: Tag = DefaultTag, tags }) {
    const tagsContent = tags.map((item, index) => <Tag data={item} key={index} />)
    return (
        <RoundedFrame title={title} titleComponent={titleComponent} cls="overflow-hidden">
            <div
                className={`flex flex-col w-full ${
                    title || titleComponent ? "rounded-b-container" : "rounded-container"
                }`}>
                <div
                    className={`flex flex-wrap p-4 bg-grey-850 ${
                        title || titleComponent ? "rounded-b-container" : "rounded-container"
                    }`}>
                    {tagsContent}
                </div>
            </div>
        </RoundedFrame>
    )
}

export default TagListFrame
