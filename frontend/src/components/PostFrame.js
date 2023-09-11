import { useState } from 'react';
import RoundedFrame from './RoundedFrame'

function PostItem({name, date, content}) {
    const [truncated, setTruncated] = useState(true)
    return (
        <RoundedFrame>
            <div className="px-6 py-5 sm:px-10 sm:py-8 hover:bg-grey-775 transition duration-200 cursor-pointer rounded-md" onClick={() => setTruncated(!truncated)}>
                <div className="flex justify-between items-start">
                    <span className="text-lg font-semibold">{name}</span>
                    <span className="ml-4 date-label">{date}</span>
                </div>
                <div className={`mt-2 ${truncated? "truncate": ""}`}>
                    {content}
                </div>
            </div>
        </RoundedFrame>
    )
}

function PostFrame({newsData}) {
    const newsItems = newsData.map((data, index) => {
        return (
            <div className="mb-3" key={index}>
                <PostItem name={data[0]} date={data[1]} content={data[2]} />
            </div>
        );
    });
    return (
        <div>
            {newsItems}
        </div>
    )
}

export default PostFrame;