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
                <div className="flex justify-between items-start space-x-4">
                    <span className="text-base emph-strong break-words min-w-0">
                        {title}
                    </span>
                    <span className="date-label">{date}</span>
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
    const newsItems = posts.map((item, index) => (
        <PostItem post={item} key={index} />
    ));
    return <div className="space-y-3">{newsItems}</div>;
}

export default PostFrame;
