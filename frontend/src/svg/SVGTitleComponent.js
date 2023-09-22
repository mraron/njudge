function SVGTitleComponent({title, svg}) {
    return (
        <div className="p-4 border-b border-default font-medium flex items-center justify-center">
            {svg}
            <span>{title}</span>
        </div>
    )
}

export default SVGTitleComponent;