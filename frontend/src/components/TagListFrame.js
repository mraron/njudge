import RoundedFrame from "./RoundedFrame";

function Tag({ tagName }) {
    return (
        <span className="w-28 text-center truncate whitespace-nowrap cursor-pointer text-sm px-2 py-1 border-1 rounded m-1 bg-grey-725 hover:bg-indigo-600 border-grey-650 hover:border-indigo-500 transition-all duration-200">
            {tagName}
        </span>
    )
}

function TagListFrame({ title, titleComponent, tagNames }) {
    const tags = tagNames.map((tagName, index) => <Tag tagName={tagName} key={index} />)
    return (
        <RoundedFrame title={title} titleComponent={titleComponent}>
            <div className="flex flex-col w-full overflow-x-auto rounded-md">
                <div className="flex flex-wrap p-4 bg-grey-850">
                    {tags}
                </div>
            </div>
        </RoundedFrame>
    );
}

export default TagListFrame;