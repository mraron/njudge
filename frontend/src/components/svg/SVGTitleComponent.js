function SVGTitleComponent({ title, svg }) {
    return (
        <div className="frame-title flex items-center justify-center">
            {svg}
            <span className="break-words min-w-0">{title}</span>
        </div>
    );
}

export default SVGTitleComponent;
