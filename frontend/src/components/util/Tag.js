function Tag({ cls, children, addMargin = true }) {
    return (
        <div className={`${cls} tag ${addMargin ? "m-1" : ""}`}>
            <span className="whitespace-nowrap truncate">{children}</span>
        </div>
    )
}

export default Tag
