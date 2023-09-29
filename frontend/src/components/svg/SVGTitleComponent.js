function SVGTitleComponent({ title, svg }) {
    return (
        <div className="p-4 border-b border-bordercol font-medium flex items-center justify-center">
            {svg}
            <span className="break-words min-w-0">{title}</span>
        </div>
    );
}

export default SVGTitleComponent;
