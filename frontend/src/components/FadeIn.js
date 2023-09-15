import {AnimatePresence, motion} from "framer-motion"

function FadeIn({children}) {
    return (
        <AnimatePresence>
            <motion.div initial={{opacity: 0.4}} animate={{opacity: 1}} exit={{opacity: 0.4}} transition={{duration: 0.3}}>
                {children}
            </motion.div>
        </AnimatePresence>
    )
}

export default FadeIn