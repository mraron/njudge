import {SVGSpinner} from "../svg/SVGs";
import {AnimatePresence, motion} from "framer-motion";

function PageLoadingAnimation({ isVisible }) {
    return (
        <AnimatePresence>
            {isVisible && (
                <motion.div initial={{ opacity: 0.2 }} animate={{ opacity: 1 }} exit={{ opacity: 0.2 }} transition={{ duration: 0.2 }}>
                    <div className="absolute top-16 left-1/2 z-20">
                        <SVGSpinner cls="w-10 h-10" />
                    </div>
                </motion.div>
            )}
        </AnimatePresence>
    )
}

export default PageLoadingAnimation