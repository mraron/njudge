import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

function CopyButton({ text }) {
    const handleCopy = () => {
        navigator.clipboard.writeText(text);
        window.flash("info.successful_copy", "success");
    };
    return (
        <button
            className={`relative h-9 w-9 bg-grey-775 rounded-md hover:bg-grey-750 border border-grey-625`}
            aria-label="Copy"
            onClick={handleCopy}>
            <FontAwesomeIcon
                icon="fa-regular fa-copy"
                className="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 w-3.5 h-3.5"
            />
        </button>
    );
}

export default CopyButton;
