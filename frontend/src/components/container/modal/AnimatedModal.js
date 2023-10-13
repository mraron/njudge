import { useContext } from "react"
import { AnimatePresence, motion } from "framer-motion"
import { Modal } from "@mui/material"
import ThemeContext from "../../../contexts/theme/ThemeContext"

function AnimatedModal({ children, isOpen, onClose }) {
    const { theme } = useContext(ThemeContext)
    return (
        <Modal
            open={isOpen}
            onClose={onClose}
            BackdropProps={{
                style: { backgroundColor: theme === "light" ? "rgba(255, 255, 255, 0.6)" : "rgba(0, 0, 0, 0.6)" },
            }}>
            <AnimatePresence>
                <div className="fixed mui-fixed left-2 right-2 top-1/2 md:left-1/2 md:right-auto transform md:-translate-x-1/2 -translate-y-1/2">
                    <motion.div
                        initial={{ opacity: 0.2, y: 20 }}
                        animate={{ opacity: 1, y: 0, transition: { duration: 0.15 } }}>
                        {children}
                    </motion.div>
                </div>
            </AnimatePresence>
        </Modal>
    )
}

export default AnimatedModal
