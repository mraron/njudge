function Tag({ cls, children }) {
    return (
        <div className={`${cls} tag`}>
            <span className="whitespace-nowrap truncate">{children}</span>
        </div>
    );
}

export default Tag;
