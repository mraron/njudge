import {SVGSpinner} from "../svg/SVGs";
import {AnimatePresence, motion} from "framer-motion";

function PageLoadingAnimation({ isVisible }) {
    return (
        <AnimatePresence>
            {isVisible && (
                <motion.div initial={{ opacity: 0 }} animate={{ opacity: 1 }} exit={{ opacity: 0 }}>
                    <div className="absolute top-32 left-1/2 z-20">
                        <SVGSpinner cls="w-10 h-10" />
                    </div>
                </motion.div>
            )}
        </AnimatePresence>
    )
}

export default PageLoadingAnimation