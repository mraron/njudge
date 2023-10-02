function Tag({ cls, children }) {
    return (
        <div className={`${cls} tag m-1`}>
            <span className="whitespace-nowrap truncate">{children}</span>
        </div>
    );
}

export default Tag;
