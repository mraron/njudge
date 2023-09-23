import RoundedFrame from "./RoundedFrame";
import {useTranslation} from "react-i18next";
import {Link} from "react-router-dom";

export function DefaultTag({data}) {
    const {t} = useTranslation()
    return (
        <span
            className="w-28 tag text-center">
            {t(data)}
        </span>
    )
}

export function LinkTag({data}) {
    const {t} = useTranslation()
    return (
        <Link to={data.href}
              className="w-28 text-center truncate whitespace-nowrap cursor-pointer text-sm px-2 py-1 border-1 rounded bg-grey-725 hover:bg-indigo-600 hover:border-transparent border-grey-650 transition-all duration-200">
            {t(data.text)}
        </Link>
    )
}

function TagListFrame({title, titleComponent, tag: Tag, tags}) {
    Tag ||= DefaultTag
    const tagsContent = tags.map((item, index) =>
        <div className="flex m-1" key={index}>
            <Tag data={item} key={index}/>
        </div>
    )
    return (
        <RoundedFrame title={title} titleComponent={titleComponent}>
            <div className="flex flex-col w-full overflow-x-auto rounded-md">
                <div className="flex flex-wrap p-4 bg-grey-850">
                    {tagsContent}
                </div>
            </div>
        </RoundedFrame>
    );
}

export default TagListFrame;