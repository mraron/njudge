function SVGTitleComponent({title, svg}) {
    return (
        <div className="py-3 px-4 border-b border-default font-medium flex items-center justify-center">
            {svg}
            <span>{title}</span>
        </div>
    )
}

export default SVGTitleComponent;