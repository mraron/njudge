import { motion, AnimatePresence } from "framer-motion"
import {useEffect, useState} from "react";
function FadeIn({ children }) {
    const [isVisible, setVisible] = useState(true)
    useEffect(() => {
        setVisible(true)
        return () => {
            setVisible(false)
        }
    }, []);
    return (
        <AnimatePresence>
            {isVisible && (
                <motion.div initial={{ opacity: 0 }} animate={{ opacity: 1 }} exit={{ opacity: 0 }} transition={{ duration: 0.3 }}>
                    {children}
                </motion.div>
            )}
        </AnimatePresence>
    )
}

export default FadeIn