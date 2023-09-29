function RoundedFrame({ children, title, titleComponent, cls = null }) {
    return (
        <div className={`${cls} rounded-frame`}>
            <div className="w-full flex flex-col">
                {title && (
                    <span className="font-medium p-4 text-center border-b-1 border-grey-700 break-words min-w-0">
                        {title}
                    </span>
                )}
                {titleComponent}
                <div className="w-full text-dropdown-list min-h-[1rem]">
                    {children}
                </div>
            </div>
        </div>
    );
}

export default RoundedFrame;
