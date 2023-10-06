function RoundedFrame({ children, title, titleComponent, cls = null }) {
    return (
        <div className={`${cls} rounded-frame text-frame`}>
            <div className="w-full flex flex-col">
                {title && <span className="frame-title break-words min-w-0">{title}</span>}
                {titleComponent}
                <div className="w-full min-h-[1rem]">{children}</div>
            </div>
        </div>
    )
}

export function SVGTitleComponent({ title, icon }) {
    return (
        <div className="frame-title flex items-center justify-center">
            {icon}
            <span className="break-words min-w-0">{title}</span>
        </div>
    )
}

export default RoundedFrame
