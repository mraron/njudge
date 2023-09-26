import { Link } from "react-router-dom";
import { useTranslation } from "react-i18next";
import RoundedFrame from "./RoundedFrame";

export function DefaultTag({ data }) {
    const { t } = useTranslation();
    return <span className="w-28 tag text-center">{t(data)}</span>;
}

export function LinkTag({ data }) {
    const { t } = useTranslation();
    return (
        <Link
            to={data.href}
            className="w-28 tag text-center hover:bg-indigo-200 hover:border-indigo-300 dark:hover:bg-indigo-600 dark:hover:border-transparent">
            {t(data.text)}
        </Link>
    );
}

function TagListFrame({ title, titleComponent, tag: Tag = DefaultTag, tags }) {
    console.log(tags);
    console.log("nigga");
    const tagsContent = tags.map((item, index) => (
        <div className="flex m-1" key={index}>
            <Tag data={item} key={index} />
        </div>
    ));
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
