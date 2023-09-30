import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { AnimatePresence, motion } from "framer-motion";
import { TERipple } from "tw-elements-react";

function CopyButton({ text, isVisible }) {
    const handleCopy = () => {
        navigator.clipboard.writeText(text);
        window.flash("info.successful_copy", "success");
    };
    return (
        <AnimatePresence>
            {isVisible && (
                <motion.div
                    initial={{ opacity: 0 }}
                    animate={{ opacity: 1, transition: { duration: 0.2 } }}>
                    <TERipple>
                        <button
                            className={`rounded-md bg-grey-775 hover:bg-grey-750 border border-bordercol relative h-9 w-9`}
                            aria-label="Copy"
                            onClick={handleCopy}>
                            <FontAwesomeIcon
                                icon="fa-regular fa-copy"
                                className="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 w-3.5 h-3.5"
                            />
                        </button>
                    </TERipple>
                </motion.div>
            )}
        </AnimatePresence>
    );
}

export default CopyButton;
