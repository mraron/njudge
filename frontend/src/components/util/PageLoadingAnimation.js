import { AnimatePresence, motion } from "framer-motion";
import { SVGSpinner } from "../svg/SVGs";

function PageLoadingAnimation({ isVisible }) {
    return (
        <AnimatePresence>
            {isVisible && (
                <motion.div
                    initial={{ opacity: 0.01 }}
                    animate={{
                        opacity: 1,
                        transition: { delay: 0.2, duration: 0.8 },
                    }}
                    exit={{ opacity: 0.01, transition: { duration: 0.1 } }}>
                    <div className="absolute top-16 left-1/2 z-10">
                        <SVGSpinner cls="w-10 h-10" />
                    </div>
                </motion.div>
            )}
        </AnimatePresence>
    );
}

export default PageLoadingAnimation;
