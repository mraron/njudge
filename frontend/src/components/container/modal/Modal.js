import { useContext } from "react"
import { AnimatePresence, motion } from "framer-motion"
import { Box, Modal as MUIModal } from "@mui/material"
import ThemeContext from "../../../contexts/theme/ThemeContext"

function Modal({ children, isOpen, onClose, width }) {
    const { theme } = useContext(ThemeContext)
    return (
        <MUIModal
            open={isOpen}
            onClose={onClose}
            BackdropProps={{
                style: { backgroundColor: theme === "dark" ? "rgba(255, 255, 255, 0.25)" : "rgba(0, 0, 0, 0.25)" },
            }}>
            <AnimatePresence>
                <div className="fixed mui-fixed left-2 right-2 top-1/2 sm:left-1/2 sm:right-auto transform sm:-translate-x-1/2 -translate-y-1/2">
                    <motion.div
                        initial={{ opacity: 0.2, y: 20 }}
                        animate={{ opacity: 1, y: 0, transition: { duration: 0.15 } }}>
                        <Box sx={{ width: { xs: "100%", sm: width } }}>{children}</Box>
                    </motion.div>
                </div>
            </AnimatePresence>
        </MUIModal>
    )
}

export default Modal
