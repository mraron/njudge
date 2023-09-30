import { useState } from "react";
import RoundedFrame from "./RoundedFrame";

function PostItem({ post }) {
    const [truncated, setTruncated] = useState(true);
    const { title, content, date } = post;
    return (
        <RoundedFrame>
            <div
                className="px-6 py-5 sm:px-10 sm:py-8 hover:bg-grey-775 cursor-pointer rounded-md"
                onClick={() => setTruncated(!truncated)}>
                <div className="flex justify-between items-start">
                    <span className="text-base font-semibold break-words min-w-0">
                        {title}
                    </span>
                    <span className="ml-4 date-label">{date}</span>
                </div>
                <div
                    className={`mt-2 ${
                        truncated ? "truncate" : "break-words"
                    }`}>
                    {content}
                </div>
            </div>
        </RoundedFrame>
    );
}

function PostFrame({ posts }) {
    const newsItems = posts.map((item, index) => {
        return (
            <div className="mb-3" key={index}>
                <PostItem post={item} />
            </div>
        );
    });
    return <div>{newsItems}</div>;
}

export default PostFrame;
