function RoundedFrame({children, title, titleComponent}) {
    return (
        <div className="bg-grey-800 border-1 rounded-md flex flex-col border-default w-full">
            <div className="flex flex-col">
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