function RoundedFrame({ children, title, titleComponent, cls = null }) {
    return (
        <div className={`${cls} rounded-frame`} style={{ fontSize: "0.96rem" }}>
            <div className="w-full flex flex-col">
                {title && (
                    <span className="frame-title break-words min-w-0">
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
