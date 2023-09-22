function RoundedFrame({children, title, titleComponent, cls}) {
    return (
        <div className={`${cls} bg-grey-800 border-1 rounded-md flex border-default w-full`}>
            <div className="w-full flex flex-col">
                {title && <span className="font-medium px-6 py-4 text-center border-b-1 border-grey-700">{title}</span>}
                {titleComponent}
                <div className="w-full text-dropdown-list">
                    {children}
                </div>
            </div>
        </div>
    );
}

export default RoundedFrame;