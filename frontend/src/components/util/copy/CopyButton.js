import { SVGCopy } from "../../svg/SVGs";

function CopyButton({ text }) {
    const handleCopy = () => {
        navigator.clipboard.writeText(text);
        window.flash("info.successful_copy", "success");
    };
    return (
        <button
            className={`relative h-9 w-9 bg-grey-775 rounded-md hover:bg-grey-750 border border-bordercol`}
            aria-label="Copy"
            onClick={handleCopy}>
            <SVGCopy
                cls={`absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 w-4 h-4`}
            />
        </button>
    );
}

export default CopyButton;
