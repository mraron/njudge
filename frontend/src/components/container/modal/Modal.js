import { useEffect } from "react";
import { AnimatePresence, motion } from "framer-motion";

function Modal({ children, isOpen, onClose }) {
    useEffect(() => {
        if (isOpen) {
            document.getElementsByTagName("body")[0].style.overflow = "hidden"
        } else {
            document.getElementsByTagName("body")[0].style.overflow = "auto"
        }
    }, [isOpen])
    return (
        isOpen &&
        <div>
            <div className="z-40 fixed inset-0 bg-white dark:bg-black opacity-50" onClick={onClose} />
            <div
                className="z-50 fixed left-3 right-3 top-1/2 md:left-1/2 md:right-auto transform md:-translate-x-1/2 -translate-y-1/2"
                onClick={(event) => event.stopPropagation()}>
                <AnimatePresence>
                    <motion.div
                        initial={{ opacity: 0.2, y: 50 }}
                        animate={{ opacity: 1, y: 0, transition: { duration: 0.2 } }}>
                        {children}
                    </motion.div>
                </AnimatePresence>
            </div>
        </div>
    )
}

export default Modal
